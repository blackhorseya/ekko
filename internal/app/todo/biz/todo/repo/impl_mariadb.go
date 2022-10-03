package repo

import (
	"database/sql"
	"strings"
	"time"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/ticket"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type mariadb struct {
	rw *sqlx.DB
}

// NewMariadb serve caller to create an ITodoRepo
func NewMariadb(rw *sqlx.DB) ITodoRepo {
	return &mariadb{
		rw: rw,
	}
}

func (i *mariadb) GetByID(ctx contextx.Contextx, id uint64) (task *ticket.Task, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	stmt := `select id, title, status, created_at, updated_at from tickets where id = ?`

	var ret ticket.Task
	err = i.rw.GetContext(timeout, &ret, stmt, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &ret, nil
}

func (i *mariadb) List(ctx contextx.Contextx, condition QueryTodoCondition) (tasks []*ticket.Task, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var args []interface{}
	stmt := []string{
		`select id, title, status, created_at, updated_at from tickets`,
	}

	if condition.Limit != 0 {
		stmt = append(stmt, `limit ? offset ?`)
		args = append(args, condition.Limit, condition.Offset)
	}

	var ret []*ticket.Task
	err = i.rw.SelectContext(timeout, &ret, strings.Join(stmt, " "), args...)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (i *mariadb) Count(ctx contextx.Contextx, condition QueryTodoCondition) (total int, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	stmt := `select count(id) as total from tickets`

	ret := 0
	row := i.rw.QueryRowxContext(timeout, stmt)
	err = row.Scan(&ret)
	if err != nil {
		return 0, err
	}

	return ret, nil
}

func (i *mariadb) Create(ctx contextx.Contextx, created *ticket.Task) (task *ticket.Task, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	now := time.Now()
	created.CreatedAt = now
	created.UpdatedAt = now

	stmt := `insert into tickets (id, title, status, created_at, updated_at) values (:id, :title, :status, :created_at, :updated_at)`

	_, err = i.rw.NamedExecContext(timeout, stmt, created)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func (i *mariadb) Update(ctx contextx.Contextx, updated *ticket.Task) (task *ticket.Task, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	updated.UpdatedAt = time.Now()

	stmt := `update tickets title=:title, status=:status, updated_at=:updated_at where id = :id`

	_, err = i.rw.NamedExecContext(timeout, stmt, updated)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (i *mariadb) Remove(ctx contextx.Contextx, id uint64) error {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	stmt := `delete from tickets where id = ?`

	_, err := i.rw.ExecContext(timeout, stmt, id)
	if err != nil {
		return err
	}

	return nil
}
