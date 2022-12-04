package mariadb

import (
	"time"

	_ "github.com/go-sql-driver/mysql" // import db driver
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const (
	defaultConns = 100
)

// Options is configuration of database
type Options struct {
	URL   string `json:"url" yaml:"url"`
	Debug bool   `json:"debug" yaml:"debug"`
	Conns int    `json:"conns" yaml:"conns"`
}

// NewOptions serve caller to create an Options
func NewOptions(v *viper.Viper, logger *zap.Logger) (*Options, error) {
	var err error
	o := new(Options)
	if err = v.UnmarshalKey("db", o); err != nil {
		return nil, errors.Wrap(err, "unmarshal db option error")
	}

	logger.Info("load database options success")

	return o, err
}

// NewMariadb init mariadb client
func NewMariadb(o *Options, logger *zap.Logger) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", o.URL)
	if err != nil {
		logger.Error("Failed to connect database", zap.Error(err), zap.String("url", o.URL))
		return nil, err
	}

	conns := o.Conns
	if conns == 0 {
		conns = defaultConns
	}

	db.SetConnMaxLifetime(time.Minute * 15)
	db.SetMaxOpenConns(conns)
	db.SetMaxIdleConns(conns)

	return db, nil
}

// ProviderSet is a provider set for mariadb client
var ProviderSet = wire.NewSet(NewOptions, NewMariadb)
