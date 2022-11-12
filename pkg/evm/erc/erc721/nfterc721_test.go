package erc721

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var cases = []string{
	//"https://speedy-nodes-nyc.moralis.io/911bb3ed10c2befdc9c7061c/eth/mainnet",
	"https://speedy-nodes-nyc.moralis.io/911bb3ed10c2befdc9c7061c/eth/rinkeby",
}

func TestNftErc721_GetItemOwner(t *testing.T) {
	for _, c := range cases {
		t.Run(c, func(t *testing.T) {
			erc721, err := NewNftErc721(c)
			assert.Nil(t, err)
			address, err := erc721.GetItemOwner("0xe169f5a086858a0e0398726a9a42cc8077C0c90d", "2101")
			assert.Nil(t, err)
			t.Log(address)
		})
	}
}

func TestA(t *testing.T) {
	arra := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t.Log(append(arra[:10], arra[11:]...))
}
