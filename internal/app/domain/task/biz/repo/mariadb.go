package repo

import (
	"database/sql"

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
	// todo: 2023/7/30|sean|implement me
	panic("implement me")
}

func (m *mariadb) CountTickets(ctx contextx.Contextx, condition ListTicketsCondition) (total int, err error) {
	// todo: 2023/7/30|sean|implement me
	panic("implement me")
}

func (m *mariadb) CreateTicket(ctx contextx.Contextx, created *taskM.Ticket) (ticket *taskM.Ticket, err error) {
	// todo: 2023/7/30|sean|implement me
	panic("implement me")
}

func (m *mariadb) UpdateTicket(ctx contextx.Contextx, updated *taskM.Ticket) error {
	// todo: 2023/7/30|sean|implement me
	panic("implement me")
}

func (m *mariadb) DeleteTicketByID(ctx contextx.Contextx, id string) error {
	// todo: 2023/7/30|sean|implement me
	panic("implement me")
}
