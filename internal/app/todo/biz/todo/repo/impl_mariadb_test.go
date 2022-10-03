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
