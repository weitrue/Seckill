/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/13 下午8:32
 * Description:
 **/

package shop

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/weitrue/Seckill/domain/stock"
	"github.com/weitrue/Seckill/infrastructure/factory"
	"github.com/weitrue/Seckill/infrastructure/mq/mqi"
	"github.com/weitrue/Seckill/infrastructure/pool/worker/taski"
	"github.com/weitrue/Seckill/infrastructure/utils"

	"github.com/sirupsen/logrus"
)

var queue mqi.Queue

func Init() {
	// 使用漏桶 限制下游redis消费速度
	queueFactory := factory.NewFactory("memory")
	if queueFactory == nil {
		panic("no memory queue factory")
	}
	queue, _ = queueFactory.New("shop")
	go func() {
		for {
			task, err := queue.Consume()
			if err != nil {
				logrus.Error(err)
				break
			}
			task.Do()
		}
	}()
}

const (
	OK         = 0
	ErrNoStock = 1001
	ErrRedis   = 1002
	ErrTimeout = 1003

	requestTimeout = 60
)

type Context struct {
	Request *http.Request
	Conn    net.Conn
	Writer  *bufio.ReadWriter
	GoodsID string
	EventID string
	UID     string
}

func Handle(ctx *Context) {
	// 处理函数 请求->是否熔断->令牌桶是否限流->业务逻辑->漏桶是否限流->减库存----
	//                        ｜                       ｜                ｜
	//                        ｜请求频繁                ｜系统繁忙         ｜
	//                        —                       <-            <-返回
	start := time.Now().Unix()
	t := func() {
		data := &utils.Response{
			Code: OK,
			Data: nil,
			Msg:  "ok",
		}
		status := http.StatusOK
		now := time.Now().Unix()
		if now-start > requestTimeout {
			data.Msg = "request timeout"
			data.Code = ErrTimeout
		} else {
			// 扣减 Redis 库存
			st, _ := stock.NewMemoryStock(ctx.EventID, ctx.GoodsID)
			if s, err := st.Sub(ctx.UID); err != nil {
				data.Msg = err.Error()
				data.Code = ErrRedis
			} else if s < 0 {
				data.Msg = "no stock"
				data.Code = ErrNoStock
			}
		}
		// 此处实现操作购物车的逻辑

		body, _ := json.Marshal(data)
		resp := &http.Response{
			Proto:         ctx.Request.Proto,
			ProtoMinor:    ctx.Request.ProtoMinor,
			ProtoMajor:    ctx.Request.ProtoMajor,
			Header:        make(http.Header),
			ContentLength: int64(len(body)),
			Body:          ioutil.NopCloser(bytes.NewReader(body)),
			StatusCode:    status,
			Close:         false,
		}
		resp.Header.Set("Content-Type", "application/json")
		resp.Write(ctx.Writer)
		ctx.Writer.Flush()
		ctx.Conn.Close()
	}

	queue.Produce(taski.TaskFunc(t))
}