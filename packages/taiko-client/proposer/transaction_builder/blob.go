package builder

import (
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"math/big"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/kzg4844"

	"github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/pkg/rpc"
	selector "github.com/taikoxyz/taiko-mono/packages/taiko-client/proposer/prover_selector"
)

// BlobTransactionBuilder is responsible for building a TaikoL1.proposeBlock transaction with txList
// bytes saved in blob.
type BlobTransactionBuilder struct {
	rpc                     *rpc.Client
	proposerPrivateKey      *ecdsa.PrivateKey
	proverSelector          selector.ProverSelector
	l1BlockBuilderTip       *big.Int
	taikoL1Address          common.Address
	proverSetAddress        common.Address
	l2SuggestedFeeRecipient common.Address
	gasLimit                uint64
	extraData               string
}

// NewBlobTransactionBuilder creates a new BlobTransactionBuilder instance based on giving configurations.
func NewBlobTransactionBuilder(
	rpc *rpc.Client,
	proposerPrivateKey *ecdsa.PrivateKey,
	proverSelector selector.ProverSelector,
	l1BlockBuilderTip *big.Int,
	taikoL1Address common.Address,
	proverSetAddress common.Address,
	l2SuggestedFeeRecipient common.Address,
	gasLimit uint64,
	extraData string,
) *BlobTransactionBuilder {
	return &BlobTransactionBuilder{
		rpc,
		proposerPrivateKey,
		proverSelector,
		l1BlockBuilderTip,
		taikoL1Address,
		proverSetAddress,
		l2SuggestedFeeRecipient,
		gasLimit,
		extraData,
	}
}

// Build implements the ProposeBlockTransactionBuilder interface.
func (b *BlobTransactionBuilder) Build(
	ctx context.Context,
	tierFees []encoding.TierFee,
	includeParentMetaHash bool,
	txListBytes []byte,
) (*txmgr.TxCandidate, error) {
	var blob = &eth.Blob{}
	if err := blob.FromData(txListBytes); err != nil {
		return nil, err
	}

	// Try to assign a prover.
	maxFee, err := b.proverSelector.AssignProver(
		ctx,
		tierFees,
	)
	if err != nil {
		return nil, err
	}

	// If the current proposer wants to include the parent meta hash, then fetch it from the protocol.
	var parentMetaHash = [32]byte{}
	if includeParentMetaHash {
		if parentMetaHash, err = getParentMetaHash(ctx, b.rpc); err != nil {
			return nil, err
		}
	}

	commitment, err := blob.ComputeKZGCommitment()
	if err != nil {
		return nil, err
	}
	blobHash := kzg4844.CalcBlobHashV1(sha256.New(), &commitment)

	signature, err := crypto.Sign(blobHash[:], b.proposerPrivateKey)
	if err != nil {
		return nil, err
	}
	signature[64] = uint8(uint(signature[64])) + 27

	var (
		to   = &b.taikoL1Address
		data []byte
	)
	if b.proverSetAddress != rpc.ZeroAddress {
		to = &b.proverSetAddress
	}

	// ABI encode the TaikoL1.proposeBlock parameters.
	encodedParams, err := encoding.EncodeBlockParams(&encoding.BlockParams{
		ExtraData:      rpc.StringToBytes32(b.extraData),
		Coinbase:       b.l2SuggestedFeeRecipient,
		ParentMetaHash: parentMetaHash,
		Signature:      signature,
	})
	if err != nil {
		return nil, err
	}

	if b.proverSetAddress != rpc.ZeroAddress {
		data, err = encoding.ProverSetABI.Pack("proposeBlock", encodedParams, []byte{})
		if err != nil {
			return nil, err
		}
	} else {
		data, err = encoding.TaikoL1ABI.Pack("proposeBlock", encodedParams, []byte{})
		if err != nil {
			return nil, err
		}
	}
	// Calculate TaikoData.State slot hash (includes slot 0-6 currently)
	var slotHashes []common.Hash
	for i := 0; i < 7; i++ {
		packedData := append(
			common.LeftPadBytes(b.taikoL1Address.Bytes(), 32),
			common.LeftPadBytes(big.NewInt(int64(i)).Bytes(), 32)...,
		)
		slotHashes = append(slotHashes, crypto.Keccak256Hash(packedData))
	}

	return &txmgr.TxCandidate{
		TxData:   data,
		Blobs:    []*eth.Blob{blob},
		To:       to,
		GasLimit: b.gasLimit,
		Value:    maxFee,
		AccessList: types.AccessList{
			{
				Address:     b.taikoL1Address,
				StorageKeys: slotHashes,
			},
		},
	}, nil
}
