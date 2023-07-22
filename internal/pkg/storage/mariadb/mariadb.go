package mariadb

import (
	"time"

	"github.com/blackhorseya/ekko/internal/pkg/config"
	_ "github.com/go-sql-driver/mysql" // import db driver
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// NewMariadb init mariadb client
func NewMariadb(config *config.Config, logger *zap.Logger) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", config.DB.DSN)
	if err != nil {
		logger.Error("Failed to connect database", zap.Error(err), zap.String("dsn", config.DB.DSN))
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 15)
	db.SetMaxOpenConns(config.DB.Conns)
	db.SetMaxIdleConns(config.DB.Conns)

	return db, nil
}

// ProviderSet is a provider set for mariadb client
var ProviderSet = wire.NewSet(NewMariadb)
