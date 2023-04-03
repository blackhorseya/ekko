package repo

import (
	"database/sql"
	"strings"

	"github.com/blackhorseya/ekko/pkg/contextx"
	tm "github.com/blackhorseya/ekko/pkg/entity/domain/task/model"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type mariadb struct {
	rw *sqlx.DB
}

func NewMariadb(rw *sqlx.DB) IRepo {
	return &mariadb{rw: rw}
}

func (i *mariadb) GetByID(ctx contextx.Contextx, id int64) (info *tm.Ticket, err error) {
	stmt := `SELECT id, title, status, created_at, updated_at FROM tickets WHERE id = ?`

	var ret task
	err = i.rw.GetContext(ctx, &ret, stmt, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		ctx.Error("GetByID", zap.Error(err), zap.String("stmt", stmt), zap.Int64("id", id))
		return nil, err
	}

	return ret.ToEntity(), nil
}

func (i *mariadb) Create(ctx contextx.Contextx, created *tm.Ticket) (info *tm.Ticket, err error) {
	now := timestamppb.Now()
	created.CreatedAt = now
	created.UpdatedAt = now
	arg := newTask(created)

	stmt := `insert into tickets (id, title, status, created_at, updated_at) values (:id, :title, :status, :created_at, :updated_at)`

	_, err = i.rw.NamedExecContext(ctx, stmt, arg)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func (i *mariadb) List(ctx contextx.Contextx, condition QueryTicketsCondition) (info []*tm.Ticket, err error) {
	var args []interface{}
	query := []string{
		`select id, title, status, created_at, updated_at from tickets`,
	}

	if condition.Limit != 0 {
		query = append(query, `limit ?`)
		args = append(args, condition.Limit)
	}

	if condition.Offset != 0 {
		query = append(query, `offset ?`)
		args = append(args, condition.Offset)
	}

	// concat query to stmt
	stmt := strings.Join(query, " ")

	var got []*task
	err = i.rw.SelectContext(ctx, &got, stmt, args...)
	if err != nil {
		return nil, err
	}
	if len(got) == 0 {
		return nil, nil
	}

	ret := make([]*tm.Ticket, len(got))
	for idx, t := range got {
		ret[idx] = t.ToEntity()
	}

	return ret, nil
}

func (i *mariadb) Count(ctx contextx.Contextx, condition QueryTicketsCondition) (total int, err error) {
	stmt := `select count(id) as total from tickets`

	ret := 0
	err = i.rw.QueryRowxContext(ctx, stmt).Scan(&ret)
	if err != nil {
		return 0, err
	}

	return ret, nil
}

func (i *mariadb) Update(ctx contextx.Contextx, updated *tm.Ticket) (info *tm.Ticket, err error) {
	updated.UpdatedAt = timestamppb.Now()

	stmt := `update tickets set title=:title, status=:status, updated_at=:updated_at where id = :id`

	_, err = i.rw.NamedExecContext(ctx, stmt, newTask(updated))
	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (i *mariadb) DeleteByID(ctx contextx.Contextx, id int64) error {
	stmt := `delete from tickets where id = ?`

	_, err := i.rw.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	return nil
}
