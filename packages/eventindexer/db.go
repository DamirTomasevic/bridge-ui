package eventindexer

import (
	"database/sql"

	"github.com/cyberhorsey/errors"
	"gorm.io/gorm"
)

var (
	ErrNoDB = errors.Validation.NewWithKeyAndDetail("ERR_NO_DB", "DB is required")
)

type DBConnectionOpts struct {
	Name     string
	Password string
	Host     string
	Database string
	OpenFunc func(dsn string) (DB, error)
}

type DB interface {
	DB() (*sql.DB, error)
	GormDB() *gorm.DB
}
