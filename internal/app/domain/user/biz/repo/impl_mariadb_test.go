package repo

import (
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/blackhorseya/ekko/pkg/contextx"
	um "github.com/blackhorseya/ekko/pkg/entity/domain/user/model"
	"github.com/blackhorseya/ekko/test/testdata"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type suiteMariadb struct {
	suite.Suite

	logger *zap.Logger
	rw     sqlmock.Sqlmock
	repo   IRepo
}

func (s *suiteMariadb) SetupTest() {
	s.logger, _ = zap.NewDevelopment()
	db, rw, _ := sqlmock.New()
	s.rw = rw
	s.repo = NewRepoByMariadb(sqlx.NewDb(db, "mysql"))
}

func (s *suiteMariadb) AssertExpectation(t *testing.T) {
	if err := s.rw.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestMariadb(t *testing.T) {
	suite.Run(t, new(suiteMariadb))
}

func (s *suiteMariadb) Test_mariadb_GetProfileByUsername() {
	column := []string{"id", "username", "password", "token", "created_at", "updated_at"}

	type args struct {
		username string
		mock     func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *um.Profile
		wantErr  bool
	}{
		{
			name: "get profile by username then error",
			args: args{username: testdata.Profile1.Username, mock: func() {
				s.rw.ExpectQuery(`select id, username, password, token, created_at, updated_at from users where username = ?`).
					WithArgs(testdata.Profile1.Username).
					WillReturnError(errors.New("error"))
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "if not found then return nil",
			args: args{username: testdata.Profile1.Username, mock: func() {
				s.rw.ExpectQuery(`select id, username, password, token, created_at, updated_at from users where username = ?`).
					WithArgs(testdata.Profile1.Username).
					WillReturnRows(sqlmock.NewRows(column))
			}},
			wantInfo: nil,
			wantErr:  false,
		},
		{
			name: "if found then return profile",
			args: args{username: testdata.Profile1.Username, mock: func() {
				s.rw.ExpectQuery(`select id, username, password, token, created_at, updated_at from users where username = ?`).
					WithArgs(testdata.Profile1.Username).
					WillReturnRows(sqlmock.NewRows(column).AddRow(
						testdata.Profile1.Id,
						testdata.Profile1.Username,
						testdata.Profile1.Password,
						testdata.Profile1.Token,
						testdata.Profile1.CreatedAt.AsTime(),
						testdata.Profile1.UpdatedAt.AsTime(),
					))
			}},
			wantInfo: testdata.Profile1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.repo.GetProfileByUsername(contextx.BackgroundWithLogger(s.logger), tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProfileByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("GetProfileByUsername() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}

			s.AssertExpectation(t)
		})
	}
}
