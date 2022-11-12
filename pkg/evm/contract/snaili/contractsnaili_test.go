package snaili

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContractSnail_TokenUrl(t *testing.T) {
	snail, err := NewContractSnail("https://bsc-dataseed2.binance.org/")
	if err != nil {
		t.Log(err)
	}

	tokenUrl, err := snail.TokenUrl("6804")
	assert.Nil(t, err)
	t.Log(tokenUrl)
}
