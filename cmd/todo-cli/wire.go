//go:build wireinject

package main

import (
	"github.com/blackhorseya/todo-app/internal/app/todo/biz"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	biz.ProviderSetViaHTTP,
)
