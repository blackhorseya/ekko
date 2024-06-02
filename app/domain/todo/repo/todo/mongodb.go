package todo

import (
	"time"

	"github.com/blackhorseya/ekko/entity/domain/todo/model"
	"github.com/blackhorseya/ekko/entity/domain/todo/repo"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
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
	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	if item.ID == "" {
		item.ID = uuid.New().String()
	}
	item.UpdatedAt = time.Now()

	coll := i.rw.Database(dbName).Collection(collName)
	_, err = coll.InsertOne(timeout, item)
	if err != nil {
		ctx.Error("mongodb create todo error", zap.Error(err), zap.Any("item", &item))
		return err
	}

	return nil
}

func (i *mongodb) Update(ctx contextx.Contextx, item *model.Todo) (err error) {
	// todo: 2024/6/3|sean|implement me
	panic("implement me")
}
