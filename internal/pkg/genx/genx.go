package genx

import (
	"github.com/blackhorseya/ekko/pkg/genx"
	"github.com/bwmarrin/snowflake"
	"github.com/google/wire"
)

type impl struct {
	node *snowflake.Node
}

func NewImpl(id int64) (genx.Generator, error) {
	node, err := snowflake.NewNode(id)
	if err != nil {
		return nil, err
	}

	return &impl{
		node: node,
	}, nil
}

func (i *impl) Int64() int64 {
	return i.node.Generate().Int64()
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl)
