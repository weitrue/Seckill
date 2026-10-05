package crons

import (
	"context"

	"offer/pkg/logger/meta"
	"offer/pkg/logger/xlog"

	"github.com/robfig/cron/v3"
)

type Logger struct{}

// NewLogger 新建日志记录器
func NewLogger() cron.Logger {
	return &Logger{}
}

func (l Logger) Info(msg string, keysAndValues ...interface{}) {
	metas := make([]meta.Field, 0)
	for i := 0; i+1 < len(keysAndValues); {
		metas = append(metas, meta.NewField(keysAndValues[i].(string), keysAndValues[i+1]))
		i = i + 2
	}

	xlog.WithContext(context.Background()).Info(msg, metas...)
}

func (l Logger) Error(err error, msg string, keysAndValues ...interface{}) {
	metas := make([]meta.Field, 0)
	for i := 0; i+1 < len(keysAndValues); {
		metas = append(metas, meta.NewField(keysAndValues[i].(string), keysAndValues[i+1]))
		i = i + 2
	}

	xlog.WithContext(context.Background()).Error(msg, err, metas...)
}
