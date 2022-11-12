/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2022/2/8 9:41 下午
 * Description:
 **/

package threading

import (
	"bytes"
	"runtime"
	"strconv"
)

// GoSafe runs the given fn using another goroutine, recovers if fn panics.
func GoSafe(fn func()) {
	go RunSafe(fn)
}

// RunSafe runs the given fn, recovers if fn panics.
func RunSafe(fn func()) {
	defer Recover()

	fn()
}

// GoroutineID Get returns the id of the current goroutine.
func GoroutineID() int64 {
	buf := make([]byte, 64)
	buf = buf[:runtime.Stack(buf, false)]
	buf = bytes.TrimPrefix(buf, []byte("goroutine "))
	buf = buf[:bytes.IndexByte(buf, ' ')]
	gid, _ := strconv.ParseInt(string(buf), 10, 64)

	return gid
}

// Recover is used with defer to do cleanup on panics.
// Use it like:
//  defer Recover(func() {})
func Recover(cleanups ...func()) {
	for _, cleanup := range cleanups {
		cleanup()
	}

	if p := recover(); p != nil {
		// logx.ErrorStack(p)
		panic("")
	}
}
