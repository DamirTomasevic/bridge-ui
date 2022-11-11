package repo

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/taikochain/taiko-mono/packages/relayer"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) (*EventRepository, error) {
	if db == nil {
		return nil, relayer.ErrNoDB
	}

	return &EventRepository{
		db: db,
	}, nil
}

func (r *EventRepository) Save(opts relayer.SaveEventOpts) (*relayer.Event, error) {
	e := &relayer.Event{
		Data:    datatypes.JSON(opts.Data),
		Status:  opts.Status,
		ChainID: opts.ChainID.Int64(),
		Name:    opts.Name,
	}
	if err := r.db.Create(e).Error; err != nil {
		return nil, errors.Wrap(err, "r.db.Create")
	}

	return e, nil
}

func (r *EventRepository) UpdateStatus(id int, status relayer.EventStatus) error {
	e := &relayer.Event{}
	if err := r.db.Where("id = ?", id).First(e).Error; err != nil {
		return errors.Wrap(err, "r.db.First")
	}

	e.Status = status
	if err := r.db.Save(e).Error; err != nil {
		return errors.Wrap(err, "r.db.Save")
	}

	return nil
}

func (r *EventRepository) FindAllByAddress(chainID *big.Int, address common.Address) ([]*relayer.Event, error) {
	e := make([]*relayer.Event, 0)
	if err := r.db.Where("chain_id = ?", chainID.Int64()).
		Find(&e, datatypes.JSONQuery("data").
			Equals(address.Hex(), "Owner")).Error; err != nil {
		return nil, errors.Wrap(err, "r.db.Find")
	}

	return e, nil
}
