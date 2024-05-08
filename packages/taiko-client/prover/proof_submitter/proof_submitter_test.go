package submitter

import (
	"bytes"
	"context"
	"os"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum-optimism/optimism/op-service/txmgr/metrics"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/suite"

	"github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/driver/chain_syncer/beaconsync"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/driver/chain_syncer/blob"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/driver/state"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/internal/testutils"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/pkg/rpc"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/proposer"
	producer "github.com/taikoxyz/taiko-mono/packages/taiko-client/prover/proof_producer"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/prover/proof_submitter/transaction"
)

type ProofSubmitterTestSuite struct {
	testutils.ClientTestSuite
	submitter  *ProofSubmitter
	contester  *ProofContester
	blobSyncer *blob.Syncer
	proposer   *proposer.Proposer
	proofCh    chan *producer.ProofWithHeader
}

func (s *ProofSubmitterTestSuite) SetupTest() {
	s.ClientTestSuite.SetupTest()

	s.proofCh = make(chan *producer.ProofWithHeader, 1024)

	builder := transaction.NewProveBlockTxBuilder(
		s.RPCClient,
		common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
		common.HexToAddress(os.Getenv("GUARDIAN_PROVER_CONTRACT_ADDRESS")),
		common.HexToAddress(os.Getenv("GUARDIAN_PROVER_MINORITY_ADDRESS")),
	)

	l1ProverPrivKey, err := crypto.ToECDSA(common.FromHex(os.Getenv("L1_PROVER_PRIVATE_KEY")))
	s.Nil(err)

	txMgr, err := txmgr.NewSimpleTxManager(
		"proofSubmitterTestSuite",
		log.Root(),
		new(metrics.NoopTxMetrics),
		txmgr.CLIConfig{
			L1RPCURL:                  os.Getenv("L1_NODE_WS_ENDPOINT"),
			NumConfirmations:          0,
			SafeAbortNonceTooLowCount: txmgr.DefaultBatcherFlagValues.SafeAbortNonceTooLowCount,
			PrivateKey:                common.Bytes2Hex(crypto.FromECDSA(l1ProverPrivKey)),
			FeeLimitMultiplier:        txmgr.DefaultBatcherFlagValues.FeeLimitMultiplier,
			FeeLimitThresholdGwei:     txmgr.DefaultBatcherFlagValues.FeeLimitThresholdGwei,
			MinBaseFeeGwei:            txmgr.DefaultBatcherFlagValues.MinBaseFeeGwei,
			MinTipCapGwei:             txmgr.DefaultBatcherFlagValues.MinTipCapGwei,
			ResubmissionTimeout:       txmgr.DefaultBatcherFlagValues.ResubmissionTimeout,
			ReceiptQueryInterval:      1 * time.Second,
			NetworkTimeout:            txmgr.DefaultBatcherFlagValues.NetworkTimeout,
			TxSendTimeout:             txmgr.DefaultBatcherFlagValues.TxSendTimeout,
			TxNotInMempoolTimeout:     txmgr.DefaultBatcherFlagValues.TxNotInMempoolTimeout,
		},
	)
	s.Nil(err)

	s.submitter, err = NewProofSubmitter(
		s.RPCClient,
		&producer.OptimisticProofProducer{},
		s.proofCh,
		common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
		"test",
		0,
		txMgr,
		builder,
	)
	s.Nil(err)
	s.contester = NewProofContester(
		s.RPCClient,
		0,
		txMgr,
		"test",
		builder,
	)

	// Init calldata syncer
	testState, err := state.New(context.Background(), s.RPCClient)
	s.Nil(err)
	s.Nil(testState.ResetL1Current(context.Background(), common.Big0))

	tracker := beaconsync.NewSyncProgressTracker(s.RPCClient.L2, 30*time.Second)

	s.blobSyncer, err = blob.NewSyncer(
		context.Background(),
		s.RPCClient,
		testState,
		tracker,
		0,
		nil,
	)
	s.Nil(err)

	// Init proposer
	prop := new(proposer.Proposer)
	l1ProposerPrivKey, err := crypto.ToECDSA(common.FromHex(os.Getenv("L1_PROPOSER_PRIVATE_KEY")))
	s.Nil(err)

	s.Nil(prop.InitFromConfig(context.Background(), &proposer.Config{
		ClientConfig: &rpc.ClientConfig{
			L1Endpoint:        os.Getenv("L1_NODE_WS_ENDPOINT"),
			L2Endpoint:        os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT"),
			TaikoL1Address:    common.HexToAddress(os.Getenv("TAIKO_L1_ADDRESS")),
			TaikoL2Address:    common.HexToAddress(os.Getenv("TAIKO_L2_ADDRESS")),
			TaikoTokenAddress: common.HexToAddress(os.Getenv("TAIKO_TOKEN_ADDRESS")),
		},
		AssignmentHookAddress:      common.HexToAddress(os.Getenv("ASSIGNMENT_HOOK_ADDRESS")),
		L1ProposerPrivKey:          l1ProposerPrivKey,
		L2SuggestedFeeRecipient:    common.HexToAddress(os.Getenv("L2_SUGGESTED_FEE_RECIPIENT")),
		ProposeInterval:            1024 * time.Hour,
		MaxProposedTxListsPerEpoch: 1,
		ProverEndpoints:            s.ProverEndpoints,
		OptimisticTierFee:          common.Big256,
		SgxTierFee:                 common.Big256,
		MaxTierFeePriceBumps:       3,
		TierFeePriceBump:           common.Big2,
		L1BlockBuilderTip:          common.Big0,
		TxmgrConfigs: &txmgr.CLIConfig{
			L1RPCURL:                  os.Getenv("L1_NODE_WS_ENDPOINT"),
			NumConfirmations:          0,
			SafeAbortNonceTooLowCount: txmgr.DefaultBatcherFlagValues.SafeAbortNonceTooLowCount,
			PrivateKey:                common.Bytes2Hex(crypto.FromECDSA(l1ProposerPrivKey)),
			FeeLimitMultiplier:        txmgr.DefaultBatcherFlagValues.FeeLimitMultiplier,
			FeeLimitThresholdGwei:     txmgr.DefaultBatcherFlagValues.FeeLimitThresholdGwei,
			MinBaseFeeGwei:            txmgr.DefaultBatcherFlagValues.MinBaseFeeGwei,
			MinTipCapGwei:             txmgr.DefaultBatcherFlagValues.MinTipCapGwei,
			ResubmissionTimeout:       txmgr.DefaultBatcherFlagValues.ResubmissionTimeout,
			ReceiptQueryInterval:      1 * time.Second,
			NetworkTimeout:            txmgr.DefaultBatcherFlagValues.NetworkTimeout,
			TxSendTimeout:             txmgr.DefaultBatcherFlagValues.TxSendTimeout,
			TxNotInMempoolTimeout:     txmgr.DefaultBatcherFlagValues.TxNotInMempoolTimeout,
		},
	}))

	s.proposer = prop
}

