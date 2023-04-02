//go:generate wire
//go:build wireinject

package biz

import (
	"github.com/blackhorseya/ekko/internal/app/domain/task/biz/repo"
	tb "github.com/blackhorseya/ekko/pkg/entity/domain/task/biz"
	"github.com/blackhorseya/ekko/pkg/genx"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

func CreateBiz(repo repo.IRepo, node genx.Generator) tb.IBiz {
	panic(wire.Build(testProviderSet))
}
