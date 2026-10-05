/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2024/10/05
 * Description: rabbitmq 实现 mq.PubSub(Publish/Subscribe []byte + Close)
 *   基于 github.com/rabbitmq/amqp091-go
 **/

package rabbitmq

import (
	"errors"
	"fmt"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/weitrue/Seckill/infrastructure/mq"
)

// ConnParams rabbitmq 连接参数
type ConnParams struct {
	Addr         string // amqp://user:pwd@host:port/vhost
	Exchange     string // exchange 名称
	ExchangeType string // direct | topic | fanout | headers,默认 direct
	Durable      bool   // 持久化 exchange / queue
	AutoAck      bool   // 消费回调成功后是否自动 ack
}

// rabbitmqQueue rabbitmq pub/sub 实现
type rabbitmqQueue struct {
	params *ConnParams
	conn   *amqp.Connection

	pubMu sync.Mutex    // 保护 pubCh 懒初始化
	pubCh *amqp.Channel // publisher 复用的 channel

	subMu    sync.Mutex      // 保护 subChs
	subChs   []*amqp.Channel // 每个 Subscribe 新开一个 channel
	closed   bool
	closedMu sync.RWMutex
}

// 编译期类型断言: 确保 rabbitmqQueue 实现了 mq.PubSub
var _ mq.PubSub = (*rabbitmqQueue)(nil)

/*Publish
 *@Description: 向 rabbitmq 发布消息(topic 作为 routing key)
 *@receiver r
 *@param topic 路由键
 *@param body  消息体
 *@param reqID 请求 ID,写入 MessageId 用于链路追踪
 *@return error
 */
func (r *rabbitmqQueue) Publish(topic string, body []byte, reqID string) error {
	logType := "rabbitmqQueuePublish"
	if r.isClosed() {
		return errors.New("rabbitmq queue closed")
	}
	ch, err := r.ensurePubChannel()
	if err != nil {
		logrus.Errorf("logType:%s, err:%s, topic:%s, reqID:%s",
			logType, err.Error(), topic, reqID)
		return err
	}
	err = ch.Publish(
		r.params.Exchange, // exchange
		topic,             // routing key
		false,             // mandatory
		false,             // immediate
		amqp.Publishing{
			ContentType:  "application/octet-stream",
			DeliveryMode: amqp.Persistent,
			MessageId:    reqID,
			Body:         body,
		},
	)
	if err != nil {
		logrus.Errorf("logType:%s, err:%s, topic:%s, reqID:%s",
			logType, err.Error(), topic, reqID)
	}
	return err
}

/*Subscribe
 *@Description: 订阅 rabbitmq 消息
 *@receiver r
 *@param topic 路由键(绑定到 exchange)
 *@param group 消费组,对应 queue 名,同名 queue 的多实例共享消费
 *@param fn    消息回调,返回 nil 表示处理成功
 *@return error
 */
func (r *rabbitmqQueue) Subscribe(topic, group string, fn mq.MsgCb) error {
	logType := "rabbitmqQueueSubscribe"
	if r.isClosed() {
		return errors.New("rabbitmq queue closed")
	}
	if group == "" {
		return errors.New("rabbitmq subscribe: group(queue name) is required")
	}
	ch, err := r.conn.Channel()
	if err != nil {
		logrus.Errorf("logType:%s, err:%s, topic:%s, group:%s",
			logType, err.Error(), topic, group)
		return err
	}
	if err = r.declare(ch, group, topic); err != nil {
		_ = ch.Close()
		logrus.Errorf("logType:%s, err:%s, topic:%s, group:%s",
			logType, err.Error(), topic, group)
		return err
	}
	// autoAck 由应用层控制: 这里关闭 amqp 自带 autoAck,改由回调结果决定
	deliveries, err := ch.Consume(
		group, // queue
		"",    // consumer tag
		false, // auto-ack(交给应用层回调决定)
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		_ = ch.Close()
		logrus.Errorf("logType:%s, err:%s, topic:%s, group:%s",
			logType, err.Error(), topic, group)
		return err
	}

	r.subMu.Lock()
	r.subChs = append(r.subChs, ch)
	r.subMu.Unlock()

	go func() {
		for d := range deliveries {
			msg := &rabbitmqMsg{delivery: d}
			if cbErr := fn(msg); cbErr == nil {
				if r.params.AutoAck {
					if ackErr := msg.Ack(); ackErr != nil {
						logrus.Errorf("logType:%s, err:%s, key:%s, topic:%s, group:%s",
							logType, ackErr.Error(), msg.GetID(), topic, group)
					}
				}
			} else {
				logrus.Errorf("logType:%s, err:%s, key:%s, topic:%s, group:%s",
					logType, cbErr.Error(), msg.GetID(), topic, group)
				// 回调失败,nack 并重入队列
				if nackErr := d.Nack(false, true); nackErr != nil {
					logrus.Errorf("logType:%s, err:%s, key:%s, topic:%s, group:%s, action:nack",
						logType, nackErr.Error(), msg.GetID(), topic, group)
				}
			}
		}
		logrus.Infof("logType:%s, msg:deliveries closed, topic:%s, group:%s",
			logType, topic, group)
	}()

	return nil
}

