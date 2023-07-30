package repo

import (
	"database/sql"
	"reflect"
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
	s.repo = NewMariadb(sqlx.NewDb(db, "mysql"))
}

func (s *SuiteMariadb) TearDownSuite() {
	err := s.rw.ExpectationsWereMet()
	if err != nil {
		s.T().Errorf("there were unfulfilled expectations: %s", err)
	}
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
							id1,
							ticket1.Title,
							ticket1.Status,
							now,
							now,
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
