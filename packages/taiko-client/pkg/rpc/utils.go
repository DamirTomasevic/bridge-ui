package rpc

import (
	"context"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"

	"github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/internal/utils"
)

var (
	ZeroAddress         common.Address
	BlobBytes                  = params.BlobTxBytesPerFieldElement * params.BlobTxFieldElementsPerBlob
	BlockMaxTxListBytes uint64 = (params.BlobTxBytesPerFieldElement - 1) * params.BlobTxFieldElementsPerBlob
)

// GetProtocolStateVariables gets the protocol states from TaikoL1 contract.
func GetProtocolStateVariables(
	taikoL1Client *bindings.TaikoL1Client,
	opts *bind.CallOpts,
) (*struct {
	A bindings.TaikoDataSlotA
	B bindings.TaikoDataSlotB
}, error) {
	slotA, slotB, err := taikoL1Client.GetStateVariables(opts)
	if err != nil {
		return nil, err
	}
	return &struct {
		A bindings.TaikoDataSlotA
		B bindings.TaikoDataSlotB
	}{slotA, slotB}, nil
}

// CheckProverBalance checks if the prover has the necessary allowance and
// balance for a prover to pay the liveness bond.
func CheckProverBalance(
	ctx context.Context,
	rpc *Client,
	prover common.Address,
	address common.Address,
	bond *big.Int,
) (bool, error) {
	ctxWithTimeout, cancel := ctxWithTimeoutOrDefault(ctx, defaultTimeout)
	defer cancel()

	// Check allowance on taiko token contract
	allowance, err := rpc.TaikoToken.Allowance(&bind.CallOpts{Context: ctxWithTimeout}, prover, address)
	if err != nil {
		return false, err
	}

	log.Info(
		"Prover allowance for TaikoL1 contract",
		"allowance", utils.WeiToEther(allowance),
		"address", prover.Hex(),
		"bond", utils.WeiToEther(bond),
	)

	// Check prover's taiko token balance
	balance, err := rpc.TaikoToken.BalanceOf(&bind.CallOpts{Context: ctxWithTimeout}, prover)
	if err != nil {
		return false, err
	}

	log.Info(
		"Prover's wallet taiko token balance",
		"balance", utils.WeiToEther(balance),
		"address", prover.Hex(),
		"bond", utils.WeiToEther(bond),
	)

	if bond.Cmp(allowance) > 0 || bond.Cmp(balance) > 0 {
		log.Info(
			"Assigned prover does not have required on-chain token balance or allowance",
			"providedProver", prover.Hex(),
			"taikoTokenBalance", utils.WeiToEther(balance),
			"allowance", utils.WeiToEther(allowance),
			"bond", utils.WeiToEther(bond),
		)
		return false, nil
	}

	return true, nil
}

// BlockProofStatus represents the proving status of the given L2 block.
type BlockProofStatus struct {
	IsSubmitted            bool
	Invalid                bool
	CurrentTransitionState *bindings.TaikoDataTransitionState
	ParentHeader           *types.Header
}

// GetBlockProofStatus checks whether the L2 block still needs a new proof or a new contest.
// Here are the possible status:
// 1. No proof on chain at all.
// 2. A valid proof has been submitted.
// 3. An invalid proof has been submitted, and there is no valid contest.
// 4. An invalid proof has been submitted, and there is a valid contest.
func GetBlockProofStatus(
	ctx context.Context,
	cli *Client,
	id *big.Int,
	proverAddress common.Address,
	proverSetAddress common.Address,
) (*BlockProofStatus, error) {
	ctxWithTimeout, cancel := ctxWithTimeoutOrDefault(ctx, defaultTimeout)
	defer cancel()

	// Get the local L2 parent header.
	parent, err := cli.L2.HeaderByNumber(ctxWithTimeout, new(big.Int).Sub(id, common.Big1))
	if err != nil {
		return nil, err
	}

	// Get the transition state from TaikoL1 contract.
	transition, err := cli.TaikoL1.GetTransition0(
		&bind.CallOpts{Context: ctxWithTimeout},
		id.Uint64(),
		parent.Hash(),
	)
	if err != nil {
		if !strings.Contains(encoding.TryParsingCustomError(err).Error(), "L1_TRANSITION_NOT_FOUND") {
			return nil, encoding.TryParsingCustomError(err)
		}

		// Status 1, no proof on chain at all.
		return &BlockProofStatus{IsSubmitted: false, ParentHeader: parent}, nil
	}

	header, err := cli.WaitL2Header(ctxWithTimeout, id)
	if err != nil {
		return nil, err
	}

	if header.Hash() != transition.BlockHash || transition.StateRoot != header.Root {
		log.Info(
			"Different block hash or state root detected, try submitting a contest",
			"localBlockHash", header.Hash(),
			"protocolTransitionBlockHash", common.BytesToHash(transition.BlockHash[:]),
			"localStateRoot", header.Root,
			"protocolTransitionStateRoot", common.BytesToHash(transition.StateRoot[:]),
		)
		return &BlockProofStatus{
			IsSubmitted:            true,
			Invalid:                true,
			CurrentTransitionState: &transition,
			ParentHeader:           parent,
		}, nil
	}

	if proverAddress == transition.Prover ||
		(proverSetAddress != ZeroAddress && transition.Prover == proverSetAddress) {
		log.Info(
			"📬 Block's proof has already been submitted by current prover",
			"blockID", id,
			"parent", parent.Hash().Hex(),
			"hash", common.Bytes2Hex(transition.BlockHash[:]),
			"stateRoot", common.Bytes2Hex(transition.StateRoot[:]),
			"timestamp", transition.Timestamp,
			"contester", transition.Contester,
		)
		return &BlockProofStatus{
			IsSubmitted:            true,
			Invalid:                transition.Contester != ZeroAddress,
			ParentHeader:           parent,
			CurrentTransitionState: &transition,
		}, nil
	}

	log.Info(
		"📬 Block's proof has already been submitted by another prover",
		"blockID", id,
		"prover", transition.Prover,
		"parent", parent.Hash().Hex(),
		"hash", common.Bytes2Hex(transition.BlockHash[:]),
		"stateRoot", common.Bytes2Hex(transition.StateRoot[:]),
		"timestamp", transition.Timestamp,
		"contester", transition.Contester,
	)

	return &BlockProofStatus{
		IsSubmitted:            true,
		Invalid:                transition.Contester != ZeroAddress,
		ParentHeader:           parent,
		CurrentTransitionState: &transition,
	}, nil
}

// SetHead makes a `debug_setHead` RPC call to set the chain's head, should only be used
// for testing purpose.
func SetHead(ctx context.Context, client *EthClient, headNum *big.Int) error {
	ctxWithTimeout, cancel := ctxWithTimeoutOrDefault(ctx, defaultTimeout)
	defer cancel()

	return client.SetHead(ctxWithTimeout, headNum)
}

// StringToBytes32 converts the given string to [32]byte.
func StringToBytes32(str string) [32]byte {
	var b [32]byte
	copy(b[:], []byte(str))

	return b
}

// ctxWithTimeoutOrDefault sets a context timeout if the deadline has not passed or is not set,
// and otherwise returns the context as passed in. cancel func is always set to an empty function
// so is safe to defer the cancel.
func ctxWithTimeoutOrDefault(ctx context.Context, defaultTimeout time.Duration) (context.Context, context.CancelFunc) {
	if utils.IsNil(ctx) {
		return context.WithTimeout(context.Background(), defaultTimeout)
	}
	if _, ok := ctx.Deadline(); !ok {
		return context.WithTimeout(ctx, defaultTimeout)
	}

	return ctx, func() {}
}
