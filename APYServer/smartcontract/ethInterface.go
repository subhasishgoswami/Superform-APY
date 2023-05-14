package smartcontract

import (
	"context"
	"math/big"

	gethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/subhasishgoswami/superform/smartcontract/VenusERC4626Reinvest"
)

type GethInterface struct {
	Client        ethclient.Client
	URL           string
	SmartContract VenusERC4626Reinvest.VenusERC4626Reinvest
}

func NewGethInterface(rpc string, smartContractAddress string) GethInterface {
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return GethInterface{}
	}

	address := gethCommon.HexToAddress(smartContractAddress)
	instance, err := VenusERC4626Reinvest.NewVenusERC4626Reinvest(address, client)
	if err != nil {
		return GethInterface{}
	}

	return GethInterface{Client: *client, URL: rpc, SmartContract: *instance}
}

func (geth *GethInterface) BlockTimeGeth(block uint64, ctx context.Context) (uint64, error) {
	blockNumber := big.NewInt(int64(block))
	blockGeth, err := geth.Client.BlockByNumber(ctx, blockNumber)
	if err != nil {
		return 0, err
	}
	return blockGeth.Time(), nil
}
