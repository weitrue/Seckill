package contract

import (
	"errors"
	"github.com/weitrue/Seckill/pkg/evm/contract/snaili"
)

type Config struct {
	Endpoint string `toml:"endpoint" json:"endpoint"`
	Contract string `toml:"contract" json:"contract"`
}

type Contract interface {
	GetOwner(tokenId string) (string, error)
}

func NewErc(c *Config) (Contract, error) {
	if c == nil || len(c.Contract) == 0 || len(c.Endpoint) == 0 {
		return nil, errors.New("err config")
	}

	switch c.Contract {
	case snaili.Contract:
		return snaili.NewContractSnail(c.Endpoint)
	default:
		return nil, errors.New("err config")
	}
}
