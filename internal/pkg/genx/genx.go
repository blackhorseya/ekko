package genx

import (
	"github.com/blackhorseya/ekko/pkg/genx"
	"github.com/bwmarrin/snowflake"
)

type impl struct {
	node *snowflake.Node
}

// NewGenerator serve caller to create a generator
func NewGenerator(id int64) (genx.Generator, error) {
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
