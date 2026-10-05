package xzap

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/weitrue/Seckill/pkg/tracespec"
)

const (
	mockTraceID = "mock-trace-id"
	mockSpanID  = "mock-span-id"
)

var mock tracespec.Trace = new(mockTrace)

func TestWithContext(t *testing.T) {
	SetUp(testLoggerConf)
	ctxLogger := WithContext(context.WithValue(context.Background(), tracespec.TracingKey, mock))
	t.Log(ctxLogger.ctx.Value("X-Trace"))
	assert.NotNil(t, ctxLogger.logger)
	assert.Nil(t, ctxLogger.fields)
}

type mockTrace struct{}

func (t mockTrace) TraceId() string {
	return mockTraceID
}

func (t mockTrace) SpanId() string {
	return mockSpanID
}

func (t mockTrace) Finish() {
}

func (t mockTrace) Fork(ctx context.Context, serviceName, operationName string) (context.Context, tracespec.Trace) {
	return nil, nil
}

func (t mockTrace) Follow(ctx context.Context, serviceName, operationName string) (context.Context, tracespec.Trace) {
	return nil, nil
}

func (t mockTrace) Visit(fn func(key, val string) bool) {
}
