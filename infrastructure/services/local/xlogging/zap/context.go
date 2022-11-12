package xzap

import (
	"context"
	"fmt"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/peer"

	"github.com/weitrue/Seckill/infrastructure/services/local/xlogging"
	"github.com/weitrue/Seckill/pkg/jwt"
	"github.com/weitrue/Seckill/pkg/tracespec"
)

var ctxMarkedKey = &ctxMarker{}

const (
	// customMsg 自定义msg
	customMsg = "custom"
)

type ctxMarker struct{}

// CtxLogger 日志上下文记录器
type CtxLogger struct {
	logger *zap.Logger
	fields []zapcore.Field
	ctx    context.Context
}

func newGatewayLoggerCaller(ctx *gin.Context, logger *zap.Logger, start time.Time, timestampFormat string) context.Context {
	var fields []zapcore.Field
	fields = append(fields, zap.String("start_time", start.Format(timestampFormat)))
	if d, ok := ctx.Deadline(); ok {
		fields = append(fields, zap.String("deadline", d.Format(timestampFormat)))
	}

	if p, ok := peer.FromContext(ctx); ok {
		fields = append(fields, zap.String("address", p.Addr.String()))
	}

	tr, ok := ctx.Value(tracespec.TracingKey).(tracespec.Trace)
	if ok {
		fields = append(fields, zap.String("trace", tr.TraceId()))
		fields = append(fields, zap.String("span", tr.SpanId()))
	}

	token, ok := jwt.FromContext(ctx)
	if ok {
		fields = append(fields, zap.Int64("userid", token.UserId))
		fields = append(fields, zap.Int64("companyid", token.CompanyId))
		fields = append(fields, zap.Int64("departmentid", token.DepartmentId))
	}

	return toContext(ctx, logger.With(append(fields, gatewayCallFields(ctx.Request.URL.Path)...)...))
}

func newServerLoggerCaller(ctx context.Context, logger *zap.Logger, methodString string, start time.Time, timestampFormat string) context.Context {
	var fields []zapcore.Field
	fields = append(fields, zap.String("grpc.start_time", start.Format(timestampFormat)))
	if d, ok := ctx.Deadline(); ok {
		fields = append(fields, zap.String("grpc.request.deadline", d.Format(timestampFormat)))
	}

	if p, ok := peer.FromContext(ctx); ok {
		fields = append(fields, zap.String("grpc.address", p.Addr.String()))
	}

	tr, ok := ctx.Value(tracespec.TracingKey).(tracespec.Trace)
	if ok {
		fields = append(fields, zap.String("trace", tr.TraceId()))
		fields = append(fields, zap.String("span", tr.SpanId()))
	}

	token, ok := jwt.FromContext(ctx)
	if ok {
		fields = append(fields, zap.Int64("userid", token.UserId))
		fields = append(fields, zap.Int64("companyid", token.CompanyId))
		fields = append(fields, zap.Int64("departmentid", token.DepartmentId))
	}

	return toContext(ctx, logger.With(append(fields, serverCallFields(methodString)...)...))
}

func newClientLoggerCaller(ctx context.Context, logger *zap.Logger, methodString string, start time.Time, timestampFormat string) context.Context {
	var fields []zapcore.Field
	fields = append(fields, zap.String("grpc.start_time", start.Format(timestampFormat)))
	if d, ok := ctx.Deadline(); ok {
		fields = append(fields, zap.String("grpc.request.deadline", d.Format(timestampFormat)))
	}

	if p, ok := peer.FromContext(ctx); ok {
		fields = append(fields, zap.String("grpc.address", p.Addr.String()))
	}

	tr, ok := ctx.Value(tracespec.TracingKey).(tracespec.Trace)
	if ok {
		fields = append(fields, zap.String("trace", tr.TraceId()))
		fields = append(fields, zap.String("span", tr.SpanId()))
	}

	token, ok := jwt.FromContext(ctx)
	if ok {
		fields = append(fields, zap.Int64("userid", token.UserId))
		fields = append(fields, zap.Int64("companyid", token.CompanyId))
		fields = append(fields, zap.Int64("departmentid", token.DepartmentId))
	}

	return toContext(ctx, logger.With(append(fields, clientLoggerFields(methodString)...)...))
}

func gatewayCallFields(path string) []zapcore.Field {
	return []zapcore.Field{
		GatewaySystemField,
		ServerField,
		zap.String("gateway.path", path),
	}
}

