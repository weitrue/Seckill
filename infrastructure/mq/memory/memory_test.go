/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/9/26 11:14 上午
 * Description:
 **/

package memory

import (
	"reflect"
	"testing"

	"github.com/weitrue/Seckill/infrastructure/mq"
	"github.com/weitrue/Seckill/infrastructure/services/local/ratelimiter"
	"github.com/weitrue/Seckill/infrastructure/worker"
)

func TestNewMemoryFactory(t *testing.T) {
	type args struct {
		name string
	}
	var tests []struct {
		name    string
		args    args
		want    mq.Queue
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMemoryMQ(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMemoryMQ() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMemoryMQ() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_memoryQueue_Close(t *testing.T) {
	type fields struct {
		queue ratelimiter.RateLimiter
	}
	var tests []struct {
		name    string
		fields  fields
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mq := &memoryQueue{
				queue: tt.fields.queue,
			}
			if err := mq.Close(); (err != nil) != tt.wantErr {
				t.Errorf("Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_memoryQueue_Consume(t *testing.T) {
	type fields struct {
		queue ratelimiter.RateLimiter
	}
	tests := []struct {
		name    string
		fields  fields
		want    worker.Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mq := &memoryQueue{
				queue: tt.fields.queue,
			}
			got, err := mq.Consume()
			if (err != nil) != tt.wantErr {
				t.Errorf("Consume() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Consume() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_memoryQueue_Produce(t *testing.T) {
	type fields struct {
		queue ratelimiter.RateLimiter
	}
	type args struct {
		task worker.Task
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mq := &memoryQueue{
				queue: tt.fields.queue,
			}
			if err := mq.Produce(tt.args.task); (err != nil) != tt.wantErr {
				t.Errorf("Produce() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_memoryQueue_Publish(t *testing.T) {
	type fields struct {
		queue ratelimiter.RateLimiter
	}
	type args struct {
		topic string
		body  []byte
		reqID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mq := &memoryQueue{
				queue: tt.fields.queue,
			}
			if err := mq.Publish(tt.args.topic, tt.args.body, tt.args.reqID); (err != nil) != tt.wantErr {
				t.Errorf("Publish() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// func Test_memoryQueue_Subscribe(t *testing.T) {
// 	type fields struct {
// 		queue ratelimiter.RateLimiter
// 	}
// 	type args struct {
// 		topic     string
// 		queueName string
// 		fn        mq.MsgCb
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mq := &memoryQueue{
// 				queue: tt.fields.queue,
// 			}
// 			if err := mq.Subscribe(tt.args.topic, tt.args.queueName, tt.args.fn); (err != nil) != tt.wantErr {
// 				t.Errorf("Subscribe() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
