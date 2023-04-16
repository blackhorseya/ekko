package biz

import (
	"testing"

	"github.com/blackhorseya/ekko/internal/app/domain/user/biz/repo"
	ub "github.com/blackhorseya/ekko/pkg/entity/domain/user/biz"
	"github.com/blackhorseya/ekko/pkg/genx"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type suiteBiz struct {
	suite.Suite

	logger *zap.Logger
	ctrl   *gomock.Controller

	repo *repo.MockIRepo
	node *genx.MockGenerator
	biz  ub.IBiz
}

func (s *suiteBiz) SetupTest() {
	s.logger, _ = zap.NewDevelopment()
	s.ctrl = gomock.NewController(s.T())

	s.repo = repo.NewMockIRepo(s.ctrl)
	s.node = genx.NewMockGenerator(s.ctrl)

	s.biz = CreateBiz(s.repo, s.node)
}

func (s *suiteBiz) TearDownTest() {
	s.ctrl.Finish()
}

func TestAll(t *testing.T) {
	suite.Run(t, new(suiteBiz))
}
