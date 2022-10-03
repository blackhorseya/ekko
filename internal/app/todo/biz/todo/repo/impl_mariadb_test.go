package repo

import (
	"database/sql/driver"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/ticket"
	"github.com/blackhorseya/todo-app/test/testdata"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

var (
	columns = []string{"id", "title", "status", "created_at", "updated_at"}
)

type suiteMariadb struct {
	suite.Suite
	logger *zap.Logger
	rw     sqlmock.Sqlmock
	repo   ITodoRepo
}

func (s *suiteMariadb) SetupTest() {
	s.logger, _ = zap.NewDevelopment()

	db, mock, _ := sqlmock.New()
	s.rw = mock

	s.repo, _ = CreateMariadb(sqlx.NewDb(db, "mysql"))
}

func TestSuiteMariadb(t *testing.T) {
	suite.Run(t, new(suiteMariadb))
}

func (s *suiteMariadb) Test_mariadb_GetByID() {
	type args struct {
		id   uint64
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantTask *ticket.Task
		wantErr  bool
	}{
		{
			name: "get by id then error",
			args: args{id: testdata.Task1.ID, mock: func() {
				s.rw.ExpectQuery(`select id, title, status, created_at, updated_at from tickets where id = ?`).
					WithArgs(testdata.Task1.ID).
					WillReturnError(errors.New("error"))
			}},
			wantTask: nil,
			wantErr:  true,
		},
		{
			name: "get by id then not found",
			args: args{id: testdata.Task1.ID, mock: func() {
				s.rw.ExpectQuery(`select id, title, status, created_at, updated_at from tickets where id = ?`).
					WithArgs(testdata.Task1.ID).
					WillReturnRows(sqlmock.NewRows(columns))
			}},
			wantTask: nil,
			wantErr:  false,
		},
		{
			name: "get by id then ok",
			args: args{id: testdata.Task1.ID, mock: func() {
				s.rw.ExpectQuery(`select id, title, status, created_at, updated_at from tickets where id = ?`).
					WithArgs(testdata.Task1.ID).
					WillReturnRows(sqlmock.NewRows(columns).AddRow(
						testdata.Task1.ID,
						testdata.Task1.Title,
						testdata.Task1.Status,
						testdata.Task1.CreatedAt,
						testdata.Task1.UpdatedAt,
					))
			}},
			wantTask: testdata.Task1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotTask, err := s.repo.GetByID(contextx.BackgroundWithLogger(s.logger), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTask, tt.wantTask) {
				t.Errorf("GetByID() gotTask = %v, want %v", gotTask, tt.wantTask)
			}

			if err := s.rw.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
