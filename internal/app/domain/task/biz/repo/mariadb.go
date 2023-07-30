package repo

import (
	"github.com/blackhorseya/ekko/entity/domain/task/model"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/jmoiron/sqlx"
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

func (m *mariadb) GetTicketByID(ctx contextx.Contextx, id string) (ticket *model.Ticket, err error) {
	// todo: 2023/7/30|sean|implement me
	panic("implement me")
}

func (m *mariadb) ListTickets(ctx contextx.Contextx, condition ListTicketsCondition) (tickets []*model.Ticket, total int, err error) {
	// todo: 2023/7/30|sean|implement me
	panic("implement me")
}

func (m *mariadb) CountTickets(ctx contextx.Contextx, condition ListTicketsCondition) (total int, err error) {
	// todo: 2023/7/30|sean|implement me
	panic("implement me")
}

func (m *mariadb) CreateTicket(ctx contextx.Contextx, created *model.Ticket) (ticket *model.Ticket, err error) {
	// todo: 2023/7/30|sean|implement me
	panic("implement me")
}

func (m *mariadb) UpdateTicket(ctx contextx.Contextx, updated *model.Ticket) error {
	// todo: 2023/7/30|sean|implement me
	panic("implement me")
}

func (m *mariadb) DeleteTicketByID(ctx contextx.Contextx, id string) error {
	// todo: 2023/7/30|sean|implement me
	panic("implement me")
}
