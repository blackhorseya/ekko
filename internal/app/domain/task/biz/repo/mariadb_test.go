package repo

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

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
