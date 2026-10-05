/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2022/12/19 4:59 下午
 * Description:
 **/

package examples

import (
	"testing"
)

type SafeMap[K comparable, V any] struct {
	values map[K]V
}

func TestSafeMap(t *testing.T) {
}
