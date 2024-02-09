package relayer

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"github.com/taikoxyz/taiko-mono/packages/relayer/bindings/bridge"
)

func Test_IsInSlice(t *testing.T) {
	if IsInSlice("fake", []string{}) {
		t.Fatal()
	}

	if !IsInSlice("real", []string{"real"}) {
		t.Fatal()
	}
}

type mockConfirmer struct {
}

var (
	notFoundTxHash = common.HexToHash("0x123")
	succeedTxHash  = common.HexToHash("0x456")
	failTxHash     = common.HexToHash("0x789")
	blockNum       = 10
)

func (m *mockConfirmer) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	if txHash == notFoundTxHash {
		return nil, ethereum.NotFound
	}

	if txHash == succeedTxHash {
		return &types.Receipt{
			Status:      types.ReceiptStatusSuccessful,
			BlockNumber: new(big.Int).Sub(big.NewInt(int64(blockNum)), big.NewInt(1)),
		}, nil
	}

	return &types.Receipt{
		Status:      types.ReceiptStatusFailed,
		BlockNumber: big.NewInt(1),
	}, nil
}

func (m *mockConfirmer) BlockNumber(ctx context.Context) (uint64, error) {
	return uint64(blockNum), nil
}

func Test_WaitReceipt(t *testing.T) {
	timeoutTicker, cancel := context.WithTimeout(context.Background(), time.Duration(2*time.Second))
	defer cancel()

	tests := []struct {
		name        string
		ctx         context.Context
		txHash      common.Hash
		wantErr     error
		wantReceipt *types.Receipt
	}{
		{
			"success",
			context.Background(),
			succeedTxHash,
			nil,
			&types.Receipt{
				Status:      types.ReceiptStatusSuccessful,
				BlockNumber: new(big.Int).Sub(big.NewInt(int64(blockNum)), big.NewInt(1)),
			},
		},
		{
			"receiptStatusUnsuccessful",
			context.Background(),
			failTxHash,
			fmt.Errorf("transaction reverted, hash: %s", failTxHash),
			nil,
		},
		{
			"ticker timeout",
			timeoutTicker,
			notFoundTxHash,
			errors.New("context deadline exceeded"),
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receipt, err := WaitReceipt(tt.ctx, &mockConfirmer{}, tt.txHash)
			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
			} else {
				assert.Nil(t, err)
			}

			assert.Equal(t, tt.wantReceipt, receipt)
		})
	}
}

func Test_WaitConfirmations(t *testing.T) {
	timeoutTicker, cancel := context.WithTimeout(context.Background(), time.Duration(2*time.Second))
	defer cancel()

	tests := []struct {
		name    string
		ctx     context.Context
		confs   uint64
		txHash  common.Hash
		wantErr error
	}{
		{
			"success",
			context.Background(),
			1,
			succeedTxHash,
			nil,
		},
		{
			"ticker timeout",
			timeoutTicker,
			1,
			notFoundTxHash,
			errors.New("context deadline exceeded"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := WaitConfirmations(tt.ctx, &mockConfirmer{}, tt.confs, tt.txHash)
			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func Test_DecodeMessageSentData(t *testing.T) {
	tests := []struct {
		name               string
		event              *bridge.BridgeMessageSent
		wantEventType      EventType
		wantCanonicalToken CanonicalToken
		wantAmount         *big.Int
		wantError          error
	}{
		{
			"receiveERC20",
			&bridge.BridgeMessageSent{
				Message: bridge.IBridgeMessage{
					// nolint lll
					Data: common.Hex2Bytes("240f6a5f000000000000000000000000000000000000000000000000000000000000008000000000000000000000000015d34aaf54267db7d7c367839aaf71a00a2c6a650000000000000000000000009965507d1a55bcc2695c58ba16fb37d819b0a4dc00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000028c590000000000000000000000006e7cfe3fd749f6e1c53bc1065fc702eb6c22f600000000000000000000000000000000000000000000000000000000000000001200000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000e00000000000000000000000000000000000000000000000000000000000000004544553540000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000095465737445524332300000000000000000000000000000000000000000000000"),
				},
			},
			EventTypeSendERC20,
			CanonicalERC20{
				ChainId:  167001,
				Addr:     common.HexToAddress("0x6e7cfe3fd749F6E1C53BC1065fc702eb6c22F600"),
				Decimals: uint8(18),
				Symbol:   "TEST",
				Name:     "TestERC20",
			},
			big.NewInt(1),
			nil,
		},
		{
			"nilData",
			&bridge.BridgeMessageSent{
				Message: bridge.IBridgeMessage{
					// nolint lll
					Value: big.NewInt(1),
					Data:  common.Hex2Bytes("00"),
				},
			},
			EventTypeSendETH,
			nil,
			big.NewInt(1),
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eventType, canonicalToken, amount, err := DecodeMessageSentData(tt.event)

			assert.Equal(t, tt.wantEventType, eventType)
			assert.Equal(t, tt.wantCanonicalToken, canonicalToken)
			assert.Equal(t, tt.wantAmount, amount)
			assert.Equal(t, tt.wantError, err)
		})
	}
}
