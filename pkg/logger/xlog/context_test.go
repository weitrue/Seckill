/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2022/12/25 11:18 AM
 * Description:
 **/

package xlog

import (
	"context"
	"testing"

	"github.com/weitrue/Seckill/pkg/logger/meta"
	xzap2 "github.com/weitrue/Seckill/pkg/logger/xzap"
)

var log, _ = xzap2.SetUp(xzap2.Config{
	ServiceName: "Logger",
	Mode:        "console",
	Path:        "test",
	Level:       "info",
	Compress:    false,
	KeepDays:    7,
})

func TestContext(t *testing.T) {
	ctx := context.Background()
	l := WithContext(ctx)

	ll := ToContext(ctx, log)

	l.Info("test", meta.NewField("test", "test"))
	t.Log(l)
	t.Log(ll)
}
