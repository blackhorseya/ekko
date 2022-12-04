package repo

import (
	"database/sql"

	"github.com/blackhorseya/todo-app/pkg/contextx"
	"github.com/blackhorseya/todo-app/pkg/entity/domain/todo/model"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type mariadb struct {
	rw *sqlx.DB
}

func NewMariadb(rw *sqlx.DB) IRepo {
	return &mariadb{rw: rw}
}

func (i *mariadb) GetByID(ctx contextx.Contextx, id int64) (info *model.Task, err error) {
	stmt := `SELECT id, title, status, created_at, updated_at FROM tickets WHERE id = ?`

	var ret task
	err = i.rw.GetContext(ctx, &ret, stmt, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return ret.ToEntity(), nil
}

func (i *mariadb) List(ctx contextx.Contextx, condition QueryTasksCondition) (info []*model.Task, err error) {
	// TODO implement me
	panic("implement me")
}

func (i *mariadb) Create(ctx contextx.Contextx, created *model.Task) (info *model.Task, err error) {
	// TODO implement me
	panic("implement me")
}

func (i *mariadb) Count(ctx contextx.Contextx, condition QueryTasksCondition) (total int, err error) {
	// TODO implement me
	panic("implement me")
}

func (i *mariadb) Update(ctx contextx.Contextx, updated *model.Task) (info *model.Task, err error) {
	// TODO implement me
	panic("implement me")
}

func (i *mariadb) DeleteByID(ctx contextx.Contextx, id int64) error {
	// TODO implement me
	panic("implement me")
}
