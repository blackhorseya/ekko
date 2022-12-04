//go:generate wire
//go:build wireinject

package biz

import (
	"github.com/blackhorseya/todo-app/internal/app/domain/todo/biz/repo"
	tb "github.com/blackhorseya/todo-app/pkg/entity/domain/todo/biz"
	"github.com/blackhorseya/todo-app/pkg/genx"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

func CreateBiz(repo repo.IRepo, node genx.Generator) tb.IBiz {
	panic(wire.Build(testProviderSet))
}