func (s *ProofSubmitterTestSuite) TestProofSubmitterRequestProofDeadlineExceeded() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	s.ErrorContains(
		s.submitter.RequestProof(
			ctx, &bindings.TaikoL1ClientBlockProposed{BlockId: common.Big256}), "context deadline exceeded",
	)
}

func (s *ProofSubmitterTestSuite) TestProofSubmitterSubmitProofMetadataNotFound() {
	s.Error(
		s.submitter.SubmitProof(
			context.Background(), &producer.ProofWithHeader{
				BlockID: common.Big256,
				Meta:    &bindings.TaikoDataBlockMetadata{},
				Header:  &types.Header{},
				Opts:    &producer.ProofRequestOptions{},
				Proof:   bytes.Repeat([]byte{0xff}, 100),
			},
		),
	)
}

func (s *ProofSubmitterTestSuite) TestSubmitProofs() {
	events := s.ProposeAndInsertEmptyBlocks(s.proposer, s.blobSyncer)

	for _, e := range events {
		s.Nil(s.submitter.RequestProof(context.Background(), e))
		proofWithHeader := <-s.proofCh
		s.Nil(s.submitter.SubmitProof(context.Background(), proofWithHeader))
	}
}

func (s *ProofSubmitterTestSuite) TestGuardianSubmitProofs() {
	events := s.ProposeAndInsertEmptyBlocks(s.proposer, s.blobSyncer)

	for _, e := range events {
		s.Nil(s.submitter.RequestProof(context.Background(), e))
		proofWithHeader := <-s.proofCh
		proofWithHeader.Tier = encoding.TierGuardianMajorityID
		s.Nil(s.submitter.SubmitProof(context.Background(), proofWithHeader))
	}
}

func (s *ProofSubmitterTestSuite) TestProofSubmitterRequestProofCancelled() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() { time.AfterFunc(2*time.Second, func() { cancel() }) }()

	s.ErrorContains(
		s.submitter.RequestProof(
			ctx, &bindings.TaikoL1ClientBlockProposed{BlockId: common.Big256}), "context canceled",
	)
}

func TestProofSubmitterTestSuite(t *testing.T) {
	suite.Run(t, new(ProofSubmitterTestSuite))
}
