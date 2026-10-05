package es

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	pageSize := int64(100)

	for page := int64(1); ; page++ {
		total := int64(371)

		fmt.Println(page)
		if page*pageSize >= total {
			break
		}
	}
}
