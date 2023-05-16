package message

import (
	"crypto/ecdsa"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/taikoxyz/taiko-mono/packages/relayer"
	"github.com/taikoxyz/taiko-mono/packages/relayer/contracts/bridge"
	"github.com/taikoxyz/taiko-mono/packages/relayer/contracts/icrosschainsync"
	"github.com/taikoxyz/taiko-mono/packages/relayer/mock"
	"github.com/taikoxyz/taiko-mono/packages/relayer/proof"
	"github.com/taikoxyz/taiko-mono/packages/relayer/repo"
	"gopkg.in/go-playground/assert.v1"
)

var dummyEcdsaKey = "8da4ef21b864d2cc526dbdb2a120bd2874c36c9d0a1fb7f8c63d7f7a8b41de8f"

func newTestProcessor(profitableOnly relayer.ProfitableOnly) *Processor {
	privateKey, _ := crypto.HexToECDSA(dummyEcdsaKey)

	prover, _ := proof.New(
		&mock.Blocker{},
	)

	return &Processor{
		eventRepo:                 &mock.EventRepository{},
		destBridge:                &mock.Bridge{},
		srcEthClient:              &mock.EthClient{},
		destEthClient:             &mock.EthClient{},
		destTokenVault:            &mock.TokenVault{},
		mu:                        &sync.Mutex{},
		ecdsaKey:                  privateKey,
		destHeaderSyncer:          &mock.HeaderSyncer{},
		prover:                    prover,
		rpc:                       &mock.Caller{},
		profitableOnly:            profitableOnly,
		headerSyncIntervalSeconds: 1,
		confTimeoutInSeconds:      900,
	}
}
func Test_NewProcessor(t *testing.T) {
	tests := []struct {
		name    string
		opts    NewProcessorOpts
		wantErr error
	}{
		{
			"success",
			NewProcessorOpts{
				Prover:                        &proof.Prover{},
				ECDSAKey:                      &ecdsa.PrivateKey{},
				RPCClient:                     &rpc.Client{},
				SrcETHClient:                  &ethclient.Client{},
				DestETHClient:                 &ethclient.Client{},
				DestBridge:                    &bridge.Bridge{},
				EventRepo:                     &repo.EventRepository{},
				DestHeaderSyncer:              &icrosschainsync.ICrossChainSync{},
				Confirmations:                 1,
				ConfirmationsTimeoutInSeconds: 900,
			},
			nil,
		},
		{
			"errNoConfirmationsTimeoutInSeconds",
			NewProcessorOpts{
				Prover:           &proof.Prover{},
				ECDSAKey:         &ecdsa.PrivateKey{},
				RPCClient:        &rpc.Client{},
				SrcETHClient:     &ethclient.Client{},
				DestETHClient:    &ethclient.Client{},
				DestBridge:       &bridge.Bridge{},
				EventRepo:        &repo.EventRepository{},
				DestHeaderSyncer: &icrosschainsync.ICrossChainSync{},
				Confirmations:    1,
			},
			relayer.ErrInvalidConfirmationsTimeoutInSeconds,
		},
		{
			"errNoConfirmations",
			NewProcessorOpts{
				Prover:                        &proof.Prover{},
				ECDSAKey:                      &ecdsa.PrivateKey{},
				RPCClient:                     &rpc.Client{},
				SrcETHClient:                  &ethclient.Client{},
				DestETHClient:                 &ethclient.Client{},
				DestBridge:                    &bridge.Bridge{},
				EventRepo:                     &repo.EventRepository{},
				DestHeaderSyncer:              &icrosschainsync.ICrossChainSync{},
				ConfirmationsTimeoutInSeconds: 900,
			},
			relayer.ErrInvalidConfirmations,
		},
		{
			"errNoSrcClient",
			NewProcessorOpts{
				Prover:                        &proof.Prover{},
				ECDSAKey:                      &ecdsa.PrivateKey{},
				RPCClient:                     &rpc.Client{},
				DestETHClient:                 &ethclient.Client{},
				DestBridge:                    &bridge.Bridge{},
				EventRepo:                     &repo.EventRepository{},
				DestHeaderSyncer:              &icrosschainsync.ICrossChainSync{},
				Confirmations:                 1,
				ConfirmationsTimeoutInSeconds: 900,
			},
			relayer.ErrNoEthClient,
		},
		{
			"errNoProver",
			NewProcessorOpts{
				ECDSAKey:                      &ecdsa.PrivateKey{},
				RPCClient:                     &rpc.Client{},
				SrcETHClient:                  &ethclient.Client{},
				DestETHClient:                 &ethclient.Client{},
				DestBridge:                    &bridge.Bridge{},
				EventRepo:                     &repo.EventRepository{},
				Confirmations:                 1,
				DestHeaderSyncer:              &icrosschainsync.ICrossChainSync{},
				ConfirmationsTimeoutInSeconds: 900,
			},
			relayer.ErrNoProver,
		},
		{
			"errNoECDSAKey",
			NewProcessorOpts{
				Prover: &proof.Prover{},

				RPCClient:                     &rpc.Client{},
				SrcETHClient:                  &ethclient.Client{},
				DestETHClient:                 &ethclient.Client{},
				DestBridge:                    &bridge.Bridge{},
				EventRepo:                     &repo.EventRepository{},
				DestHeaderSyncer:              &icrosschainsync.ICrossChainSync{},
				Confirmations:                 1,
				ConfirmationsTimeoutInSeconds: 900,
			},
			relayer.ErrNoECDSAKey,
		},
		{
			"noRpcClient",
			NewProcessorOpts{
				Prover:                        &proof.Prover{},
				ECDSAKey:                      &ecdsa.PrivateKey{},
				SrcETHClient:                  &ethclient.Client{},
				DestETHClient:                 &ethclient.Client{},
				DestBridge:                    &bridge.Bridge{},
				EventRepo:                     &repo.EventRepository{},
				DestHeaderSyncer:              &icrosschainsync.ICrossChainSync{},
				Confirmations:                 1,
				ConfirmationsTimeoutInSeconds: 900,
			},
			relayer.ErrNoRPCClient,
		},
		{
			"noDestEthClient",
			NewProcessorOpts{
				Prover:                        &proof.Prover{},
				ECDSAKey:                      &ecdsa.PrivateKey{},
				RPCClient:                     &rpc.Client{},
				SrcETHClient:                  &ethclient.Client{},
				DestBridge:                    &bridge.Bridge{},
				EventRepo:                     &repo.EventRepository{},
				DestHeaderSyncer:              &icrosschainsync.ICrossChainSync{},
				Confirmations:                 1,
				ConfirmationsTimeoutInSeconds: 900,
			},
			relayer.ErrNoEthClient,
		},
		{
			"errNoDestBridge",
			NewProcessorOpts{
				Prover:                        &proof.Prover{},
				ECDSAKey:                      &ecdsa.PrivateKey{},
				RPCClient:                     &rpc.Client{},
				SrcETHClient:                  &ethclient.Client{},
				DestETHClient:                 &ethclient.Client{},
				EventRepo:                     &repo.EventRepository{},
				DestHeaderSyncer:              &icrosschainsync.ICrossChainSync{},
				Confirmations:                 1,
				ConfirmationsTimeoutInSeconds: 900,
			},
			relayer.ErrNoBridge,
		},
		{
			"errNoEventRepo",
			NewProcessorOpts{
				Prover:                        &proof.Prover{},
				ECDSAKey:                      &ecdsa.PrivateKey{},
				RPCClient:                     &rpc.Client{},
				SrcETHClient:                  &ethclient.Client{},
				DestETHClient:                 &ethclient.Client{},
				DestBridge:                    &bridge.Bridge{},
				DestHeaderSyncer:              &icrosschainsync.ICrossChainSync{},
				Confirmations:                 1,
				ConfirmationsTimeoutInSeconds: 900,
			},
			relayer.ErrNoEventRepository,
		},
		{
			"errNoTaikoL2",
			NewProcessorOpts{
				Prover:                        &proof.Prover{},
				ECDSAKey:                      &ecdsa.PrivateKey{},
				RPCClient:                     &rpc.Client{},
				SrcETHClient:                  &ethclient.Client{},
				DestETHClient:                 &ethclient.Client{},
				EventRepo:                     &repo.EventRepository{},
				DestBridge:                    &bridge.Bridge{},
				Confirmations:                 1,
				ConfirmationsTimeoutInSeconds: 900,
			},
			relayer.ErrNoTaikoL2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewProcessor(tt.opts)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