/*Close
 *@Description: 关闭所有 channel 和 connection
 *@receiver r
 *@return error
 */
func (r *rabbitmqQueue) Close() error {
	r.closedMu.Lock()
	if r.closed {
		r.closedMu.Unlock()
		return nil
	}
	r.closed = true
	r.closedMu.Unlock()

	r.subMu.Lock()
	for _, ch := range r.subChs {
		_ = ch.Close()
	}
	r.subChs = nil
	r.subMu.Unlock()

	r.pubMu.Lock()
	if r.pubCh != nil {
		_ = r.pubCh.Close()
		r.pubCh = nil
	}
	r.pubMu.Unlock()

	if r.conn != nil {
		return r.conn.Close()
	}
	return nil
}

// ensurePubChannel 懒初始化 publisher channel 并声明 exchange
func (r *rabbitmqQueue) ensurePubChannel() (*amqp.Channel, error) {
	r.pubMu.Lock()
	defer r.pubMu.Unlock()
	if r.pubCh != nil {
		return r.pubCh, nil
	}
	ch, err := r.conn.Channel()
	if err != nil {
		return nil, err
	}
	if err = r.declareExchange(ch); err != nil {
		_ = ch.Close()
		return nil, err
	}
	r.pubCh = ch
	return ch, nil
}

// declareExchange 声明 exchange(Publish/Subscribe 两侧都需要)
func (r *rabbitmqQueue) declareExchange(ch *amqp.Channel) error {
	if r.params.Exchange == "" {
		// 允许使用默认 exchange(空字符串),此时 routing key = queue name
		return nil
	}
	return ch.ExchangeDeclare(
		r.params.Exchange,
		r.params.ExchangeType,
		r.params.Durable,
		false, // auto-delete
		false, // internal
		false, // no-wait
		nil,   // args
	)
}

// declare 声明 exchange + queue + binding
func (r *rabbitmqQueue) declare(ch *amqp.Channel, queueName, routingKey string) error {
	if err := r.declareExchange(ch); err != nil {
		return err
	}
	if _, err := ch.QueueDeclare(
		queueName,
		r.params.Durable,
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // args
	); err != nil {
		return err
	}
	if r.params.Exchange == "" {
		// 默认 exchange 下不需要显式绑定
		return nil
	}
	return ch.QueueBind(
		queueName,
		routingKey,
		r.params.Exchange,
		false, // no-wait
		nil,   // args
	)
}

func (r *rabbitmqQueue) isClosed() bool {
	r.closedMu.RLock()
	defer r.closedMu.RUnlock()
	return r.closed
}

/*NewRabbitMQ
 *@Description: 创建 rabbitmq 实例
 *   从 viper 读取 queue.<name>.addrs / exchange / exchangeType / durable / autoAck
 *@param name 配置 key
 *@return mq.PubSub
 *@return error
 */
func NewRabbitMQ(name string) (mq.PubSub, error) {
	logType := "NewRabbitMQ"
	p := &ConnParams{
		Addr:         viper.GetString(fmt.Sprintf("queue.%s.addrs", name)),
		Exchange:     viper.GetString(fmt.Sprintf("queue.%s.exchange", name)),
		ExchangeType: viper.GetString(fmt.Sprintf("queue.%s.exchangeType", name)),
		Durable:      viper.GetBool(fmt.Sprintf("queue.%s.durable", name)),
		AutoAck:      viper.GetBool(fmt.Sprintf("queue.%s.autoAck", name)),
	}
	if p.Addr == "" {
		return nil, fmt.Errorf("rabbitmq addr empty, name:%s", name)
	}
	if p.ExchangeType == "" {
		p.ExchangeType = "direct"
	}
	conn, err := amqp.Dial(p.Addr)
	if err != nil {
		logrus.Errorf("logType:%s, err:%s, name:%s, addr:%s",
			logType, err.Error(), name, p.Addr)
		return nil, err
	}
	return &rabbitmqQueue{
		params: p,
		conn:   conn,
	}, nil
}

// rabbitmqMsg rabbitmq 消息包装
type rabbitmqMsg struct {
	delivery amqp.Delivery
}

// 编译期类型断言: 确保 rabbitmqMsg 实现了 mq.ConsumerMsg
var _ mq.ConsumerMsg = (*rabbitmqMsg)(nil)

/*GetBody
 *@Description: 获取消息体
 *@receiver m
 *@return []byte
 */
func (m *rabbitmqMsg) GetBody() []byte {
	return m.delivery.Body
}

/*GetID
 *@Description: 获取消息 ID(即 publish 时写入的 reqID)
 *@receiver m
 *@return string
 */
func (m *rabbitmqMsg) GetID() string {
	return m.delivery.MessageId
}

/*Ack
 *@Description: 手动 ack 单条消息
 *@receiver m
 *@return error
 */
func (m *rabbitmqMsg) Ack() error {
	return m.delivery.Ack(false)
}
