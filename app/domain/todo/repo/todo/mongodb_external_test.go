//go:build external

package todo

import (
	"github.com/blackhorseya/ekko/entity/domain/todo/repo"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
)

type suiteMongodbExternal struct {
	suite.Suite

	rw   *mongo.Client
	repo repo.ITodoRepo
}
