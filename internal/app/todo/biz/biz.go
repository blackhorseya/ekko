package biz

import (
	"github.com/blackhorseya/todo-app/internal/app/todo/biz/health"
	"github.com/blackhorseya/todo-app/internal/app/todo/biz/todo"
	"github.com/google/wire"
)

var (
	// ProviderSet is a business provider set
	ProviderSet = wire.NewSet(health.ProviderSet, todo.ProviderSet)

	// ProviderSetViaHTTP is a http provider set for wire
	ProviderSetViaHTTP = wire.NewSet(health.ProviderSet, todo.ProviderSetViaHTTP)
)
