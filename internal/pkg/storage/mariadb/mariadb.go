package mariadb

import (
	"time"

	"github.com/blackhorseya/ekko/internal/pkg/config"
	_ "github.com/go-sql-driver/mysql" // import db driver
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/github" // import GitHub source
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

// NewMigration init migration
func NewMigration(config *config.Config, rw *sqlx.DB) (*migrate.Migrate, error) {
	instance, err := mysql.WithInstance(rw.DB, &mysql.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(config.DB.Source, "mysql", instance)
	if err != nil {
		return nil, err
	}

	return m, nil
}
