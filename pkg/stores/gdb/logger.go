package gdb

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-stack/stack"
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
)

const (
	traceInfo  = "%s [%.3fms] [rows:%v] %s"
	traceWarn  = "%s %s [%.3fms] [rows:%v] %s"
	traceError = "%s %s [%.3fms] [rows:%v] %s"
)

// Logger gorm 日志记录器
//
//	走 zap.L() 全局 logger,业务侧通过 zap.ReplaceGlobals() 配置后生效
//	保留 ctx 入参以匹配 gorm/logger.Interface,后续可从 ctx 携带 trace 字段
type Logger struct {
	LogLevel      logger.LogLevel
	SlowThreshold time.Duration
}

// NewLogger 新建日志记录器
func NewLogger(logLevel logger.LogLevel, slowThreshold time.Duration) *Logger {
	return &Logger{LogLevel: logLevel, SlowThreshold: slowThreshold}
}

// LogMode 设置日志记录模式
func (l *Logger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

// Info Info日志记录
func (l *Logger) Info(ctx context.Context, msg string, data ...interface{}) {
	_ = ctx
	if l.LogLevel >= logger.Info {
		zap.L().Info("DB", zap.String("content", fmt.Sprintf(msg, data...)))
	}
}

// Warn Warn日志记录
func (l *Logger) Warn(ctx context.Context, msg string, data ...interface{}) {
	_ = ctx
	if l.LogLevel >= logger.Warn {
		zap.L().Warn("DB", zap.String("content", fmt.Sprintf(msg, data...)))
	}
}

// Error Error日志记录
func (l *Logger) Error(ctx context.Context, msg string, data ...interface{}) {
	_ = ctx
	if l.LogLevel >= logger.Error {
		zap.L().Error("DB", zap.String("content", fmt.Sprintf(msg, data...)))
	}
}

// Trace Trace日志记录
func (l *Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel > logger.Silent {
		elapsed := time.Since(begin)
		switch {
		case err != nil && l.LogLevel >= logger.Error:
			sql, rows := fc()
			if rows == -1 {
				l.Error(ctx, traceError, FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				l.Error(ctx, traceError, FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= logger.Warn:
			sql, rows := fc()
			slowLog := fmt.Sprintf("Slow SQL Greater Than %v", l.SlowThreshold)
			if rows == -1 {
				l.Warn(ctx, traceWarn, FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				l.Warn(ctx, traceWarn, FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		case l.LogLevel == logger.Info:
			sql, rows := fc()
			if rows == -1 {
				l.Info(ctx, traceInfo, FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				l.Info(ctx, traceInfo, FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		}
	}
}

// FileWithLineNum 获取调用堆栈信息
func FileWithLineNum() string {
	cs := stack.Trace().TrimBelow(stack.Caller(2)).TrimRuntime()

	for _, c := range cs {
		s := fmt.Sprintf("%+v", c)
		if !strings.Contains(s, "gorm.io/gorm") {
			return s
		}
	}

	return ""
}
