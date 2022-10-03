package repo

import (
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/todo"
	"github.com/jmoiron/sqlx"
)

type mariadb struct {
	rw *sqlx.DB
}

// NewMariadb serve caller to create an ITodoRepo
func NewMariadb(rw *sqlx.DB) ITodoRepo {
	return &mariadb{
		rw: rw,
	}
}

func (i *mariadb) GetByID(ctx contextx.Contextx, id uint64) (task *todo.Task, err error) {
	// timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	// defer cancel()
	//
	// filter := bson.M{"_id": id}
	// var ret todo.Task
	// coll := i.client.Database(dbName).Collection(collName)
	// err = coll.FindOne(timeout, filter).Decode(&ret)
	// if err != nil {
	// 	if err == mongo.ErrNoDocuments {
	// 		return nil, nil
	// 	}
	//
	// 	return nil, err
	// }
	//
	// return &ret, nil
	// todo: 2022/10/4|sean|mariadb me
	panic("mariadb me")
}

func (i *mariadb) List(ctx contextx.Contextx, condition QueryTodoCondition) (tasks []*todo.Task, err error) {
	// timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	// defer cancel()
	//
	// coll := i.client.Database(dbName).Collection(collName)
	// cur, err := coll.Find(timeout, bson.D{}, options.Find().SetLimit(int64(limit)).SetSkip(int64(offset)))
	// if err != nil {
	// 	return nil, err
	// }
	// defer cur.Close(timeout)
	//
	// var ret []*todo.Task
	// err = cur.All(timeout, &ret)
	// if err != nil {
	// 	if errors.Is(err, mongo.ErrNoDocuments) {
	// 		return nil, nil
	// 	}
	//
	// 	return nil, err
	// }
	//
	// return ret, nil

	// todo: 2022/10/4|sean|mariadb me
	panic("mariadb me")
}

func (i *mariadb) Count(ctx contextx.Contextx) (total int, err error) {
	// timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	// defer cancel()
	//
	// coll := i.client.Database(dbName).Collection(collName)
	// ret, err := coll.CountDocuments(timeout, bson.M{})
	// if err != nil {
	// 	return 0, err
	// }
	//
	// return int(ret), nil

	// todo: 2022/10/4|sean|mariadb me
	panic("mariadb me")
}

func (i *mariadb) Create(ctx contextx.Contextx, created *todo.Task) (task *todo.Task, err error) {
	// timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	// defer cancel()
	//
	// now := time.Now()
	// created.ID = primitive.NewObjectIDFromTimestamp(now)
	// created.CreatedAt = primitive.NewDateTimeFromTime(now)
	// created.UpdatedAt = primitive.NewDateTimeFromTime(now)
	//
	// coll := i.client.Database(dbName).Collection(collName)
	// res, err := coll.InsertOne(timeout, created)
	// if err != nil {
	// 	return nil, err
	// }
	//
	// created.ID = res.InsertedID.(primitive.ObjectID)
	//
	// return created, nil

	// todo: 2022/10/4|sean|mariadb me
	panic("mariadb me")
}

func (i *mariadb) Update(ctx contextx.Contextx, updated *todo.Task) (task *todo.Task, err error) {
	// timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	// defer cancel()
	//
	// updated.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	//
	// filter := bson.M{"_id": updated.ID}
	// opt := options.FindOneAndReplace().SetUpsert(false).SetReturnDocument(options.After)
	// coll := i.client.Database(dbName).Collection(collName)
	// var ret *todo.Task
	// err = coll.FindOneAndReplace(timeout, filter, updated, opt).Decode(&ret)
	// if err != nil {
	// 	return nil, err
	// }
	//
	// return ret, nil

	// todo: 2022/10/4|sean|mariadb me
	panic("mariadb me")
}

func (i *mariadb) Remove(ctx contextx.Contextx, id uint64) error {
	// timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	// defer cancel()
	//
	// coll := i.client.Database(dbName).Collection(collName)
	// _, err := coll.DeleteOne(timeout, bson.M{"_id": id})
	// if err != nil {
	// 	return err
	// }
	//
	// return nil

	// todo: 2022/10/4|sean|mariadb me
	panic("mariadb me")
}
