package contract

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"math"
	"math/big"
	"testing"
)

var erc, _ = NewErc(&Config{
	Endpoint: "http://127.0.0.1:8545",
	Contract: "",
})

func TestHexToAddress(t *testing.T) {
	address := common.HexToAddress("0xe01511d7333A18e969758BBdC9C7f50CcF30160A")

	t.Log(address.Hex())        // 0xe01511d7333A18e969758BBdC9C7f50CcF30160A
	t.Log(address.Hash().Hex()) // 0x000000000000000000000000e01511d7333a18e969758bbdc9c7f50ccf30160a
}

func TestClient_BalanceAt(t *testing.T) {

	balanceAt, err := erc.BalanceAt(context.Background(), "0xbDA5747bFD65F08deb54cb465eB87D40e51B197E")
	assert.Nil(t, err)
	t.Log(balanceAt)
	balance := new(big.Float)
	balance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(balance, big.NewFloat(math.Pow10(18)))
	t.Log(ethValue)

	_, err = erc.GetOwner("1")
	assert.Equal(t, errors.New("unsupported"), err)
}

func TestGenerateKey(t *testing.T) {
	pub, pri, err := erc.GenerateKey()
	assert.Nil(t, err)
	t.Log(pub, pri)
}
