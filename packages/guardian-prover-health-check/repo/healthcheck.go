package repo

import (
	"context"
	"net/http"

	"github.com/morkid/paginate"
	guardianproverhealthcheck "github.com/taikoxyz/taiko-mono/packages/guardian-prover-health-check"
	"gorm.io/gorm"
)

var (
	expectedHealthChecksPer24Hours = 7200
)

type HealthCheckRepository struct {
	db DB
}

func NewHealthCheckRepository(db DB) (*HealthCheckRepository, error) {
	if db == nil {
		return nil, ErrNoDB
	}

	return &HealthCheckRepository{
		db: db,
	}, nil
}

func (r *HealthCheckRepository) startQuery() *gorm.DB {
	return r.db.GormDB().Table("health_checks")
}

func (r *HealthCheckRepository) Get(
	ctx context.Context,
	req *http.Request,
) (paginate.Page, error) {
	pg := paginate.New(&paginate.Config{
		DefaultSize: 100,
	})

	reqCtx := pg.With(r.startQuery())

	page := reqCtx.Request(req).Response(&[]guardianproverhealthcheck.HealthCheck{})

	return page, nil
}

func (r *HealthCheckRepository) GetByGuardianProverAddress(
	ctx context.Context,
	req *http.Request,
	address string,
) (paginate.Page, error) {
	pg := paginate.New(&paginate.Config{
		DefaultSize: 100,
	})

	reqCtx := pg.With(r.startQuery().Order("created_at desc").
		Where("recovered_address = ?", address))

	page := reqCtx.Request(req).Response(&[]guardianproverhealthcheck.HealthCheck{})

	return page, nil
}

func (r *HealthCheckRepository) GetMostRecentByGuardianProverAddress(
	ctx context.Context,
	req *http.Request,
	address string,
) (*guardianproverhealthcheck.HealthCheck, error) {
	hc := &guardianproverhealthcheck.HealthCheck{}

	if err := r.startQuery().Order("created_at desc").
		Where("recovered_address = ?", address).Limit(1).
		Scan(hc).Error; err != nil {
		return nil, err
	}

	return hc, nil
}

func (r *HealthCheckRepository) Save(opts guardianproverhealthcheck.SaveHealthCheckOpts) error {
	b := &guardianproverhealthcheck.HealthCheck{
		Alive:            opts.Alive,
		ExpectedAddress:  opts.ExpectedAddress,
		RecoveredAddress: opts.RecoveredAddress,
		SignedResponse:   opts.SignedResponse,
		GuardianProverID: opts.GuardianProverID,
		LatestL1Block:    opts.LatestL1Block,
		LatestL2Block:    opts.LatestL2Block,
	}
	if err := r.startQuery().Create(b).Error; err != nil {
		return err
	}

	return nil
}

func (r *HealthCheckRepository) GetUptimeByGuardianProverAddress(
	ctx context.Context,
	address string,
) (float64, int, error) {
	var count int64

	var query string = `SELECT COUNT(*) 
	FROM health_checks 
	WHERE recovered_address = ? AND
	created_at > NOW() - INTERVAL 1 DAY`

	if err := r.db.GormDB().Raw(query, address).Scan(&count).Error; err != nil {
		return 0, 0, err
	}

	uptimePercentage := (float64(count) / float64(expectedHealthChecksPer24Hours)) * 100

	return uptimePercentage, int(count), nil
}
