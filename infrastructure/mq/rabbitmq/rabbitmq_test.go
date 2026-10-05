/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2024/10/05
 * Description:
 **/

package rabbitmq

import (
	"testing"
)

// TestNewRabbitMQEmptyAddr 验证空地址构造会报错
func TestNewRabbitMQEmptyAddr(t *testing.T) {
	// 不设置 viper,addr 读到空
	if _, err := NewRabbitMQ("not-configured"); err == nil {
		t.Error("expect error when addr is empty")
	}
}
