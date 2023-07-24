package repo

import (
	"database/sql/driver"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	issueM "github.com/blackhorseya/ekko/entity/domain/issue/model"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/blackhorseya/ekko/test/testdata"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

var (
	columns = []string{"id", "title", "status", "created_at", "updated_at"}
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

type suiteMariadb struct {
	suite.Suite
	logger *zap.Logger
	rw     sqlmock.Sqlmock
	repo   IRepo
}

func (s *suiteMariadb) SetupTest() {
	s.logger, _ = zap.NewDevelopment()
	db, mock, _ := sqlmock.New()
	s.rw = mock
	s.repo = CreateMariadb(sqlx.NewDb(db, "mysql"))
}

func (s *suiteMariadb) assertExpectation(t *testing.T) {
	if err := s.rw.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestMariadb(t *testing.T) {
	suite.Run(t, new(suiteMariadb))
}

func (s *suiteMariadb) Test_mariadb_GetByID() {
	type args struct {
		id   int64
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *issueM.Ticket
		wantErr  bool
	}{
		{
			name: "get by id then error",
			args: args{id: testdata.Ticket1.Id, mock: func() {
				s.rw.ExpectQuery(`SELECT id, title, status, created_at, updated_at FROM tickets WHERE id = ?`).
					WithArgs(testdata.Ticket1.Id).
					WillReturnError(errors.New("error"))
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "get by id then not found",
			args: args{id: testdata.Ticket1.Id, mock: func() {
				s.rw.ExpectQuery(`SELECT id, title, status, created_at, updated_at FROM tickets WHERE id = ?`).
					WithArgs(testdata.Ticket1.Id).
					WillReturnRows(sqlmock.NewRows(columns))
			}},
			wantInfo: nil,
			wantErr:  false,
		},
		{
			name: "get by id then ok",
			args: args{id: testdata.Ticket1.Id, mock: func() {
				s.rw.ExpectQuery(`SELECT id, title, status, created_at, updated_at FROM tickets WHERE id = ?`).
					WithArgs(testdata.Ticket1.Id).
					WillReturnRows(sqlmock.NewRows(columns).AddRow(
						testdata.Ticket1.Id,
						testdata.Ticket1.Title,
						testdata.Ticket1.Status,
						testdata.Ticket1.CreatedAt.AsTime(),
						testdata.Ticket1.UpdatedAt.AsTime(),
					))
			}},
			wantInfo: testdata.Ticket1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.repo.GetByID(contextx.WithLogger(s.logger), tt.args.id)
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

func (s *suiteMariadb) Test_mariadb_Create() {
	type args struct {
		created *issueM.Ticket
		mock    func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *issueM.Ticket
		wantErr  bool
	}{
		{
			name: "create then error",
			args: args{created: testdata.Ticket1, mock: func() {
				s.rw.ExpectExec(`insert into tickets`).
					WithArgs(
						testdata.Ticket1.Id,
						testdata.Ticket1.Title,
						testdata.Ticket1.Status,
						AnyTime{},
						AnyTime{},
					).
					WillReturnError(errors.New("error"))
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "create then ok",
			args: args{created: testdata.Ticket1, mock: func() {
				s.rw.ExpectExec(`insert into tickets`).
					WithArgs(
						testdata.Ticket1.Id,
						testdata.Ticket1.Title,
						testdata.Ticket1.Status,
						AnyTime{},
						AnyTime{},
					).
					WillReturnResult(sqlmock.NewResult(1, 1))
			}},
			wantInfo: testdata.Ticket1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.repo.Create(contextx.WithLogger(s.logger), tt.args.created)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("Create() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}

			s.assertExpectation(t)
		})
	}
}

func (s *suiteMariadb) Test_mariadb_List() {
	type args struct {
		condition QueryTicketsCondition
		mock      func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo []*issueM.Ticket
		wantErr  bool
	}{
		{
			name: "list then error",
			args: args{condition: QueryTicketsCondition{}, mock: func() {
				s.rw.ExpectQuery(`select id, title, status, created_at, updated_at from tickets`).
					WillReturnError(errors.New("error"))
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "list then not found",
			args: args{condition: QueryTicketsCondition{}, mock: func() {
				s.rw.ExpectQuery(`select id, title, status, created_at, updated_at from tickets`).
					WillReturnRows(sqlmock.NewRows(columns))
			}},
			wantInfo: nil,
			wantErr:  false,
		},
		{
			name: "list then ok",
			args: args{condition: QueryTicketsCondition{}, mock: func() {
				s.rw.ExpectQuery(`select id, title, status, created_at, updated_at from tickets`).
					WillReturnRows(sqlmock.NewRows(columns).AddRow(
						testdata.Ticket1.Id,
						testdata.Ticket1.Title,
						testdata.Ticket1.Status,
						testdata.Ticket1.CreatedAt.AsTime(),
						testdata.Ticket1.UpdatedAt.AsTime(),
					))
			}},
			wantInfo: []*issueM.Ticket{testdata.Ticket1},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.repo.List(contextx.WithLogger(s.logger), tt.args.condition)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("List() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}

			s.assertExpectation(t)
		})
	}
}

func (s *suiteMariadb) Test_mariadb_DeleteByID() {
	type args struct {
		id   int64
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "delete then error",
			args: args{id: testdata.Ticket1.Id, mock: func() {
				s.rw.ExpectExec(`delete from tickets`).
					WithArgs(testdata.Ticket1.Id).
					WillReturnError(errors.New("error"))
			}},
			wantErr: true,
		},
		{
			name: "delete then ok",
			args: args{id: testdata.Ticket1.Id, mock: func() {
				s.rw.ExpectExec(`delete from tickets`).
					WithArgs(testdata.Ticket1.Id).
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

			if err := s.repo.DeleteByID(contextx.WithLogger(s.logger), tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteByID() error = %v, wantErr %v", err, tt.wantErr)
			}

			s.assertExpectation(t)
		})
	}
}

func (s *suiteMariadb) Test_mariadb_Update() {
	type args struct {
		updated *issueM.Ticket
		mock    func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "update then error",
			args: args{updated: testdata.Ticket1, mock: func() {
				s.rw.ExpectExec(`update tickets`).
					WithArgs(
						testdata.Ticket1.Title,
						testdata.Ticket1.Status,
						AnyTime{},
						testdata.Ticket1.Id,
					).
					WillReturnError(errors.New("error"))
			}},
			wantErr: true,
		},
		{
			name: "update then ok",
			args: args{updated: testdata.Ticket1, mock: func() {
				s.rw.ExpectExec(`update tickets`).
					WithArgs(
						testdata.Ticket1.Title,
						testdata.Ticket1.Status,
						AnyTime{},
						testdata.Ticket1.Id,
					).
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

			err := s.repo.Update(contextx.WithLogger(s.logger), tt.args.updated)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			s.assertExpectation(t)
		})
	}
}
