package transaction

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"

	"github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/pkg/rpc"
)

var (
	ErrUnretryableSubmission = errors.New("unretryable submission error")
	ZeroAddress              common.Address
)

// TxBuilder will build a transaction with the given nonce.
type TxBuilder func(txOpts *bind.TransactOpts) (*txmgr.TxCandidate, error)

// ProveBlockTxBuilder is responsible for building ProveBlock transactions.
type ProveBlockTxBuilder struct {
	rpc                           *rpc.Client
	taikoL1Address                common.Address
	guardianProverMajorityAddress common.Address
	guardianProverMinorityAddress common.Address
}

// NewProveBlockTxBuilder creates a new ProveBlockTxBuilder instance.
func NewProveBlockTxBuilder(
	rpc *rpc.Client,
	taikoL1Address common.Address,
	guardianProverMajorityAddress common.Address,
	guardianProverMinorityAddress common.Address,
) *ProveBlockTxBuilder {
	return &ProveBlockTxBuilder{rpc, taikoL1Address, guardianProverMajorityAddress, guardianProverMinorityAddress}
}

// Build creates a new TaikoL1.ProveBlock transaction with the given nonce.
func (a *ProveBlockTxBuilder) Build(
	blockID *big.Int,
	meta *bindings.TaikoDataBlockMetadata,
	transition *bindings.TaikoDataTransition,
	tierProof *bindings.TaikoDataTierProof,
	tier uint16,
) TxBuilder {
	return func(txOpts *bind.TransactOpts) (*txmgr.TxCandidate, error) {
		var (
			data     []byte
			to       common.Address
			err      error
			guardian = tier >= encoding.TierGuardianMinorityID
		)

		log.Info(
			"Build proof submission transaction",
			"blockID", blockID,
			"gasLimit", txOpts.GasLimit,
			"guardian", guardian,
		)

		if !guardian {
			to = a.taikoL1Address

			input, err := encoding.EncodeProveBlockInput(meta, transition, tierProof)
			if err != nil {
				return nil, err
			}
			if data, err = encoding.TaikoL1ABI.Pack("proveBlock", blockID.Uint64(), input); err != nil {
				if isSubmitProofTxErrorRetryable(err, blockID) {
					return nil, err
				}
				return nil, ErrUnretryableSubmission
			}
		} else {
			if tier > encoding.TierGuardianMinorityID {
				to = a.guardianProverMajorityAddress
			} else if tier == encoding.TierGuardianMinorityID && a.guardianProverMinorityAddress != ZeroAddress {
				to = a.guardianProverMinorityAddress
			} else {
				return nil, fmt.Errorf("tier %d need set guardianProverMinorityAddress", tier)
			}
			if data, err = encoding.GuardianProverABI.Pack("approve", *meta, *transition, *tierProof); err != nil {
				if isSubmitProofTxErrorRetryable(err, blockID) {
					return nil, err
				}
				return nil, ErrUnretryableSubmission
			}
		}

		return &txmgr.TxCandidate{
			TxData:   data,
			To:       &to,
			Blobs:    nil,
			GasLimit: txOpts.GasLimit,
			Value:    txOpts.Value,
		}, nil
	}
}
