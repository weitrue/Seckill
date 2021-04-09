/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/9 下午12:55
 * Description: etcd 集群存储节点元信息配置
 **/

package cluster

import (
	"Seckill/infrastructure/stores/etcd"
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	"github.com/sirupsen/logrus"
)

var (
	defaultCluster *cluster
	once           = &sync.Once{}
)

type Node struct {
	Addr    string `json:"addr"`    // 节点服务地址
	Version string `json:"version"` // 节点服务版本
	Proto   string `json:"proto"`   // 协议
}

type cluster struct {
	one            sync.Once
	client         *clientv3.Client         // etcd 客户端
	service        string                   // 服务
	deregisterChan map[string]chan struct{} // 需要注销节点的任务列表
	nodes          map[string]*Node         // 节点列表
	sync.RWMutex
}

func Init(service string) {
	// 初始化集群服务信息
	once.Do(func() {
		defaultCluster = &cluster{
			client:         etcd.GetETCDClient(),
			service:        service,
			one:            sync.Once{},
			deregisterChan: make(map[string]chan struct{}),
			nodes:          make(map[string]*Node),
		}
	})
}

func (c *cluster) getNodeKey(node *Node) string {
	// 格式化key serviceName/nodes/xx-xx-xx-xx-port
	id := strings.Replace(node.Addr, ".", "-", -1)
	id = strings.Replace(id, ":", "-", -1)
	return fmt.Sprintf("/%s/nodes/%s", c.service, id)
}

// 节点管理--Register注册节点,DeRegister注销节点,Discover发现节点
func Register(node *Node, ttl int) error {
	// 注册节点信息到 etcd 并设置有效期
	const minTTl = 2
	c := defaultCluster
	// 根据节点信息格式化key
	key := c.getNodeKey(node)
	if ttl < minTTl {
		ttl = minTTl
	}
	var errChan = make(chan error)
	go func() {
		kv := clientv3.NewKV(c.client)
		closeChan := make(chan struct{})
		// 创建租约对象
		lease := clientv3.NewLease(c.client)
		// 序列化node
		val, _ := json.Marshal(node)
		var curLeaseId clientv3.LeaseID = 0
		ticker := time.NewTicker(time.Duration(ttl/2) * time.Second)

		// 注册逻辑
		register := func() error {
			if curLeaseId == 0 {
				// 租约已失效或者未注册，需要创建租约
				leaseResp, err := lease.Grant(context.TODO(), int64(ttl))
				if err != nil {
					return err
				}
				// key:serviceName/nodes/xx-xx-xx-xx-port, val
				if _, err := kv.Put(context.TODO(), key, string(val), clientv3.WithLease(leaseResp.ID)); err != nil {
					return err
				}
				curLeaseId = leaseResp.ID
			} else {
				// 续租租约，如果租约已经过期将curLeaseId复位到0重新走创建租约的逻辑
				if _, err := lease.KeepAlive(context.TODO(), curLeaseId); err == rpctypes.ErrLeaseNotFound {
					curLeaseId = 0
				}
			}
			return nil
		}
		if err := register(); err != nil {
			logrus.Error("register node failed, error:", err)
			errChan <- err
		}
		close(errChan)
		for {
			select {
			case <-ticker.C:
				if err := register(); err != nil {
					logrus.Error("register node failed, error:", err)
					panic(err)
				}
			case <-closeChan:
				ticker.Stop()
				return
			}
		}
	}()
	err := <-errChan
	return err
}

func Deregister(node *Node) error {
	// 从 etcd 中删除节点信息
	c := defaultCluster
	c.Lock()
	defer c.Unlock()
	key := c.getNodeKey(node)
	if ch, ok := c.deregisterChan[key]; ok {
		close(ch)
		delete(c.deregisterChan, key)
	}
	_, err := c.client.Delete(context.Background(), key, clientv3.WithPrefix())
	return err
}

func Discover() (output []*Node, err error) {
	// 发现新节点
	c := defaultCluster
	key := fmt.Sprintf("/%s/nodes/", c.service)
	// 初始化时记录集群中的节点
	c.one.Do(func() {
		var resp *clientv3.GetResponse
		resp, err = c.client.Get(context.Background(), key, clientv3.WithPrefix())
		if err != nil {
			return
		}

		for _, kv := range resp.Kvs {
			k := string(kv.Key)
			if len(k) > len(key) {
				var node *Node
				// 反序列化
				json.Unmarshal(kv.Value, &node)
				if node != nil {
					c.Lock()
					c.nodes[k] = node
					c.Unlock()
				}
			}
		}
		watch := c.client.Watch(context.Background(), key, clientv3.WithPrefix())
		go func() {
			for {
				select {
				case resp := <-watch:
					for _, event := range resp.Events {
						k := string(event.Kv.Key)
						if len(k) <= len(key) {
							continue
						}

						switch event.Type {
						case clientv3.EventTypePut:
							var node *Node
							json.Unmarshal(event.Kv.Value, &node)
							if node != nil {
								c.Lock()
								c.nodes[k] = node
								c.Unlock()
							}
						case clientv3.EventTypeDelete:
							c.Lock()
							if _, ok := c.nodes[k]; ok {
								delete(c.nodes, k)
							}
							c.Unlock()
						}
					}
				}
			}
		}()
	})
	if err != nil {
		return nil, err
	}
	// 加写锁
	c.RLock()
	for _, node := range c.nodes {
		output = append(output, node)
	}
	c.RUnlock()
	return
}
