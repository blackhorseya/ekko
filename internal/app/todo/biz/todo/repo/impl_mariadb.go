package repo

import (
	"database/sql"
	"time"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/ticket"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
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

func (i *mariadb) GetByID(ctx contextx.Contextx, id uint64) (task *ticket.Task, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	stmt := `select id, title, status, created_at, updated_at from tickets where id = ?`

	var ret ticket.Task
	err = i.rw.GetContext(timeout, &ret, stmt, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &ret, nil
}

func (i *mariadb) List(ctx contextx.Contextx, condition QueryTodoCondition) (tasks []*ticket.Task, err error) {
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
	// var ret []*ticket.Task
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

func (i *mariadb) Create(ctx contextx.Contextx, created *ticket.Task) (task *ticket.Task, err error) {
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

func (i *mariadb) Update(ctx contextx.Contextx, updated *ticket.Task) (task *ticket.Task, err error) {
	// timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	// defer cancel()
	//
	// updated.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	//
	// filter := bson.M{"_id": updated.ID}
	// opt := options.FindOneAndReplace().SetUpsert(false).SetReturnDocument(options.After)
	// coll := i.client.Database(dbName).Collection(collName)
	// var ret *ticket.Task
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
