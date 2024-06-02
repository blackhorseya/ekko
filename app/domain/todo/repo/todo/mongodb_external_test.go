//go:build external

package todo

import (
	"testing"

	"github.com/blackhorseya/ekko/app/infra/configx"
	"github.com/blackhorseya/ekko/entity/domain/todo/model"
	"github.com/blackhorseya/ekko/entity/domain/todo/repo"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/blackhorseya/ekko/pkg/storage/mongodbx"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
)

type suiteMongodbExternal struct {
	suite.Suite

	rw   *mongo.Client
	repo repo.ITodoRepo
}

func (s *suiteMongodbExternal) SetupTest() {
	err := configx.LoadConfig("")
	s.Require().NoError(err)

	app, err := configx.LoadApplication(&configx.C.PlatformRest)
	s.Require().NoError(err)

	rw, err := mongodbx.NewClientWithDSN(app.Storage.Mongodb.DSN)
	s.Require().NoError(err)
	s.rw = rw

	s.repo = NewMongodb(s.rw)
}

func (s *suiteMongodbExternal) TearDownTest() {
	if s.rw != nil {
		_ = s.rw.Disconnect(contextx.Background())
	}
}

func TestMongodbExternal(t *testing.T) {
	suite.Run(t, new(suiteMongodbExternal))
}

func (s *suiteMongodbExternal) Test_mongodb_Create() {
	todo := model.NewTodo("test")
	err := s.repo.Create(contextx.Background(), todo)
	s.NoError(err)
}
