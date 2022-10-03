package databases

import (
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql" // import db driver
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const (
	defaultConns = 100
)

var pool sync.Map

// NewMariadb init mariadb client
func NewMariadb(o *Options, logger *zap.Logger) (*sqlx.DB, error) {
	entry, load := pool.LoadOrStore("jarvan", creatConnection(o, logger))
	if !load {
		logger.Info("[DB] client created")
	}

	return entry.(*sqlx.DB), nil
}

func creatConnection(o *Options, logger *zap.Logger) *sqlx.DB {
	db, err := sqlx.Open("mysql", o.URL)
	if err != nil {
		logger.Error("Connection to database failed.", zap.Error(err))
		return nil
	}

	conns := o.Conns
	if conns == 0 {
		conns = defaultConns
	}

	db.SetConnMaxLifetime(time.Minute * 15)
	db.SetMaxOpenConns(conns)
	db.SetMaxIdleConns(conns)

	return db
}
