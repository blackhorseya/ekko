package repo

import (
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/blackhorseya/todo-app/pkg/contextx"
	"github.com/blackhorseya/todo-app/pkg/entity/domain/todo/model"
	"github.com/blackhorseya/todo-app/test/testdata"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

var (
	columns = []string{"id", "title", "status", "created_at", "updated_at"}
)

type SuiteMariadb struct {
	suite.Suite
	logger *zap.Logger
	rw     sqlmock.Sqlmock
	repo   IRepo
}

func (s *SuiteMariadb) SetupTest() {
	s.logger, _ = zap.NewDevelopment()
	db, mock, _ := sqlmock.New()
	s.rw = mock
	s.repo = CreateMariadb(sqlx.NewDb(db, "mysql"))
}

func (s *SuiteMariadb) assertExpectation(t *testing.T) {
	if err := s.rw.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestMariadb(t *testing.T) {
	suite.Run(t, new(SuiteMariadb))
}

func (s *SuiteMariadb) Test_mariadb_GetByID() {
	type args struct {
		id   int64
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *model.Task
		wantErr  bool
	}{
		{
			name: "get by id then error",
			args: args{id: testdata.Task1.ID, mock: func() {
				s.rw.ExpectQuery(`SELECT id, title, status, created_at, updated_at FROM tickets WHERE id = ?`).
					WithArgs(testdata.Task1.ID).
					WillReturnError(errors.New("error"))
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "get by id then not found",
			args: args{id: testdata.Task1.ID, mock: func() {
				s.rw.ExpectQuery(`SELECT id, title, status, created_at, updated_at FROM tickets WHERE id = ?`).
					WithArgs(testdata.Task1.ID).
					WillReturnRows(sqlmock.NewRows(columns))
			}},
			wantInfo: nil,
			wantErr:  false,
		},
		{
			name: "get by id then ok",
			args: args{id: testdata.Task2.Id, mock: func() {
				s.rw.ExpectQuery(`SELECT id, title, status, created_at, updated_at FROM tickets WHERE id = ?`).
					WithArgs(testdata.Task2.Id).
					WillReturnRows(sqlmock.NewRows(columns).AddRow(
						testdata.Task2.Id,
						testdata.Task2.Title,
						testdata.Task2.Status,
						testdata.Task2.CreatedAt.AsTime(),
						testdata.Task2.UpdatedAt.AsTime(),
					))
			}},
			wantInfo: testdata.Task2,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.repo.GetByID(contextx.BackgroundWithLogger(s.logger), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("GetByID() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}

			s.assertExpectation(t)
		})
	}
}
