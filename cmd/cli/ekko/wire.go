//go:generate wire
//go:build wireinject

package main

import (
	"github.com/blackhorseya/ekko/internal/adapter/cli"
	"github.com/blackhorseya/ekko/internal/app/domain/issue/biz"
	"github.com/blackhorseya/ekko/internal/app/domain/issue/biz/repo"
	"github.com/blackhorseya/ekko/internal/pkg/config"
	"github.com/blackhorseya/ekko/internal/pkg/genx"
	"github.com/blackhorseya/ekko/internal/pkg/httpx"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	// infrastructure
	config.NewViperSet,
	httpx.ClientSet,
	genx.ProviderSet,

	// adapters
	cli.ProviderSet,

	// implementation
	biz.IssueSet,
	repo.HTTPClientSet,

	// main
	NewService,
)

// InitializeService serve caller to initialize service
func InitializeService(i int64) (*Service, error) {
	panic(wire.Build(providerSet))
}
