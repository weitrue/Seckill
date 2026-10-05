/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/10 上午9:49
 * Description: 内存预扣库存
 *   - 职责:在把请求丢进异步队列之前,先做一次原子内存扣减,快速过滤掉绝大多数"明显超额"的请求,
 *     让真正到达 Redis(由 redisstock 做 Lua 原子扣减+去重)的流量不至于被打爆
 *   - 冷启动:懒加载,首次 Sub 时通过外部注入的 Loader 拉取当前库存写入内存
 *     Loader 由 application 层组装(通常是闭包: redisstock.NewRedisStock + Get)
 *   - 并发:同 (activityID,goodsID) 在同一个 *Factory 内复用 *memoryStack 实例,库存值用 atomic.Int64 操作
 *   - 预扣失败回滚:Sub 发现 decr 后 < 0 立刻把那 1 还回去,避免负值累积
 *   - 多实例一致性:每个进程各自预扣,允许总预扣数 > 真实库存,Redis Lua 作为最终真实源兜底
 *   - 以下方法暂未实现(没有调用点),保留 panic 占位:Set / Get / Del
 **/

package memstock

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/weitrue/Seckill/internal/domain/stock"
	"github.com/weitrue/Seckill/internal/infrastructure/config"
)

// Loader 内存库存的冷启动数据源
//
//	首次扣减某 (activityID,goodsID) 时,通过 Loader 读取当前库存写入内存
//	典型实现: 调用 redisstock.NewRedisStock(aid,gid).Get(ctx)
type Loader func(ctx context.Context, activityID, goodsID string) (int64, error)

// Factory memstock 工厂,持有 Loader 和实例缓存
//
//	一个 Factory 内 (activityID, goodsID) 共享一个 memoryStack 实例
//	不同 Factory 之间互不干扰,方便单测
type Factory struct {
	loader Loader
	stacks sync.Map // map[string]*memoryStack
}

/*NewFactory
 *@Description: 构造 memstock Factory
 *@param loader 冷启动数据源(nil 时首次 Sub 会返回错误)
 *@return *Factory
 */
func NewFactory(loader Loader) *Factory {
	return &Factory{loader: loader}
}

/*New
 *@Description: 创建或复用一个内存库存实例(同 key 返回同一个);
 *   签名匹配 stock.Factory,可作为 method value 注入 application
 *@receiver f
 *@param activityID
 *@param goodsID
 *@return stock.Stock
 *@return error
 */
func (f *Factory) New(activityID, goodsID string) (stock.Stock, error) {
	if activityID == "" || goodsID == "" {
		return nil, errors.New("invalid event id or goods id")
	}
	key := fmt.Sprintf("%s:%s:%s", config.GetServiceName(), activityID, goodsID)
	if v, ok := f.stacks.Load(key); ok {
		return v.(*memoryStack), nil
	}
	fresh := &memoryStack{
		activityID: activityID,
		goodsID:    goodsID,
		key:        key,
		loader:     f.loader,
	}
	actual, _ := f.stacks.LoadOrStore(key, fresh)
	return actual.(*memoryStack), nil
}

// memoryStack 内存库存单元,一个 (activity,goods) 对应一个
type memoryStack struct {
	activityID string // 活动ID
	goodsID    string // 商品ID
	key        string // 库存 key,序列化成 serviceName:activityID:goodsID
	loader     Loader // 冷启动数据源

	stock    int64     // 原子库存计数(通过 atomic 操作)
	initOnce sync.Once // 懒加载保护
	initErr  error     // 懒加载错误保存
}

// 编译期类型断言: 确保 *memoryStack 实现了 stock.Stock
var _ stock.Stock = (*memoryStack)(nil)

/*lazyInit
 *@Description: 懒加载,首次 Sub 时通过 Loader 拉取当前库存
 *@receiver m
 *@param ctx 请求上下文,用于后续链路追踪与超时控制
 *@return error 加载失败的错误(initOnce 保证只执行一次,后续 Sub 直接返回这个 err)
 */
func (m *memoryStack) lazyInit(ctx context.Context) error {
	m.initOnce.Do(func() {
		if m.loader == nil {
			m.initErr = errors.New("memstock: loader is nil, cannot cold-start")
			return
		}
		val, err := m.loader(ctx, m.activityID, m.goodsID)
		if err != nil {
			m.initErr = err
			return
		}
		atomic.StoreInt64(&m.stock, val)
	})
	return m.initErr
}

/*Sub
 *@Description: 内存预扣,原子操作;扣后 < 0 自动回滚那 1
 *@receiver m
 *@param ctx 请求上下文
 *@param uid 用户 ID(内存层不做 uid 去重,去重由 redisstock Lua 兜底)
 *@return int64 预扣后剩余库存;预扣失败返回 -1
 *@return error 冷启动失败时返回 err,预扣失败本身不返 err(与 redisstock.Sub 语义对齐)
 */
func (m *memoryStack) Sub(ctx context.Context, uid string) (int64, error) {
	if err := m.lazyInit(ctx); err != nil {
		return -1, err
	}
	v := atomic.AddInt64(&m.stock, -1)
	if v < 0 {
		// 预扣失败,把这 1 还回去,避免负值累积
		atomic.AddInt64(&m.stock, 1)
		return -1, nil
	}
	return v, nil
}

/*GetActivityID
 *@Description: 返回活动 ID
 *@receiver m
 *@return string
 */
func (m *memoryStack) GetActivityID() string {
	return m.activityID
}

/*GetGoodsID
 *@Description: 返回商品 ID
 *@receiver m
 *@return string
 */
func (m *memoryStack) GetGoodsID() string {
	return m.goodsID
}

// Set 暂未实现:库存写入由 redisstock 承担,内存层通过 Loader 读取
func (m *memoryStack) Set(ctx context.Context, val, expire int64) error {
	panic("implement me")
}

// Get 暂未实现:当前没有调用点,如需快照可读取 atomic.LoadInt64(&m.stock)
func (m *memoryStack) Get(ctx context.Context) (int64, error) {
	panic("implement me")
}

// Del 暂未实现:当前没有调用点
func (m *memoryStack) Del(ctx context.Context) error {
	panic("implement me")
}
