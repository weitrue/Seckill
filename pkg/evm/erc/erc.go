package erc

import (
	"github.com/pkg/errors"
	"github.com/weitrue/Seckill/pkg/evm/erc/erc721"
)

type Erc interface {
	GetItemOwner(address string, tokenId string) (string, error)
}

type NftErc struct {
	Endpoint string `toml:"endpoint" json:"endpoint"`
	Standard string `toml:"standard" json:"standard"`
}

func NewErc(c *NftErc) (Erc, error) {
	if c == nil || len(c.Standard) == 0 || len(c.Endpoint) == 0 {
		return nil, errors.New("err config")
	}

	switch c.Standard {
	case erc721.ERC721:
		return erc721.NewNftErc721(c.Endpoint)
	default:
		return nil, errors.New("err config")
	}
}
