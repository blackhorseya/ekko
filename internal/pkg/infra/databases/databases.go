package databases

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Options is configuration of database
type Options struct {
	URL   string `json:"url" yaml:"url"`
	Conns int    `json:"conns" yaml:"conns"`
}

// NewOptions serve caller to create an Options
func NewOptions(v *viper.Viper, logger *zap.Logger) (*Options, error) {
	var err error
	o := new(Options)
	if err = v.UnmarshalKey("db", o); err != nil {
		return nil, errors.Wrap(err, "unmarshal db option error")
	}

	logger.Info("load database options success", zap.String("url", o.URL))

	return o, err
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewOptions, NewMariadb)
