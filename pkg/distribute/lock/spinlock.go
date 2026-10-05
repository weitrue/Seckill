/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2023/2/16 4:24 PM
 * Description:
 **/

package lock

import (
	"sync"
)

type Lock struct {
	l    sync.Mutex
	Name string
}
