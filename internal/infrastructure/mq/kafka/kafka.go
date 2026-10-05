/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/13 下午6:35
 * Description: kafka 实现 mq.PubSub(Publish/Subscribe []byte + Close)
 **/

package kafka

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/weitrue/Seckill/internal/infrastructure/mq"
)

// kafkaQueue kafka pub/sub 实现
type kafkaQueue struct {
	params  *ConnParams
	brokers []string
	pub     sarama.SyncProducer
	sub     *cluster.Consumer
}

// ConnParams kafka 连接参数
type ConnParams struct {
	Addr     string // 多 broker 用 "," 分隔
	NeedPub  bool   // 是否启用生产者
	NeedSub  bool   // 是否启用消费者(由 Subscribe 实际触发)
	AutoAck  bool   // 消费回调成功后是否自动 ack
	Exchange string // 消费组名
}

// 编译期类型断言: 确保 kafkaQueue 实现了 mq.PubSub
var _ mq.PubSub = (*kafkaQueue)(nil)

/*Publish
 *@Description: 向 kafka 发布消息
 *@receiver m
 *@param topic 主题
 *@param body  消息体
 *@param reqID 请求 ID,用于链路追踪
 *@return error
 */
func (m *kafkaQueue) Publish(topic string, body []byte, reqID string) error {
	if m.pub == nil {
		return errors.New("kafka producer not initialized")
	}
	k := &sarama.ProducerMessage{
		Topic:     topic,
		Key:       sarama.StringEncoder(reqID),
		Value:     sarama.ByteEncoder(body),
		Timestamp: time.Now(),
	}
	_, _, err := m.pub.SendMessage(k)

	return err
}

/*Subscribe
 *@Description: 订阅 kafka 主题
 *@receiver m
 *@param topic 主题
 *@param group 消费组(为空则使用 ConnParams.Exchange)
 *@param fn    消息回调
 *@return error
 */
func (m *kafkaQueue) Subscribe(topic, group string, fn mq.MsgCb) error {
	logType := "kafkaQueueSubscribe"
	if group == "" {
		group = m.params.Exchange
	}
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Group.Return.Notifications = true
	consumer, err := cluster.NewConsumer(m.brokers, group, []string{topic}, config)
	if err != nil {
		logrus.Errorf("logType:%s, err:%s, brokers:%s, topic:%s, group:%s",
			logType, err.Error(), strings.Join(m.brokers, ","), topic, group)
		return err
	}
	m.sub = consumer

	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				if err != nil {
					logrus.Errorf("logType:%s, err:%s, topic:%s, group:%s",
						logType, err.Error(), topic, group)
				}
			case n := <-consumer.Notifications():
				logrus.Infof("logType:%s, notify:%+v, topic:%s, group:%s",
					logType, n, topic, group)
			case msg, ok := <-consumer.Messages():
				if !ok {
					logrus.Infof("logType:%s, msg:queue closed, brokers:%s, topic:%s, group:%s",
						logType, strings.Join(m.brokers, ","), topic, group)
					return
				}
				k := &kafkaMsg{body: msg, handle: m}
				if cbErr := fn(k); cbErr == nil {
					if m.params.AutoAck {
						_ = k.Ack()
					}
				} else {
					logrus.Errorf("logType:%s, err:%s, key:%s, topic:%s, group:%s",
						logType, cbErr.Error(), k.GetID(), topic, group)
				}
			}
		}
	}()

	return nil
}

/*Close
 *@Description: 关闭 kafka 生产者和消费者
 *@receiver m
 *@return error
 */
func (m *kafkaQueue) Close() error {
	if m.pub != nil {
		if err := m.pub.Close(); err != nil {
			return err
		}
	}
	if m.sub != nil {
		if err := m.sub.Close(); err != nil {
			return err
		}
	}
	return nil
}

/*NewKafkaQueue
 *@Description: 创建 kafka 队列实例
 *   从 viper 读取 queue.<name>.addrs / needPub / needSub / autoAck / exchange
 *@param name 配置 key
 *@return mq.PubSub
 *@return error
 */
func NewKafkaQueue(name string) (mq.PubSub, error) {
	kq := new(kafkaQueue)
	p := &ConnParams{
		Addr:     viper.GetString(fmt.Sprintf("queue.%s.addrs", name)),
		NeedPub:  viper.GetBool(fmt.Sprintf("queue.%s.needPub", name)),
		NeedSub:  viper.GetBool(fmt.Sprintf("queue.%s.needSub", name)),
		AutoAck:  viper.GetBool(fmt.Sprintf("queue.%s.autoAck", name)),
		Exchange: viper.GetString(fmt.Sprintf("queue.%s.exchange", name)),
	}
	kq.params = p
	kq.brokers = strings.Split(p.Addr, ",")
	if p.NeedPub {
		kc := sarama.NewConfig()
		kc.Producer.RequiredAcks = sarama.WaitForAll // 等所有 ISR ack
		kc.Producer.Retry.Max = 10                   // 最多重试 10 次
		kc.Producer.Return.Successes = true
		pub, err := sarama.NewSyncProducer(kq.brokers, kc)
		if err != nil {
			return nil, err
		}

		kq.pub = pub
	}

	return kq, nil
}

// kafkaMsg kafka 消息包装
type kafkaMsg struct {
	body   *sarama.ConsumerMessage
	handle *kafkaQueue
}

// 编译期类型断言: 确保 kafkaMsg 实现了 mq.ConsumerMsg
var _ mq.ConsumerMsg = (*kafkaMsg)(nil)

/*GetBody
 *@Description: 获取消息体
 *@receiver m
 *@return []byte
 */
func (m *kafkaMsg) GetBody() []byte {
	return m.body.Value
}

/*GetID
 *@Description: 获取消息 ID(即 publish 时的 reqID)
 *@receiver m
 *@return string
 */
func (m *kafkaMsg) GetID() string {
	return string(m.body.Key)
}

/*Ack
 *@Description: 手动 ack
 *@receiver m
 *@return error
 */
func (m *kafkaMsg) Ack() error {
	m.handle.sub.MarkOffset(m.body, "")
	return nil
}
