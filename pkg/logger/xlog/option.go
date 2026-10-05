/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2022/12/22 9:13 AM
 * Description:
 **/

package xlog

import (
	"context"
	"time"

	"github.com/weitrue/Seckill/pkg/logger/meta"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/weitrue/Seckill/pkg/errcode"
)

// Option 可选参数
type Option func(*option)

type option struct {
	shouldLog           Decider
	codeFunc            ErrorToCode
	levelFunc           CodeToLevel
	durationFunc        DurationToField
	messageFunc         MessageProducer
	recoveryHandlerFunc RecoveryHandlerFuncContext
}

func WithDecider(decider Decider) Option {
	return func(opt *option) {
		opt.shouldLog = decider
	}
}

// Decider 决策器 定义抑制拦截器日志的规则
type Decider func(methodName string, err error) bool

func WithErrorToCode(err ErrorToCode) Option {
	return func(opt *option) {
		opt.codeFunc = err
	}
}

// ErrorToCode 定义error 映射 code
type ErrorToCode func(err error) uint32

func WithCodeToLevel(code CodeToLevel) Option {
	return func(opt *option) {
		opt.levelFunc = code
	}
}

// CodeToLevel rpc返回码与zap日志级别映射
type CodeToLevel func(code uint32) zapcore.Level

func WithDurationToField(duration DurationToField) Option {
	return func(opt *option) {
		opt.durationFunc = duration
	}
}

// DurationToField 生成日志持续时间
type DurationToField func(duration time.Duration) meta.Field

func WithMessageProducer(producer MessageProducer) Option {
	return func(opt *option) {
		opt.messageFunc = producer
	}
}

// MessageProducer 生成日志消息
type MessageProducer func(ctx context.Context, msg string, level zapcore.Level, err error, data ...meta.Field)

func WithRecoveryHandlerFuncContext(recovery RecoveryHandlerFuncContext) Option {
	return func(opt *option) {
		opt.recoveryHandlerFunc = recovery
	}
}

// RecoveryHandlerFuncContext 将recovery以error形式返回 上下文可以用于提取请求的元数据和上下文
type RecoveryHandlerFuncContext func(ctx context.Context, p any) (err error)

var defaultOptions = &option{
	levelFunc:           DefaultCodeToLevel,
	shouldLog:           DefaultDeciderMethod,
	codeFunc:            DefaultErrorToCode,
	durationFunc:        DefaultDurationToField,
	messageFunc:         DefaultMessageProducer,
	recoveryHandlerFunc: DefaultRecoveryHandlerFunc,
}

// DefaultCodeToLevel 根据RPC服务端返回码返回zap日志级别
func DefaultCodeToLevel(code uint32) zapcore.Level {
	switch codes.Code(code) {
	case codes.OK:
		return zap.InfoLevel
	case codes.Canceled:
		return zap.InfoLevel
	case codes.Unknown:
		return zap.ErrorLevel
	case codes.InvalidArgument:
		return zap.InfoLevel
	case codes.DeadlineExceeded:
		return zap.WarnLevel
	case codes.NotFound:
		return zap.InfoLevel
	case codes.AlreadyExists:
		return zap.InfoLevel
	case codes.PermissionDenied:
		return zap.WarnLevel
	case codes.Unauthenticated:
		return zap.InfoLevel // unauthenticated requests can happen
	case codes.ResourceExhausted:
		return zap.WarnLevel
	case codes.FailedPrecondition:
		return zap.WarnLevel
	case codes.Aborted:
		return zap.WarnLevel
	case codes.OutOfRange:
		return zap.WarnLevel
	case codes.Unimplemented:
		return zap.ErrorLevel
	case codes.Internal:
		return zap.ErrorLevel
	case codes.Unavailable:
		return zap.WarnLevel
	case codes.DataLoss:
		return zap.ErrorLevel
	default:
		if code >= 7000 {
			return zap.InfoLevel
		}

		return zap.ErrorLevel
	}
}

// DefaultClientCodeToLevel 根据RPC客户端返回码返回zap日志级别
func DefaultClientCodeToLevel(code uint32) zapcore.Level {
	switch codes.Code(code) {
	case codes.OK:
		return zap.DebugLevel
	case codes.Canceled:
		return zap.DebugLevel
	case codes.Unknown:
		return zap.InfoLevel
	case codes.InvalidArgument:
		return zap.DebugLevel
	case codes.DeadlineExceeded:
		return zap.InfoLevel
	case codes.NotFound:
		return zap.DebugLevel
	case codes.AlreadyExists:
		return zap.DebugLevel
	case codes.PermissionDenied:
		return zap.InfoLevel
	case codes.Unauthenticated:
		return zap.InfoLevel // unauthenticated requests can happen
	case codes.ResourceExhausted:
		return zap.DebugLevel
	case codes.FailedPrecondition:
		return zap.DebugLevel
	case codes.Aborted:
		return zap.DebugLevel
	case codes.OutOfRange:
		return zap.DebugLevel
	case codes.Unimplemented:
		return zap.WarnLevel
	case codes.Internal:
		return zap.WarnLevel
	case codes.Unavailable:
		return zap.WarnLevel
	case codes.DataLoss:
		return zap.WarnLevel
	default:
		if code >= 7000 {
			return zap.InfoLevel
		}

		return zap.InfoLevel
	}
}

// DefaultDeciderMethod 决策器是否记录日志的默认实现，默认是记录日志
func DefaultDeciderMethod(methodName string, err error) bool {
	return true
}

// DefaultErrorToCode error映射code
func DefaultErrorToCode(err error) uint32 {
	if err == nil {
		return uint32(codes.OK)
	}

	switch e := err.(type) {
	case interface{ GRPCStatus() *status.Status }:
		return uint32(status.Code(err))
	case *errcode.Err:
		return e.Code()
	default:
		return uint32(codes.Unknown)
	}
}

// DefaultDurationToField 请求持续时间转换为Zap字段
var DefaultDurationToField = DurationToTimeMillisField

// DurationToTimeMillisField 持续时间转换为毫秒并使用key[time_ms]
func DurationToTimeMillisField(duration time.Duration) meta.Field {
	return meta.NewField("grpc.duration", durationToMilliseconds(duration))
}

// durationToMilliseconds 时间转换为毫秒级别
func durationToMilliseconds(duration time.Duration) float32 {
	return float32(duration.Nanoseconds()/1000) / 1000
}

// DefaultMessageProducer 写日志
func DefaultMessageProducer(ctx context.Context, msg string, level zapcore.Level, err error, data ...meta.Field) {
	// 重新从上下文提取日志
	WithContext(ctx).Print(msg, level, err, data...)
}

// DefaultRecoveryHandlerFunc 恐慌默认处理
func DefaultRecoveryHandlerFunc(ctx context.Context, p any) (err error) {
	err = status.Errorf(codes.Internal, "%v", p)
	WithContext(ctx).Panic("panic", err)

	return err
}
