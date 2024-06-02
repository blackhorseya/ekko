package todo

import (
	"time"

	"github.com/blackhorseya/ekko/entity/domain/todo/model"
	"github.com/blackhorseya/ekko/entity/domain/todo/repo"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	defaultTimeout = 5 * time.Second
	dbName         = "ekko"
	collName       = "todos"
)

type mongodb struct {
	rw *mongo.Client
}

// NewMongodb is to create a new mongodb
func NewMongodb(rw *mongo.Client) repo.ITodoRepo {
	return &mongodb{rw: rw}
}

func (i *mongodb) List(
	ctx contextx.Contextx,
	condition repo.ListCondition,
) (items []*model.Todo, total int, err error) {
	// todo: 2024/6/3|sean|implement me
	panic("implement me")
}

func (i *mongodb) GetByID(ctx contextx.Contextx, id string) (item *model.Todo, err error) {
	// todo: 2024/6/3|sean|implement me
	panic("implement me")
}

func (i *mongodb) Create(ctx contextx.Contextx, item *model.Todo) (err error) {
	// todo: 2024/6/3|sean|implement me
	panic("implement me")
}

func (i *mongodb) Update(ctx contextx.Contextx, item *model.Todo) (err error) {
	// todo: 2024/6/3|sean|implement me
	panic("implement me")
}
