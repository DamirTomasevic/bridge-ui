package chainsyncer

import (
	"bytes"
	"context"

	"os"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/suite"

	"github.com/taikoxyz/taiko-mono/packages/taiko-client/driver/state"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/internal/testutils"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/pkg/rpc"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/proposer"
)

type ChainSyncerTestSuite struct {
	testutils.ClientTestSuite
	s          *L2ChainSyncer
	snapshotID string
	p          testutils.Proposer
}

func (s *ChainSyncerTestSuite) SetupTest() {
	s.ClientTestSuite.SetupTest()

	state, err := state.New(context.Background(), s.RPCClient)
	s.Nil(err)

	syncer, err := New(
		context.Background(),
		s.RPCClient,
		state,
		false,
		1*time.Hour,
		0,
		nil,
		nil,
	)
	s.Nil(err)
	s.s = syncer

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
			ProverSetAddress:  common.HexToAddress(os.Getenv("PROVER_SET_ADDRESS")),
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
		ExtraData:                  "test",
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

	s.p = prop
}

func (s *ChainSyncerTestSuite) TestGetInnerSyncers() {
	s.NotNil(s.s.BeaconSyncer())
	s.NotNil(s.s.BlobSyncer())
}

func (s *ChainSyncerTestSuite) TestSync() {
	s.Nil(s.s.Sync())
}

func (s *ChainSyncerTestSuite) TestAheadOfProtocolVerifiedHead2() {
	s.TakeSnapshot()
	// propose a couple blocks
	s.ProposeAndInsertEmptyBlocks(s.p, s.s.blobSyncer)

	// NOTE: need to prove the proposed blocks to be verified, writing helper function
	// generate transactopts to interact with TaikoL1 contract with.
	privKey, err := crypto.ToECDSA(common.FromHex(os.Getenv("L1_PROVER_PRIVATE_KEY")))
	s.Nil(err)
	opts, err := bind.NewKeyedTransactorWithChainID(privKey, s.RPCClient.L1.ChainID)
	s.Nil(err)

	head, err := s.RPCClient.L1.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	l2Head, err := s.RPCClient.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)
	s.Equal("test", string(bytes.TrimRight(l2Head.Extra, "\x00")))
	log.Info("L1HeaderByNumber head", "number", head.Number)
	// (equiv to s.state.GetL2Head().Number)
	log.Info("L2HeaderByNumber head", "number", l2Head.Number)

	// increase evm time to make blocks verifiable.
	s.IncreaseTime(uint64((1024 * time.Hour).Seconds()))

	// interact with TaikoL1 contract to allow for verification of L2 blocks
	tx, err := s.s.rpc.TaikoL1.VerifyBlocks(opts, uint64(3))
	s.Nil(err)
	s.NotNil(tx)

	head2, err := s.RPCClient.L1.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	l2Head2, err := s.RPCClient.L2.HeaderByNumber(context.Background(), nil)
	s.Nil(err)

	log.Info("L1HeaderByNumber head2", "number", head2.Number)
	log.Info("L2HeaderByNumber head", "number", l2Head2.Number)

	s.RevertSnapshot()
}

func TestChainSyncerTestSuite(t *testing.T) {
	suite.Run(t, new(ChainSyncerTestSuite))
}

func (s *ChainSyncerTestSuite) TakeSnapshot() {
	// record snapshot state to revert to before changes
	s.snapshotID = s.SetL1Snapshot()
}

func (s *ChainSyncerTestSuite) RevertSnapshot() {
	// revert to the snapshot state so protocol configs are unaffected
	s.RevertL1Snapshot(s.snapshotID)
	s.Nil(rpc.SetHead(context.Background(), s.RPCClient.L2, common.Big0))
}

func (s *ChainSyncerTestSuite) TestAheadOfProtocolVerifiedHead() {
	s.True(s.s.AheadOfProtocolVerifiedHead(0))
}
