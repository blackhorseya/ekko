//go:build external

package mongodb

import (
	"testing"
	"time"

	"github.com/blackhorseya/ekko/entity/domain/workflow/agg"
	"github.com/blackhorseya/ekko/entity/domain/workflow/model"
	"github.com/blackhorseya/ekko/entity/domain/workflow/repo"
	"github.com/blackhorseya/ekko/pkg/configx"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/blackhorseya/ekko/pkg/storage/mongodbx"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type suiteExternal struct {
	suite.Suite

	rw   *mongo.Client
	repo repo.IIssueRepo
}

func (s *suiteExternal) SetupTest() {
	err := configx.LoadWithPathAndName("", "ekko")
	s.Require().NoError(err)

	client, err := mongodbx.NewClient()
	s.Require().NoError(err)
	s.rw = client

	s.repo = NewIssueRepo(s.rw)
}

func (s *suiteExternal) TearDownTest() {
	if s.rw != nil {
		_ = s.rw.Disconnect(contextx.Background())
	}
}

func TestExternal(t *testing.T) {
	suite.Run(t, new(suiteExternal))
}

func (s *suiteExternal) Test_Impl_List() {
	ctx := contextx.Background()

	opts := repo.ListIssueOptions{
		OwnerID: "test",
	}
	items, total, err := s.repo.List(ctx, opts)
	s.Require().NoError(err)

	ctx.Debug("list issue success", zap.Any("items", items), zap.Int("total", total))
}

func (s *suiteExternal) Test_Impl_Create() {
	ctx := contextx.Background()

	item := &agg.Issue{
		Ticket: &model.Ticket{
			ID:        "",
			Title:     "test",
			Completed: false,
			OwnerID:   "test",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
	}
	err := s.repo.Create(ctx, item)
	s.Require().NoError(err)

	ctx.Debug("create issue success", zap.Any("item", &item))
}

func (s *suiteExternal) Test_Impl_GetByID() {
	ctx := contextx.Background()

	item, err := s.repo.GetByID(ctx, "65ed58f20d1152510f7bff43")
	s.Require().NoError(err)

	ctx.Debug("get issue by id success", zap.Any("item", &item))
}

func (s *suiteExternal) Test_Impl_Update() {
	ctx := contextx.Background()

	item := &agg.Issue{
		Ticket: &model.Ticket{
			ID:        "65ed58f20d1152510f7bff43",
			Title:     "test",
			Completed: true,
			OwnerID:   "test",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
	}
	err := s.repo.Update(ctx, item)
	s.Require().NoError(err)

	ctx.Debug("update issue success", zap.Any("item", &item))
}
