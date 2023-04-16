package repo

import (
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/blackhorseya/ekko/pkg/entity/domain/user/model"
	"github.com/jmoiron/sqlx"
)

type mariadb struct {
	rw *sqlx.DB
}

// NewMariadb serve caller to create an IRepo
func NewMariadb(rw *sqlx.DB) IRepo {
	return &mariadb{
		rw: rw,
	}
}

func (m *mariadb) GetProfileByUsername(ctx contextx.Contextx, username string) (info *model.Profile, err error) {
	// todo: 2023/4/16|sean|impl me
	panic("implement me")
}

func (m *mariadb) GetProfileByID(ctx contextx.Contextx, id int64) (info *model.Profile, err error) {
	// todo: 2023/4/16|sean|impl me
	panic("implement me")
}

func (m *mariadb) Register(ctx contextx.Contextx, who *model.Profile) error {
	// todo: 2023/4/16|sean|impl me
	panic("implement me")
}
