package log

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Options declare log's configuration
type Options struct {
	Level string
}

// NewOptions serve caller to create Options
func NewOptions(v *viper.Viper) (*Options, error) {
	var (
		err error
		o   = new(Options)
	)

	if err = v.UnmarshalKey("log", o); err != nil {
		return nil, err
	}

	return o, nil
}

// New serve caller to create zap.Logger
func New(o *Options) (*zap.Logger, error) {
	var (
		err   error
		level = zap.NewAtomicLevel()
	)

	if err = level.UnmarshalText([]byte(o.Level)); err != nil {
		return nil, err
	}

	if gin.Mode() == gin.ReleaseMode {
		return zap.NewProduction()
	}

	return zap.NewDevelopment()
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(New, NewOptions)
