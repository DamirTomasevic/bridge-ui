package message

import (
	"crypto/ecdsa"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/taikochain/taiko-mono/packages/relayer"
	"github.com/taikochain/taiko-mono/packages/relayer/contracts"
	"github.com/taikochain/taiko-mono/packages/relayer/proof"
	"github.com/taikochain/taiko-mono/packages/relayer/repo"
	"gopkg.in/go-playground/assert.v1"
)

func Test_NewProcessor(t *testing.T) {
	tests := []struct {
		name    string
		opts    NewProcessorOpts
		wantErr error
	}{
		{
			"success",
			NewProcessorOpts{
				Prover:           &proof.Prover{},
				ECDSAKey:         &ecdsa.PrivateKey{},
				RPCClient:        &rpc.Client{},
				DestETHClient:    &ethclient.Client{},
				DestBridge:       &contracts.Bridge{},
				EventRepo:        &repo.EventRepository{},
				DestHeaderSyncer: &contracts.IHeaderSync{},
			},
			nil,
		},
		{
			"errNoProver",
			NewProcessorOpts{
				ECDSAKey:         &ecdsa.PrivateKey{},
				RPCClient:        &rpc.Client{},
				DestETHClient:    &ethclient.Client{},
				DestBridge:       &contracts.Bridge{},
				EventRepo:        &repo.EventRepository{},
				DestHeaderSyncer: &contracts.IHeaderSync{},
			},
			relayer.ErrNoProver,
		},
		{
			"errNoECDSAKey",
			NewProcessorOpts{
				Prover: &proof.Prover{},

				RPCClient:        &rpc.Client{},
				DestETHClient:    &ethclient.Client{},
				DestBridge:       &contracts.Bridge{},
				EventRepo:        &repo.EventRepository{},
				DestHeaderSyncer: &contracts.IHeaderSync{},
			},
			relayer.ErrNoECDSAKey,
		},
		{
			"noRpcClient",
			NewProcessorOpts{
				Prover:           &proof.Prover{},
				ECDSAKey:         &ecdsa.PrivateKey{},
				DestETHClient:    &ethclient.Client{},
				DestBridge:       &contracts.Bridge{},
				EventRepo:        &repo.EventRepository{},
				DestHeaderSyncer: &contracts.IHeaderSync{},
			},
			relayer.ErrNoRPCClient,
		},
		{
			"noDestEthClient",
			NewProcessorOpts{
				Prover:           &proof.Prover{},
				ECDSAKey:         &ecdsa.PrivateKey{},
				RPCClient:        &rpc.Client{},
				DestBridge:       &contracts.Bridge{},
				EventRepo:        &repo.EventRepository{},
				DestHeaderSyncer: &contracts.IHeaderSync{},
			},
			relayer.ErrNoEthClient,
		},
		{
			"errNoDestBridge",
			NewProcessorOpts{
				Prover:           &proof.Prover{},
				ECDSAKey:         &ecdsa.PrivateKey{},
				RPCClient:        &rpc.Client{},
				DestETHClient:    &ethclient.Client{},
				EventRepo:        &repo.EventRepository{},
				DestHeaderSyncer: &contracts.IHeaderSync{},
			},
			relayer.ErrNoBridge,
		},
		{
			"errNoEventRepo",
			NewProcessorOpts{
				Prover:           &proof.Prover{},
				ECDSAKey:         &ecdsa.PrivateKey{},
				RPCClient:        &rpc.Client{},
				DestETHClient:    &ethclient.Client{},
				DestBridge:       &contracts.Bridge{},
				DestHeaderSyncer: &contracts.IHeaderSync{},
			},
			relayer.ErrNoEventRepository,
		},
		{
			"errNoTaikoL2",
			NewProcessorOpts{
				Prover:        &proof.Prover{},
				ECDSAKey:      &ecdsa.PrivateKey{},
				RPCClient:     &rpc.Client{},
				DestETHClient: &ethclient.Client{},
				EventRepo:     &repo.EventRepository{},
				DestBridge:    &contracts.Bridge{},
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
