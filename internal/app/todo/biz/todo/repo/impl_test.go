package repo

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/blackhorseya/todo-app/internal/pkg/base/contextx"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/todo"
	"github.com/blackhorseya/todo-app/test/testdata"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type repoSuite struct {
	suite.Suite
	pool     *dockertest.Pool
	resource *dockertest.Resource
	client   *mongo.Client
	repo     IRepo
}

func (s *repoSuite) SetupTest() {
	logger := zap.NewNop()
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

	uri := fmt.Sprintf("mongodb://localhost:%s/", resource.GetPort("27017/tcp"))
	s.client, err = mongo.Connect(contextx.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	repo, err := CreateIRepo(logger, s.client)
	if err != nil {
		panic(err)
	}
	s.repo = repo
}

func (s *repoSuite) TearDownSuite() {
	_ = s.client.Disconnect(contextx.Background())

	_ = s.pool.Purge(s.resource)
}

func TestRepoSuite(t *testing.T) {
	suite.Run(t, new(repoSuite))
}

func (s *repoSuite) Test_impl_GetByID() {
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
			name: "get task by id then success",
			args: args{id: testdata.Task1.ID, mock: func() {
				_, _ = s.client.Database(dbName).Collection(collName).InsertOne(contextx.Background(), testdata.Task1)
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

			gotTask, err := s.repo.GetByID(contextx.Background(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTask, tt.wantTask) {
				t.Errorf("GetByID() gotTask = %v, want %v", gotTask, tt.wantTask)
			}

			_ = s.client.Database(dbName).Collection(collName).Drop(contextx.Background())
		})
	}
}

func (s *repoSuite) Test_impl_Update() {
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
			name: "update a task then success",
			args: args{updated: testdata.TaskUpdate1, mock: func() {
				_, _ = s.client.Database(dbName).Collection(collName).InsertOne(contextx.Background(), testdata.Task1)
			}},
			wantTask: testdata.TaskUpdate1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotTask, err := s.repo.Update(contextx.Background(), tt.args.updated)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTask, tt.wantTask) {
				t.Errorf("Update() gotTask = %v, want %v", gotTask, tt.wantTask)
			}

			_ = s.client.Database(dbName).Collection(collName).Drop(contextx.Background())
		})
	}
}

func (s *repoSuite) Test_impl_Remove() {
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
			name: "remove task by id then success",
			args: args{id: testdata.Task1.ID, mock: func() {
				_, _ = s.client.Database(dbName).Collection(collName).InsertOne(contextx.Background(), testdata.Task1)
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if err := s.repo.Remove(contextx.Background(), tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Remove() error = %v, wantErr %v", err, tt.wantErr)
			}

			_ = s.client.Database(dbName).Collection(collName).Drop(contextx.Background())
		})
	}
}

func (s *repoSuite) Test_impl_Create() {
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
			name:     "create a task then success",
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

			gotTask, err := s.repo.Create(contextx.Background(), tt.args.newTask)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTask, tt.wantTask) {
				t.Errorf("Create() gotTask = %v, want %v", gotTask, tt.wantTask)
			}

			_ = s.client.Database(dbName).Collection(collName).Drop(contextx.Background())
		})
	}
}
