package repo

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/gocommon/pkg/restclient"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/todo"
	"github.com/blackhorseya/todo-app/test/testdata"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
)

type suiteRepo struct {
	suite.Suite
	logger     *zap.Logger
	pool       *dockertest.Pool
	resource   *dockertest.Resource
	client     *mongo.Client
	restclient *restclient.MockRestClient
	repo       IRepo
}

func (s *suiteRepo) SetupTest() {
	s.logger, _ = zap.NewDevelopment()

	s.restclient = new(restclient.MockRestClient)

	pool, err := dockertest.NewPool("")
	if err != nil {
		panic(err)
	}
	s.pool = pool

	resource, err := pool.Run("mongo", "4.4.10", nil)
	if err != nil {
		panic(err)
	}
	s.resource = resource

	err = pool.Retry(func() error {
		uri := fmt.Sprintf("mongodb://localhost:%s/", resource.GetPort("27017/tcp"))
		s.client, err = mongo.Connect(contextx.BackgroundWithLogger(s.logger), options.Client().ApplyURI(uri))
		if err != nil {
			return err
		}

		return s.client.Ping(contextx.BackgroundWithLogger(s.logger), readpref.Primary())
	})
	if err != nil {
		panic(err)
	}

	repo, err := CreateIRepo(s.client, s.restclient)
	if err != nil {
		panic(err)
	}
	s.repo = repo
}

func (s *suiteRepo) TearDownTest() {
	_ = s.client.Disconnect(contextx.BackgroundWithLogger(s.logger))
	_ = s.pool.Purge(s.resource)
}

func TestSuiteRepo(t *testing.T) {
	suite.Run(t, new(suiteRepo))
}

func (s *suiteRepo) Test_impl_GetByID() {
	type args struct {
		id   primitive.ObjectID
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantTask *todo.Task
		wantErr  bool
	}{
		{
			name:     "get by id then not found",
			args:     args{id: testdata.Task1.ID},
			wantTask: nil,
			wantErr:  false,
		},
		{
			name: "get by id then success",
			args: args{id: testdata.Task1.ID, mock: func() {
				_, _ = s.client.Database(dbName).Collection(collName).InsertOne(contextx.BackgroundWithLogger(s.logger), testdata.Task1)
			}},
			wantTask: testdata.Task1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotTask, err := s.repo.GetByID(contextx.BackgroundWithLogger(s.logger), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTask, tt.wantTask) {
				t.Errorf("GetByID() gotTask = %v, want %v", gotTask, tt.wantTask)
			}

			_, _ = s.client.Database(dbName).Collection(collName).DeleteMany(contextx.BackgroundWithLogger(s.logger), bson.M{})
		})
	}
}

func (s *suiteRepo) Test_impl_List() {
	type args struct {
		limit  int
		offset int
		mock   func()
	}
	tests := []struct {
		name      string
		args      args
		wantTasks []*todo.Task
		wantErr   bool
	}{
		{
			name:      "list then not found",
			args:      args{limit: 1, offset: 0},
			wantTasks: nil,
			wantErr:   false,
		},
		{
			name: "list then success",
			args: args{limit: 1, offset: 0, mock: func() {
				_, _ = s.client.Database(dbName).Collection(collName).InsertOne(contextx.BackgroundWithLogger(s.logger), testdata.Task1)
			}},
			wantTasks: []*todo.Task{testdata.Task1},
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotTasks, err := s.repo.List(contextx.BackgroundWithLogger(s.logger), tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTasks, tt.wantTasks) {
				t.Errorf("List() gotTasks = %v, want %v", gotTasks, tt.wantTasks)
			}

			_, _ = s.client.Database(dbName).Collection(collName).DeleteMany(contextx.BackgroundWithLogger(s.logger), bson.M{})
		})
	}
}

func (s *suiteRepo) Test_impl_Count() {
	type args struct {
		mock func()
	}
	tests := []struct {
		name      string
		args      args
		wantTotal int
		wantErr   bool
	}{
		{
			name: "count then success",
			args: args{mock: func() {
				_, _ = s.client.Database(dbName).Collection(collName).InsertOne(contextx.BackgroundWithLogger(s.logger), testdata.Task1)
			}},
			wantTotal: 1,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotTotal, err := s.repo.Count(contextx.BackgroundWithLogger(s.logger))
			if (err != nil) != tt.wantErr {
				t.Errorf("Count() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotTotal != tt.wantTotal {
				t.Errorf("Count() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}

			_, _ = s.client.Database(dbName).Collection(collName).DeleteMany(contextx.BackgroundWithLogger(s.logger), bson.M{})
		})
	}
}

func (s *suiteRepo) Test_impl_Create() {
	type args struct {
		newTask *todo.Task
		mock    func()
	}
	tests := []struct {
		name     string
		args     args
		wantTask *todo.Task
		wantErr  bool
	}{
		{
			name:     "create then success",
			args:     args{newTask: testdata.Task1},
			wantTask: testdata.Task1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotTask, err := s.repo.Create(contextx.BackgroundWithLogger(s.logger), tt.args.newTask)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTask, tt.wantTask) {
				t.Errorf("Create() gotTask = %v, want %v", gotTask, tt.wantTask)
			}

			_, _ = s.client.Database(dbName).Collection(collName).DeleteMany(contextx.BackgroundWithLogger(s.logger), bson.M{})
		})
	}
}

func (s *suiteRepo) Test_impl_Update() {
	type args struct {
		updated *todo.Task
		mock    func()
	}
	tests := []struct {
		name     string
		args     args
		wantTask *todo.Task
		wantErr  bool
	}{
		{
			name:     "update then error",
			args:     args{updated: testdata.Task1},
			wantTask: nil,
			wantErr:  true,
		},
		{
			name: "update then success",
			args: args{updated: testdata.Task1, mock: func() {
				_, _ = s.client.Database(dbName).Collection(collName).InsertOne(contextx.BackgroundWithLogger(s.logger), testdata.Task1)
			}},
			wantTask: testdata.Task1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotTask, err := s.repo.Update(contextx.BackgroundWithLogger(s.logger), tt.args.updated)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTask, tt.wantTask) {
				t.Errorf("Update() gotTask = %v, want %v", gotTask, tt.wantTask)
			}

			_, _ = s.client.Database(dbName).Collection(collName).DeleteMany(contextx.BackgroundWithLogger(s.logger), bson.M{})
		})
	}
}

func (s *suiteRepo) Test_impl_Remove() {
	type args struct {
		id   primitive.ObjectID
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "remove then success",
			args: args{id: testdata.Task1.ID, mock: func() {
				_, _ = s.client.Database(dbName).Collection(collName).InsertOne(contextx.BackgroundWithLogger(s.logger), testdata.Task1)
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if err := s.repo.Remove(contextx.BackgroundWithLogger(s.logger), tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Remove() error = %v, wantErr %v", err, tt.wantErr)
			}

			_, _ = s.client.Database(dbName).Collection(collName).DeleteMany(contextx.BackgroundWithLogger(s.logger), bson.M{})
		})
	}
}
