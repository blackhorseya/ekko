package biz

import (
	"testing"

	taskB "github.com/blackhorseya/ekko/entity/domain/task/biz"
	taskR "github.com/blackhorseya/ekko/internal/app/domain/task/biz/repo"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type SuiteTester struct {
	suite.Suite

	logger *zap.Logger
	ctrl   *gomock.Controller
	repo   *taskR.MockIRepo
	biz    taskB.IBiz
}

func (s *SuiteTester) SetupTest() {
	s.logger = zap.NewExample()
	s.ctrl = gomock.NewController(s.T())
	s.repo = taskR.NewMockIRepo(s.ctrl)
	s.biz = NewImpl(s.repo)
}

func (s *SuiteTester) TearDownTest() {
	s.ctrl.Finish()
}

func TestAll(t *testing.T) {
	t.Parallel()
	t.Helper()

	suite.Run(t, new(SuiteTester))
}
