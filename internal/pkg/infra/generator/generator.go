package generator

import (
	"github.com/bwmarrin/snowflake"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

// Options declare app's configuration
type Options struct {
	NodeID int
}

// NewOptions serve caller to create Options
func NewOptions(v *viper.Viper) (*Options, error) {
	var (
		err error
		o   = new(Options)
	)

	if err = v.UnmarshalKey("app", o); err != nil {
		return nil, err
	}

	return o, nil
}

// New serve caller to create snowflake.Node
func New(o *Options) (*snowflake.Node, error) {
	node, err := snowflake.NewNode(int64(o.NodeID))
	if err != nil {
		return nil, err
	}

	return node, nil
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(New, NewOptions)
