/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/13 下午6:35
 * Description:
 **/

package kafka

import (
	"fmt"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/weitrue/Seckill/infrastructure/worker"
)

type kafkaQueue struct {
	params  *ConnParams
	brokers []string
	pub     sarama.SyncProducer
	sub     *cluster.Consumer
}

type ConnParams struct {
	Addr     string // separate by "," for more than one servers case, such as kafka brokers
	NeedPub  bool   // for pub
	NeedSub  bool   // for sub
	AutoAck  bool   // for sub
	Exchange string
}

func (m *kafkaQueue) Publish(topic string, body []byte, reqID string) error {
	k := &sarama.ProducerMessage{
		Topic:     topic,
		Key:       sarama.StringEncoder(body),
		Value:     sarama.ByteEncoder(reqID),
		Timestamp: time.Now(),
	}
	_, _, err := m.pub.SendMessage(k)

	return err
}

func (m *kafkaQueue) Subscribe(topic, queueName string, fn MsgCb) error {
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Group.Return.Notifications = true
	consumer, err := cluster.NewConsumer(m.brokers, m.params.Exchange, []string{topic}, config)
	if err != nil {
		return err
	}
	m.sub = consumer

	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				logrus.Error("consumer err", "err", err)
			case n := <-consumer.Notifications():
				logrus.Info("consumer rebalanced", "notify", n)
			case msg, ok := <-consumer.Messages():
				if !ok {
					logrus.Info("msg queue closed", "brokers", strings.Join(m.brokers, ","),
						"topic", topic, "group", m.params.Exchange)
					return
				}
				k := &kafkaMsg{body: msg, handle: m}
				if err := fn(k); err == nil {
					if m.params.AutoAck {
						_ = k.Ack()
					}
				} else {
					logrus.Error("msg cb err", "key", k.GetID(), "err", err)
				}
			}
		}
	}()

	return nil
}

func (m *kafkaQueue) Produce(task worker.Task) error {
	panic("implement me")
}

func (m *kafkaQueue) Consume() (worker.Task, error) {
	panic("implement me")
}

func (m *kafkaQueue) Close() error {
	if m.pub != nil {
		err := m.pub.Close()
		if err != nil {
			return err
		}
	}
	if m.sub != nil {
		err := m.sub.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func NewKafkaQueue(name string) (*kafkaQueue, error) {
	kafka := new(kafkaQueue)
	p := &ConnParams{
		Addr:     viper.GetString(fmt.Sprintf("queue.%s.addrs", name)),
		NeedPub:  viper.GetBool(fmt.Sprintf("queue.%s.needPub", name)),
		NeedSub:  viper.GetBool(fmt.Sprintf("queue.%s.needSub", name)),
		AutoAck:  viper.GetBool(fmt.Sprintf("queue.%s.autoAck", name)),
		Exchange: viper.GetString(fmt.Sprintf("queue.%s.exchange", name)),
	}
	kafka.params = p
	kafka.brokers = strings.Split(p.Addr, ",")
	if p.NeedPub {
		kc := sarama.NewConfig()
		kc.Producer.RequiredAcks = sarama.WaitForAll // Wait for all in-sync replicas to ack the message
		kc.Producer.Retry.Max = 10                   // Retry up to 10 times to produce the message
		kc.Producer.Return.Successes = true
		pub, err := sarama.NewSyncProducer(kafka.brokers, kc)
		if err != nil {
			return nil, err
		}

		kafka.pub = pub
	}

	return kafka, nil
}

type KConsumerMsg interface {
	GetBody() []byte
	GetID() string
	Ack() error
}

type MsgCb func(m KConsumerMsg) error

type kafkaMsg struct {
	body   *sarama.ConsumerMessage
	handle *kafkaQueue
}

func (m *kafkaMsg) GetBody() []byte {
	return m.body.Value
}

func (m *kafkaMsg) GetID() string {
	return string(m.body.Key)
}

func (m *kafkaMsg) Ack() error {
	m.handle.sub.MarkOffset(m.body, "")
	return nil
}
