package mock

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

var (
	MockChainID = big.NewInt(167001)
)

type EthClient struct {
}

func (c *EthClient) ChainID(ctx context.Context) (*big.Int, error) {
	return MockChainID, nil
}

func (c *EthClient) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	return &types.Header{
		Number: number,
	}, nil
}
