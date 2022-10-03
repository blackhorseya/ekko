package node

import (
	"github.com/bwmarrin/snowflake"
	"github.com/google/wire"
)

// Generator declare node generator functions
//
//go:generate mockery --all --inpackage
type Generator interface {
	Generate() snowflake.ID
}

type impl struct {
	node *snowflake.Node
}

func NewImpl() (Generator, error) {
	node, err := snowflake.NewNode(0)
	if err != nil {
		return nil, err
	}

	return &impl{
		node: node,
	}, nil
}

func (i *impl) Generate() snowflake.ID {
	return i.node.Generate()
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl)
