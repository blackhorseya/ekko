package repo

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type suiteMariadb struct {
	suite.Suite
	logger *zap.Logger
	rw     sqlmock.Sqlmock
	repo   IHealthRepo
}

func (s *suiteMariadb) SetupTest() {
	s.logger, _ = zap.NewDevelopment()

	db, mock, _ := sqlmock.New(sqlmock.MonitorPingsOption(true))
	s.rw = mock

	s.repo, _ = CreateMariadb(sqlx.NewDb(db, "mysql"))
}

func TestSuiteMariadb(t *testing.T) {
	suite.Run(t, new(suiteMariadb))
}

func (s *suiteMariadb) Test_mariadb_Ping() {
	type args struct {
		timeout time.Duration
		mock    func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ping then error",
			args: args{timeout: 1 * time.Second, mock: func() {
				s.rw.ExpectPing().WillReturnError(errors.New("error"))
			}},
			wantErr: true,
		},
		{
			name: "ping then timeout",
			args: args{timeout: 1 * time.Millisecond, mock: func() {
				s.rw.ExpectPing().WillDelayFor(2 * time.Millisecond)
			}},
			wantErr: true,
		},
		{
			name: "ping then pong",
			args: args{timeout: 1 * time.Second, mock: func() {
				s.rw.ExpectPing().WillDelayFor(50 * time.Millisecond)
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if err := s.repo.Ping(contextx.BackgroundWithLogger(s.logger), tt.args.timeout); (err != nil) != tt.wantErr {
				t.Errorf("Ping() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err := s.rw.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
