package repo

import (
	"math/big"

	"github.com/taikoxyz/taiko-mono/packages/relayer"
	"gorm.io/gorm"
)

type BlockRepository struct {
	db DB
}

func NewBlockRepository(db DB) (*BlockRepository, error) {
	if db == nil {
		return nil, ErrNoDB
	}

	return &BlockRepository{
		db: db,
	}, nil
}

func (r *BlockRepository) startQuery() *gorm.DB {
	return r.db.GormDB().Table("processed_blocks")
}

func (r *BlockRepository) Save(opts relayer.SaveBlockOpts) error {
	exists := &relayer.Block{}
	_ = r.startQuery().
		Where("block_height = ?", opts.Height).
		Where("chain_id = ?", opts.ChainID.Int64()).
		Where("dest_chain_id = ?", opts.DestChainID.Int64()).First(exists)
	// block processed already
	if exists.Height == opts.Height {
		return nil
	}

	b := &relayer.Block{
		Height:      opts.Height,
		Hash:        opts.Hash.String(),
		ChainID:     opts.ChainID.Int64(),
		DestChainID: opts.DestChainID.Int64(),
		EventName:   opts.EventName,
	}
	if err := r.startQuery().Create(b).Error; err != nil {
		return err
	}

	return nil
}

func (r *BlockRepository) GetLatestBlockProcessedForEvent(
	eventName string,
	chainID *big.Int,
	destChainID *big.Int,
) (*relayer.Block, error) {
	b := &relayer.Block{}
	if err := r.
		startQuery().
		Raw(`SELECT id, block_height, hash, event_name, chain_id, dest_chain_id 
		FROM processed_blocks 
		WHERE block_height = 
		( SELECT MAX(block_height) from processed_blocks 
		WHERE chain_id = ? AND dest_chain_id = ? AND event_name = ? )`,
			chainID.Int64(),
			destChainID.Int64(),
			eventName).
		FirstOrInit(b).Error; err != nil {
		return nil, err
	}

	return b, nil
}
