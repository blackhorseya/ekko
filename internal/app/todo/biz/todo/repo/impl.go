package repo

import (
	"github.com/blackhorseya/todo-app/internal/pkg/base/contextx"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/todo"
	"go.mongodb.org/mongo-driver/mongo"
)

type impl struct {
	client *mongo.Client
}

// NewImpl serve caller to create an IRepo
func NewImpl(client *mongo.Client) IRepo {
	return &impl{client: client}
}

func (i *impl) GetByID(ctx contextx.Contextx, id string) (task *todo.Task, err error) {
	// todo: 2021-05-02|10:15|doggy|implement me
	panic("implement me")
}

func (i *impl) List(ctx contextx.Contextx, limit, offset int) (tasks []*todo.Task, err error) {
	// todo: 2021-05-02|10:15|doggy|implement me
	panic("implement me")
}

func (i *impl) Count(ctx contextx.Contextx) (total int, err error) {
	// todo: 2021-05-02|10:15|doggy|implement me
	panic("implement me")
}

func (i *impl) Create(ctx contextx.Contextx, title string) (task *todo.Task, err error) {
	panic("implement me")
}

func (i *impl) Update(ctx contextx.Contextx, updated *todo.Task) (task *todo.Task, err error) {
	// todo: 2021-05-02|10:15|doggy|implement me
	panic("implement me")
}

func (i *impl) Remove(ctx contextx.Contextx, id string) error {
	// todo: 2021-05-02|10:15|doggy|implement me
	panic("implement me")
}
