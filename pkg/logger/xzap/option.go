/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2022/12/21 5:21 PM
 * Description:
 **/

package xzap

import (
	"go.uber.org/zap/zapcore"
)

const (
	levelInfo   = "info"
	levelError  = "error"
	levelSevere = "severe"
)

// Option 可选参数
type Option func(*Opt)

type Opt struct {
	level  zapcore.Level
	fields map[string]string
}

// WithField 添加field(s)到日志中
func WithField(key, value string) Option {
	return func(opt *Opt) {
		opt.fields[key] = value
	}
}

func withLogLevel(level string) Option {
	return func(opt *Opt) {
		switch level {
		case levelInfo:
			opt.level = zapcore.InfoLevel
		case levelSevere:
			opt.level = zapcore.WarnLevel
		case levelError:
			opt.level = zapcore.ErrorLevel
		default:
			opt.level = zapcore.DebugLevel
		}
	}
}
