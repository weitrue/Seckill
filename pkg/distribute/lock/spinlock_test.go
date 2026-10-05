/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2023/2/20 8:04 PM
 * Description:
 **/

package lock

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestLock(t *testing.T) {
	wg := sync.WaitGroup{}
	l := Lock{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			l.l.Lock()
			defer l.l.Unlock()
			fmt.Println(fmt.Sprintf("lock num: %d", num))
			time.Sleep(time.Second)
		}(i)
	}
	wg.Wait()
	p := sync.Pool{
		New: func() any {
			return nil
		},
	}
	p.Get()

}
