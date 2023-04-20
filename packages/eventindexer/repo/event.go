package repo

import (
	"context"

	"github.com/pkg/errors"
	"github.com/taikoxyz/taiko-mono/packages/eventindexer"
	"gorm.io/datatypes"
)

type EventRepository struct {
	db eventindexer.DB
}

func NewEventRepository(db eventindexer.DB) (*EventRepository, error) {
	if db == nil {
		return nil, eventindexer.ErrNoDB
	}

	return &EventRepository{
		db: db,
	}, nil
}

func (r *EventRepository) Save(ctx context.Context, opts eventindexer.SaveEventOpts) (*eventindexer.Event, error) {
	e := &eventindexer.Event{
		Data:    datatypes.JSON(opts.Data),
		ChainID: opts.ChainID.Int64(),
		Name:    opts.Name,
		Event:   opts.Event,
		Address: opts.Address,
	}

	if err := r.db.GormDB().Create(e).Error; err != nil {
		return nil, errors.Wrap(err, "r.db.Create")
	}

	return e, nil
}

func (r *EventRepository) FindUniqueProvers(
	ctx context.Context,
) ([]eventindexer.UniqueProversResponse, error) {
	addrs := make([]eventindexer.UniqueProversResponse, 0)

	if err := r.db.GormDB().
		Raw("SELECT address, count(*) AS count FROM events WHERE event = ? GROUP BY address",
			eventindexer.EventNameBlockProven).
		FirstOrInit(&addrs).Error; err != nil {
		return nil, errors.Wrap(err, "r.db.FirstOrInit")
	}

	return addrs, nil
}

func (r *EventRepository) FindUniqueProposers(
	ctx context.Context,
) ([]eventindexer.UniqueProposersResponse, error) {
	addrs := make([]eventindexer.UniqueProposersResponse, 0)

	if err := r.db.GormDB().
		Raw("SELECT address, count(*) AS count FROM events WHERE event = ? GROUP BY address",
			eventindexer.EventNameBlockProposed).
		FirstOrInit(&addrs).Error; err != nil {
		return nil, errors.Wrap(err, "r.db.FirstOrInit")
	}

	return addrs, nil
}

func (r *EventRepository) GetCountByAddressAndEventName(
	ctx context.Context,
	address string,
	event string,
) (int, error) {
	var count int

	if err := r.db.GormDB().
		Raw("SELECT count(*) AS count FROM events WHERE event = ? AND address = ?", event, address).
		FirstOrInit(&count).Error; err != nil {
		return 0, errors.Wrap(err, "r.db.FirstOrInit")
	}

	return count, nil
}
