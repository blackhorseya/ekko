//go:generate wire
//go:build wireinject

package biz

import (
	"github.com/blackhorseya/ekko/internal/app/domain/issue/biz/repo"
	ib "github.com/blackhorseya/ekko/pkg/entity/domain/issue/biz"
	"github.com/blackhorseya/ekko/pkg/genx"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

func CreateBiz(repo repo.IRepo, node genx.Generator) ib.IBiz {
	panic(wire.Build(testProviderSet))
}