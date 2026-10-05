/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2022/12/19 3:15 下午
 * Description:
 **/

package examples

import (
	"context"
	"fmt"
	"testing"
	"time"

	"golang.org/x/sync/errgroup"
)

func TestName(t *testing.T) {
	eg, ctx := errgroup.WithContext(context.Background())
	eg.Go(func() error {
		time.Sleep(time.Second * 2)
		fmt.Println("eg.Go")
		return nil
	})

	newCtx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	select {
	case <-newCtx.Done():
		fmt.Println("newCtx.Done")
	case <-ctx.Done():
		fmt.Println("ctx.Done")
	}
}
