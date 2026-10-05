/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2022/12/21 11:04 PM
 * Description:
 **/

package xlog

import (
	"context"
	"fmt"

	"github.com/weitrue/Seckill/pkg/logger"
	"github.com/weitrue/Seckill/pkg/logger/meta"
	"github.com/weitrue/Seckill/pkg/logger/xzap"

	"go.uber.org/zap/zapcore"
)

var ctxMarkedKey = &ctxMarker{}

type ctxMarker struct{}

// CtxLogger 日志上下文记录器
type CtxLogger struct {
	logger *xzap.ZapLogger
	fields []meta.Field
}

// WithContext 获取当前上下文日志记录器
func WithContext(ctx context.Context) logger.Logger {
	return newContextLogger(ctx)
}

// ToContext 返回新的上下文并添加日志到上下文用于提取
func ToContext(ctx context.Context, logger *xzap.ZapLogger) context.Context {
	l, ok := ctx.Value(ctxMarkedKey).(*CtxLogger)
	if ok {
		return ctx
	}

	l = &CtxLogger{
		logger: logger,
	}

	return context.WithValue(ctx, ctxMarkedKey, l)
}

// newContextLogger 获取当前上下文日志记录器
func newContextLogger(ctx context.Context) *CtxLogger {
	l, ok := ctx.Value(ctxMarkedKey).(*CtxLogger)
	if !ok || l == nil {
		fmt.Println(xzap.Logger)
		l = &CtxLogger{
			logger: xzap.Logger,
		}
	}

	return l
}

// AddFields 添加zap Field 到日志中
func AddFields(ctx context.Context, fields ...meta.Field) {
	l := newContextLogger(ctx)
	l.fields = append(l.fields, fields...)
}

// extract 提取context log中fields
func (l *CtxLogger) extract() *xzap.ZapLogger {
	return l.logger.With(l.fields...)
}

// Debug debug log
func (l *CtxLogger) Debug(msg string, data ...meta.Field) {
	l.extract().Debug(msg, data...)
}

// Info log
func (l *CtxLogger) Info(msg string, data ...meta.Field) {
	l.extract().Info(msg, data...)
}

// Warn log
func (l *CtxLogger) Warn(msg string, err error, data ...meta.Field) {
	l.extract().Warn(msg, err, data...)
}

// Error error log
func (l *CtxLogger) Error(msg string, err error, data ...meta.Field) {
	l.extract().Error(msg, err, data...)
}

// Print 调用 zap.Logger write
func (l *CtxLogger) Print(msg string, level zapcore.Level, err error, data ...meta.Field) {
	l.extract().Print(msg, level, err, data...)
}

// Panic 调用 zap.Logger Panic
func (l *CtxLogger) Panic(msg string, err error, data ...meta.Field) {
	l.extract().Panic(msg, err, data...)
}

// Debugf debug
func (l *CtxLogger) Debugf(msg, format string, data ...any) {
	l.extract().Debug(msg, nil, meta.NewField("content", fmt.Sprintf(format, data...)))
}

// Infof info
func (l *CtxLogger) Infof(msg, format string, data ...any) {
	l.extract().Info(msg, nil, meta.NewField("content", fmt.Sprintf(format, data...)))
}

// Warnf warn
func (l *CtxLogger) Warnf(msg, format string, err error, data ...any) {
	l.extract().Warn(msg, err, meta.NewField("content", fmt.Sprintf(format, data...)))
}

// Errorf error
func (l *CtxLogger) Errorf(msg, format string, err error, data ...any) {
	l.extract().Error(msg, err, meta.NewField("content", fmt.Sprintf(format, data...)))
}

// Panicf panic
func (l *CtxLogger) Panicf(msg, format string, err error, data ...any) {
	l.extract().Panic(msg, err, meta.NewField("content", fmt.Sprintf(format, data...)))
}
