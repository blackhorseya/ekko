//go:build wireinject

package main

import (
	"github.com/blackhorseya/gocommon/pkg/restclient"
	"github.com/blackhorseya/todo-app/internal/app/todo/api/cmd"
	"github.com/blackhorseya/todo-app/internal/app/todo/biz"
	"github.com/google/wire"
	"github.com/spf13/cobra"
)

var providerSet = wire.NewSet(
	restclient.ProviderSet,
	cmd.ProviderSet,
	biz.ProviderSetViaHTTP,
)

func CreateApp() (*cobra.Command, error) {
	panic(wire.Build(providerSet))
}
