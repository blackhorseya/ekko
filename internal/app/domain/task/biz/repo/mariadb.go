package repo

import (
	"database/sql"
	"fmt"

	taskM "github.com/blackhorseya/ekko/entity/domain/task/model"
	"github.com/blackhorseya/ekko/internal/app/domain/task/biz/repo/dao"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type mariadb struct {
	rw *sqlx.DB
}

// NewMariadb will create an object that represent the IRepo interface
func NewMariadb(rw *sqlx.DB) IRepo {
	return &mariadb{
		rw: rw,
	}
}

func (m *mariadb) GetTicketByID(ctx contextx.Contextx, id string) (ticket *taskM.Ticket, err error) {
	stmt := `SELECT id, title, status, created_at, updated_at FROM tickets WHERE id = ?`

	var got dao.Ticket
	err = m.rw.GetContext(ctx, &got, stmt, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return got.ToEntity(), nil
}

func (m *mariadb) ListTickets(ctx contextx.Contextx, condition ListTicketsCondition) (tickets []*taskM.Ticket, total int, err error) {
	stmt := `SELECT id, title, status, created_at, updated_at FROM tickets`
	count := fmt.Sprintf(`SELECT COUNT(*) FROM (%s) AS total`, stmt)

	err = m.rw.GetContext(ctx, &total, count)
	if err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return nil, 0, nil
	}

	stmt = fmt.Sprintf("%s LIMIT ? OFFSET ?", stmt)

	var got dao.Tickets
	err = m.rw.SelectContext(ctx, &got, stmt, condition.Limit, condition.Offset)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, total, nil
		}

		return nil, 0, err
	}

	return got.ToEntity(), total, nil
}

func (m *mariadb) CreateTicket(ctx contextx.Contextx, created *taskM.Ticket) (ticket *taskM.Ticket, err error) {
	stmt := `INSERT INTO tickets (id, title, status, created_at, updated_at) VALUES (:id, :title, :status, :created_at, :updated_at)`

	arg := dao.NewTicket(created)
	_, err = m.rw.NamedExecContext(ctx, stmt, arg)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func (m *mariadb) UpdateTicket(ctx contextx.Contextx, updated *taskM.Ticket) error {
	stmt := `UPDATE tickets SET title = :title, status = :status, updated_at = :updated_at WHERE id = :id`

	arg := dao.NewTicket(updated)
	_, err := m.rw.NamedExecContext(ctx, stmt, arg)
	if err != nil {
		return err
	}

	return nil
}

func (m *mariadb) DeleteTicketByID(ctx contextx.Contextx, id string) error {
	stmt := `DELETE FROM tickets WHERE id = ?`

	_, err := m.rw.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	return nil
}
