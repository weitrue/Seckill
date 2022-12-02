package snaili

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContractSnail_TokenUrl(t *testing.T) {
	snail, err := NewContractSnail("https://bsc-dataseed1.binance.org") // "http://3.0.249.9:8545 // https://bsc-dataseed1.binance.org
	if err != nil {
		t.Log(err)
	}

	tokenUrl, err := snail.TokenUrl("1015")
	assert.Nil(t, err)
	t.Log(tokenUrl)

}
