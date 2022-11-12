/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2022/2/8 9:49 下午
 * Description:
 **/

package threading

import (
	"io/ioutil"
	"log"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Placeholder is a placeholder object that can be used globally.
var Placeholder PlaceholderType

type (
	// PlaceholderType represents a placeholder type.
	PlaceholderType = struct{}
)

func TestRunSafe(t *testing.T) {
	log.SetOutput(ioutil.Discard)

	i := 0

	defer func() {
		assert.Equal(t, 1, i)
	}()

	ch := make(chan PlaceholderType)
	go RunSafe(func() {
		defer func() {
			ch <- Placeholder
		}()

		panic("panic")
	})

	<-ch
	i++
}

func TestGoroutineID(t *testing.T) {
	assert.True(t, GoroutineID() > 0)
}

func TestRescue(t *testing.T) {
	var count int32
	assert.NotPanics(t, func() {
		defer Recover(func() {
			atomic.AddInt32(&count, 2)
		}, func() {
			atomic.AddInt32(&count, 3)
		})

		panic("hello")
	})
	assert.Equal(t, int32(5), atomic.LoadInt32(&count))
}
