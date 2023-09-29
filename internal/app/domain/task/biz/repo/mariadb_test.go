package repo

import (
	"database/sql"
	"fmt"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	taskM "github.com/blackhorseya/ekko/entity/domain/task/model"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var now = time.Now()

type SuiteMariadb struct {
	suite.Suite

	logger *zap.Logger
	rw     sqlmock.Sqlmock
	repo   IRepo
}

func (s *SuiteMariadb) SetupTest() {
	s.logger = zap.NewExample()
	db, mock, err := sqlmock.New()
	if err != nil {
		s.T().Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	s.rw = mock
	s.repo, _ = NewMariadb(sqlx.NewDb(db, "mysql"))
}

func TestMariadb(t *testing.T) {
	t.Parallel()
	t.Helper()

	suite.Run(t, new(SuiteMariadb))
}

func (s *SuiteMariadb) Test_mariadb_GetTicketByID() {
	id1 := "1"
	ticket1 := &taskM.Ticket{
		Id:          id1,
		Title:       "title",
		Description: "",
		Status:      taskM.TicketStatus_TICKET_STATUS_TODO,
		CreatedAt:   timestamppb.New(now),
		UpdatedAt:   timestamppb.New(now),
	}

	column := []string{"id", "title", "status", "created_at", "updated_at"}
	stmt := `SELECT id, title, status, created_at, updated_at FROM tickets WHERE id = ?`

	type args struct {
		id   string
		mock func()
	}
	tests := []struct {
		name       string
		args       args
		wantTicket *taskM.Ticket
		wantErr    bool
	}{
		{
			name: "query then error",
			args: args{id: id1, mock: func() {
				s.rw.ExpectQuery(stmt).
					WithArgs(id1).
					WillReturnError(errors.New("error"))
			}},
			wantTicket: nil,
			wantErr:    true,
		},
		{
			name: "not found return nil",
			args: args{id: id1, mock: func() {
				s.rw.ExpectQuery(stmt).
					WithArgs(id1).
					WillReturnError(sql.ErrNoRows)
			}},
			wantTicket: nil,
			wantErr:    false,
		},
		{
			name: "get ticket then ok",
			args: args{id: id1, mock: func() {
				s.rw.ExpectQuery(stmt).
					WithArgs(id1).
					WillReturnRows(sqlmock.NewRows(column).
						AddRow(
							ticket1.Id,
							ticket1.Title,
							ticket1.Status,
							ticket1.CreatedAt.AsTime(),
							ticket1.UpdatedAt.AsTime(),
						),
					)
			}},
			wantTicket: ticket1,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotTicket, err := s.repo.GetTicketByID(contextx.WithLogger(s.logger), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTicketByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTicket, tt.wantTicket) {
				t.Errorf("GetTicketByID() gotTicket = %v, want %v", gotTicket, tt.wantTicket)
			}
		})
	}
}

func (s *SuiteMariadb) Test_mariadb_ListTickets() {
	condition := ListTicketsCondition{Limit: 10, Offset: 0}
	ticket1 := &taskM.Ticket{
		Id:          "1",
		Title:       "title",
		Description: "",
		Status:      taskM.TicketStatus_TICKET_STATUS_TODO,
		CreatedAt:   timestamppb.New(now),
		UpdatedAt:   timestamppb.New(now),
	}
	tickets := []*taskM.Ticket{ticket1}

	column := []string{"id", "title", "status", "created_at", "updated_at"}
	stmt := `SELECT id, title, status, created_at, updated_at FROM tickets`
	count := fmt.Sprintf(`SELECT COUNT(*) FROM (%s) AS total`, stmt)

	type args struct {
		mock func()
	}
	tests := []struct {
		name        string
		args        args
		wantTickets []*taskM.Ticket
		wantTotal   int
		wantErr     bool
	}{
		{
			name: "count all then error",
			args: args{mock: func() {
				s.rw.ExpectQuery(regexp.QuoteMeta(count)).WillReturnError(errors.New("error"))
			}},
			wantTickets: nil,
			wantTotal:   0,
			wantErr:     true,
		},
		{
			name: "count all got 0 return nil",
			args: args{mock: func() {
				s.rw.ExpectQuery(regexp.QuoteMeta(count)).WillReturnRows(sqlmock.NewRows([]string{"total"}).AddRow(0))
			}},
			wantTickets: nil,
			wantTotal:   0,
			wantErr:     false,
		},
		{
			name: "select then error",
			args: args{mock: func() {
				s.rw.ExpectQuery(regexp.QuoteMeta(count)).WillReturnRows(sqlmock.NewRows([]string{"total"}).AddRow(10))
				s.rw.ExpectQuery(regexp.QuoteMeta(stmt)).
					WithArgs(condition.Limit, condition.Offset).
					WillReturnError(errors.New("error"))
			}},
			wantTickets: nil,
			wantTotal:   0,
			wantErr:     true,
		},
		{
			name: "select not found then return nil",
			args: args{mock: func() {
				s.rw.ExpectQuery(regexp.QuoteMeta(count)).WillReturnRows(sqlmock.NewRows([]string{"total"}).AddRow(10))
				s.rw.ExpectQuery(regexp.QuoteMeta(stmt)).
					WithArgs(condition.Limit, condition.Offset).
					WillReturnError(sql.ErrNoRows)
			}},
			wantTickets: nil,
			wantTotal:   10,
			wantErr:     false,
		},
		{
			name: "select then ok",
			args: args{mock: func() {
				s.rw.ExpectQuery(regexp.QuoteMeta(count)).WillReturnRows(sqlmock.NewRows([]string{"total"}).AddRow(100))
				s.rw.ExpectQuery(regexp.QuoteMeta(stmt)).
					WithArgs(condition.Limit, condition.Offset).
					WillReturnRows(sqlmock.NewRows(column).
						AddRow(
							ticket1.Id,
							ticket1.Title,
							ticket1.Status,
							ticket1.CreatedAt.AsTime(),
							ticket1.UpdatedAt.AsTime(),
						),
					)
			}},
			wantTickets: tickets,
			wantTotal:   100,
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotTickets, gotTotal, err := s.repo.ListTickets(contextx.WithLogger(s.logger), condition)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListTickets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTickets, tt.wantTickets) {
				t.Errorf("ListTickets() gotTickets = %v, want %v", gotTickets, tt.wantTickets)
			}
			if gotTotal != tt.wantTotal {
				t.Errorf("ListTickets() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
			err = s.rw.ExpectationsWereMet()
			if err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func (s *SuiteMariadb) Test_mariadb_CreateTicket() {
	ticket1 := &taskM.Ticket{
		Id:          "1",
		Title:       "title",
		Description: "",
		Status:      taskM.TicketStatus_TICKET_STATUS_TODO,
		CreatedAt:   timestamppb.New(now),
		UpdatedAt:   timestamppb.New(now),
	}
	stmt := `INSERT INTO tickets (id, title, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`

	type args struct {
		created *taskM.Ticket
		mock    func()
	}
	tests := []struct {
		name       string
		args       args
		wantTicket *taskM.Ticket
		wantErr    bool
	}{
		{
			name: "create then error",
			args: args{created: ticket1, mock: func() {
				s.rw.ExpectExec(regexp.QuoteMeta(stmt)).WithArgs(
					ticket1.Id,
					ticket1.Title,
					ticket1.Status,
					ticket1.CreatedAt.AsTime(),
					ticket1.UpdatedAt.AsTime(),
				).WillReturnError(errors.New("error"))
			}},
			wantTicket: nil,
			wantErr:    true,
		},
		{
			name: "create then ok",
			args: args{created: ticket1, mock: func() {
				s.rw.ExpectExec(regexp.QuoteMeta(stmt)).WithArgs(
					ticket1.Id,
					ticket1.Title,
					ticket1.Status,
					ticket1.CreatedAt.AsTime(),
					ticket1.UpdatedAt.AsTime(),
				).WillReturnResult(sqlmock.NewResult(1, 1))
			}},
			wantTicket: ticket1,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotTicket, err := s.repo.CreateTicket(contextx.WithLogger(s.logger), tt.args.created)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTicket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTicket, tt.wantTicket) {
				t.Errorf("CreateTicket() gotTicket = %v, want %v", gotTicket, tt.wantTicket)
			}
			err = s.rw.ExpectationsWereMet()
			if err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func (s *SuiteMariadb) Test_mariadb_UpdateTicket() {
	ticket1 := &taskM.Ticket{
		Id:          "1",
		Title:       "title",
		Description: "",
		Status:      taskM.TicketStatus_TICKET_STATUS_TODO,
		CreatedAt:   timestamppb.New(now),
		UpdatedAt:   timestamppb.New(now),
	}
	stmt := `UPDATE tickets SET title = ?, status = ?, updated_at = ? WHERE id = ?`

	type args struct {
		updated *taskM.Ticket
		mock    func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "update then error",
			args: args{updated: ticket1, mock: func() {
				s.rw.ExpectExec(regexp.QuoteMeta(stmt)).
					WithArgs(ticket1.Title, ticket1.Status, ticket1.UpdatedAt.AsTime(), ticket1.Id).
					WillReturnError(errors.New("error"))
			}},
			wantErr: true,
		},
		{
			name: "update then error",
			args: args{updated: ticket1, mock: func() {
				s.rw.ExpectExec(regexp.QuoteMeta(stmt)).
					WithArgs(ticket1.Title, ticket1.Status, ticket1.UpdatedAt.AsTime(), ticket1.Id).
					WillReturnResult(sqlmock.NewResult(1, 1))
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if err := s.repo.UpdateTicket(contextx.WithLogger(s.logger), tt.args.updated); (err != nil) != tt.wantErr {
				t.Errorf("UpdateTicket() error = %v, wantErr %v", err, tt.wantErr)
			}
			err := s.rw.ExpectationsWereMet()
			if err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func (s *SuiteMariadb) Test_mariadb_DeleteTicketByID() {
	id1 := "1"

	stmt := `DELETE FROM tickets WHERE id = ?`

	type args struct {
		id   string
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "delete by id then error",
			args: args{id: id1, mock: func() {
				s.rw.ExpectExec(regexp.QuoteMeta(stmt)).
					WithArgs(id1).
					WillReturnError(errors.New("error"))
			}},
			wantErr: true,
		},
		{
			name: "delete by id then ok",
			args: args{id: id1, mock: func() {
				s.rw.ExpectExec(regexp.QuoteMeta(stmt)).
					WithArgs(id1).
					WillReturnResult(sqlmock.NewResult(1, 1))
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if err := s.repo.DeleteTicketByID(contextx.WithLogger(s.logger), tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteTicketByID() error = %v, wantErr %v", err, tt.wantErr)
			}
			err := s.rw.ExpectationsWereMet()
			if err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
