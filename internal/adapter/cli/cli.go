package cli

import (
	"errors"

	"github.com/blackhorseya/ekko/pkg/adapters"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewCLI)

type impl struct {
}

// NewCLI serve caller to create adapter cli
func NewCLI() adapters.CLI {
	return &impl{}
}

func (i *impl) Execute() error {
	// todo: 2023/4/4|sean|implement me
	return errors.New("implement me")
}
