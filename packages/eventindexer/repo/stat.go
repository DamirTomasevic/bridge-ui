package repo

import (
	"context"

	"github.com/pkg/errors"
	"github.com/taikoxyz/taiko-mono/packages/eventindexer"
)

type StatRepository struct {
	db eventindexer.DB
}

func NewStatRepository(db eventindexer.DB) (*StatRepository, error) {
	if db == nil {
		return nil, eventindexer.ErrNoDB
	}

	return &StatRepository{
		db: db,
	}, nil
}

func (r *StatRepository) Save(ctx context.Context, opts eventindexer.SaveStatOpts) (*eventindexer.Stat, error) {
	s := &eventindexer.Stat{}

	if err := r.db.
		GormDB().
		FirstOrCreate(s).
		Error; err != nil {
		return nil, errors.Wrap(err, "r.db.gormDB.FirstOrCreate")
	}

	if opts.ProofReward != nil {
		s.AverageProofReward = *opts.ProofReward
	}

	if opts.ProofTime != nil {
		s.NumProofs++
		s.AverageProofTime = *opts.ProofTime
	}

	if err := r.db.GormDB().Save(s).Error; err != nil {
		return nil, errors.Wrap(err, "r.db.Save")
	}

	return s, nil
}

func (r *StatRepository) Find(ctx context.Context) (*eventindexer.Stat, error) {
	s := &eventindexer.Stat{}

	if err := r.db.
		GormDB().
		FirstOrCreate(s).
		Error; err != nil {
		return nil, errors.Wrap(err, "r.db.gormDB.FirstOrCreate")
	}

	return s, nil
}