func serverCallFields(methodString string) []zapcore.Field {
	service := path.Dir(methodString)[1:]
	method := path.Base(methodString)
	return []zapcore.Field{
		GrpcSystemField,
		ServerField,
		zap.String("grpc.service", service),
		zap.String("grpc.method", method),
	}
}

func clientLoggerFields(methodString string) []zapcore.Field {
	service := path.Dir(methodString)[1:]
	method := path.Base(methodString)
	return []zapcore.Field{
		GrpcSystemField,
		ClientField,
		zap.String("grpc.service", service),
		zap.String("grpc.method", method),
	}
}

// toContext 返回新的上下文并添加日志到上下文用于提取
func toContext(ctx context.Context, logger *zap.Logger) context.Context {
	l := &CtxLogger{
		logger: logger,
	}

	return context.WithValue(ctx, ctxMarkedKey, l)
}

// WithContext 获取当前上下文日志记录器
func WithContext(ctx context.Context) *CtxLogger {
	l, ok := ctx.Value(ctxMarkedKey).(*CtxLogger)
	if !ok || l == nil {
		return NewContextLogger(ctx)
	}

	l.ctx = ctx

	return l
}

// NewContextLogger 获取当前上下文日志记录器
func NewContextLogger(ctx context.Context) *CtxLogger {
	return &CtxLogger{
		logger: logger,
		ctx:    ctx,
	}
}

// addFields 添加zap Field 到日志中
func addFields(ctx context.Context, fields ...zap.Field) {
	l := WithContext(ctx)
	l.fields = append(l.fields, fields...)
}

// tagsToFields 将上下文中的日志标签转换为zap中的Field
func tagsToFields(ctx context.Context) []zapcore.Field {
	var fields []zapcore.Field
	tags := xlogging.Extract(ctx)
	for k, v := range tags.Values() {
		fields = append(fields, zap.Any(k, v))
	}
	return fields
}

// extract 提取xzap中最新的Logger
func (l *CtxLogger) extract() *zap.Logger {
	fields := tagsToFields(l.ctx)
	fields = append(fields, l.fields...)
	return l.logger.With(fields...)
}

// Debug 调用 zap.Logger Debug
func (l *CtxLogger) Debug(msg string, fields ...zap.Field) {
	l.extract().WithOptions(zap.AddCallerSkip(1)).Debug(msg, fields...)
}

// Info 调用 zap.Logger Info
func (l *CtxLogger) Info(msg string, fields ...zap.Field) {
	l.extract().WithOptions(zap.AddCallerSkip(1)).Info(msg, fields...)
}

// Warn 调用 zap.Logger Warn
func (l *CtxLogger) Warn(msg string, fields ...zap.Field) {
	l.extract().WithOptions(zap.AddCallerSkip(1)).Warn(msg, fields...)
}

// Error 调用 zap.Logger Error
func (l *CtxLogger) Error(msg string, fields ...zap.Field) {
	l.extract().WithOptions(zap.AddCallerSkip(1)).Error(msg, fields...)
}

// Panic 调用 zap.Logger Panic
func (l *CtxLogger) Panic(msg string, fields ...zap.Field) {
	l.extract().WithOptions(zap.AddCallerSkip(1)).Panic(msg, fields...)
}

// Debugf 调用 zap.Logger Debug
func (l *CtxLogger) Debugf(format string, data ...interface{}) {
	l.Debug(customMsg, zap.String("content", fmt.Sprintf(format, data...)))
}

// Infof 调用 zap.Logger Info
func (l *CtxLogger) Infof(format string, data ...interface{}) {
	l.Info(customMsg, zap.String("content", fmt.Sprintf(format, data...)))
}

// Warnf 调用 zap.Logger Warn
func (l *CtxLogger) Warnf(format string, data ...interface{}) {
	l.Warn(customMsg, zap.String("content", fmt.Sprintf(format, data...)))
}

// Errorf 调用 zap.Logger Error
func (l *CtxLogger) Errorf(format string, data ...interface{}) {
	l.Error(customMsg, zap.String("content", fmt.Sprintf(format, data...)))
}

// Panicf 调用 zap.Logger Panic
func (l *CtxLogger) Panicf(format string, data ...interface{}) {
	l.Error(customMsg, zap.String("content", fmt.Sprintf(format, data...)))
}
