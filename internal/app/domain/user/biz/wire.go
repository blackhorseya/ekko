//go:generate wire
//go:build wireinject

package biz

import (
	"github.com/blackhorseya/ekko/internal/app/domain/user/biz/repo"
	ub "github.com/blackhorseya/ekko/pkg/entity/domain/user/biz"
	"github.com/blackhorseya/ekko/pkg/genx"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

func CreateBiz(repo repo.IRepo, node genx.Generator) ub.IBiz {
	panic(wire.Build(testProviderSet))
}
