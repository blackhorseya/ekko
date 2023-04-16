package repo

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
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
