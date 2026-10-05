/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/9/24 11:26 上午
 * Description:
 **/

package factory

import (
	"testing"

	"github.com/weitrue/Seckill/infrastructure/mq"
)

// TestBuiltinDriversRegistered 验证内置驱动都已注册
func TestBuiltinDriversRegistered(t *testing.T) {
	for _, driver := range []string{DriverMemory, DriverKafka, DriverRabbitMQ} {
		if NewFactory(driver) == nil {
			t.Errorf("builtin driver %q not registered", driver)
		}
	}
}

// TestNewFactoryUnknown 验证未注册的 driver 返回 nil
func TestNewFactoryUnknown(t *testing.T) {
	if NewFactory("unknown-driver") != nil {
		t.Error("expect nil for unknown driver")
	}
}

// TestRegisterNilPanic 验证注册 nil 工厂会 panic
func TestRegisterNilPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expect panic when registering nil factory")
		}
	}()
	Register("nil-driver", nil)
}

// TestRegisterDuplicatePanic 验证重复注册会 panic
func TestRegisterDuplicatePanic(t *testing.T) {
	driver := "dup-driver"
	Register(driver, Func(func(name string) (mq.MQ, error) { return nil, nil }))
	defer func() {
		if r := recover(); r == nil {
			t.Error("expect panic when registering duplicate driver")
		}
		// 清理,避免影响其他测试
		driversMu.Lock()
		delete(drivers, driver)
		driversMu.Unlock()
	}()
	Register(driver, Func(func(name string) (mq.MQ, error) { return nil, nil }))
}

// TestNewQueueDriverMismatch 验证 pub/sub 驱动不能被 NewQueue 拿到
func TestNewQueueDriverMismatch(t *testing.T) {
	// kafka / rabbitmq 的构造函数需要网络,所以这里只能用一个"构造成功但类型不匹配"的替身驱动
	stubDriver := "stub-pubsub"
	Register(stubDriver, Func(func(name string) (mq.MQ, error) {
		return &stubPubSub{}, nil
	}))
	defer func() {
		driversMu.Lock()
		delete(drivers, stubDriver)
		driversMu.Unlock()
	}()

	if _, err := NewQueue(stubDriver, ""); err == nil {
		t.Error("expect error when driver does not implement mq.Queue")
	}
	if _, err := NewPubSub(stubDriver, ""); err != nil {
		t.Errorf("expect no error for pubsub driver, got %v", err)
	}
}

// TestNewQueueUnknownDriver 验证未注册 driver 调用辅助函数会报错
func TestNewQueueUnknownDriver(t *testing.T) {
	if _, err := NewQueue("not-exist", ""); err == nil {
		t.Error("expect error for unknown driver")
	}
	if _, err := NewPubSub("not-exist", ""); err == nil {
		t.Error("expect error for unknown driver")
	}
}

// stubPubSub 是只实现 mq.PubSub 的替身,用于测试类型断言
type stubPubSub struct{}

func (s *stubPubSub) Close() error                              { return nil }
func (s *stubPubSub) Publish(topic string, body []byte, reqID string) error {
	return nil
}
func (s *stubPubSub) Subscribe(topic, group string, fn mq.MsgCb) error {
	return nil
}
