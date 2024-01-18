package repo

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/taikoxyz/taiko-mono/packages/eventindexer"
	"github.com/taikoxyz/taiko-mono/packages/eventindexer/pkg/db"
	"gopkg.in/go-playground/assert.v1"
)

func Test_NewProcessedBlockRepo(t *testing.T) {
	tests := []struct {
		name    string
		db      eventindexer.DB
		wantErr error
	}{
		{
			"success",
			&db.DB{},
			nil,
		},
		{
			"noDb",
			nil,
			eventindexer.ErrNoDB,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewProcessedBlockRepository(tt.db)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestIntegration_Block_Save(t *testing.T) {
	db, close, err := testMysql(t)
	assert.Equal(t, nil, err)

	defer close()

	blockRepo, err := NewProcessedBlockRepository(db)
	assert.Equal(t, nil, err)
	tests := []struct {
		name    string
		opts    eventindexer.SaveProcessedBlockOpts
		wantErr error
	}{
		{
			"success",
			eventindexer.SaveProcessedBlockOpts{
				ChainID: big.NewInt(1),
				Height:  100,
				Hash:    common.HexToHash("0x1234"),
			},
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err = blockRepo.Save(tt.opts)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestIntegration_Block_GetLatestBlockProcessedForEvent(t *testing.T) {
	db, close, err := testMysql(t)
	assert.Equal(t, nil, err)

	defer close()

	blockRepo, err := NewProcessedBlockRepository(db)
	assert.Equal(t, nil, err)
	tests := []struct {
		name      string
		eventName string
		chainID   *big.Int
		wantErr   error
	}{
		{
			"success",
			eventindexer.EventNameBlockProposed,
			big.NewInt(1),
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := blockRepo.GetLatestBlockProcessed(tt.chainID)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
