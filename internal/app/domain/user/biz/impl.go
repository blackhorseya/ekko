package biz

import (
	"github.com/blackhorseya/ekko/internal/app/domain/user/biz/repo"
	"github.com/blackhorseya/ekko/pkg/contextx"
	ub "github.com/blackhorseya/ekko/pkg/entity/domain/user/biz"
	um "github.com/blackhorseya/ekko/pkg/entity/domain/user/model"
	"github.com/blackhorseya/ekko/pkg/genx"
)

type impl struct {
	repo repo.IRepo
	node genx.Generator
}

// NewImpl serve caller to create an IBiz
func NewImpl(repo repo.IRepo, node genx.Generator) ub.IBiz {
	return &impl{
		repo: repo,
		node: node,
	}
}

func (i *impl) Signup(ctx contextx.Contextx, username, password string) (info *um.Profile, err error) {
	// todo: 2023/4/16|sean|impl me
	panic("implement me")
}

func (i *impl) Login(ctx contextx.Contextx, username, password string) (info *um.Profile, err error) {
	// todo: 2023/4/16|sean|impl me
	panic("implement me")
}

func (i *impl) Logout(ctx contextx.Contextx, who *um.Profile) error {
	// todo: 2023/4/16|sean|impl me
	panic("implement me")
}

func (i *impl) WhoAmI(ctx contextx.Contextx, token string) (info *um.Profile, err error) {
	// todo: 2023/4/16|sean|impl me
	panic("implement me")
}
