// +build integration

package repo

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/todo-app/internal/pkg/base/contextx"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/todo"
	"github.com/stretchr/testify/suite"
)

var (
	uuid1 = "43fa0832-fd3a-4ba7-a3c7-8b4a36506a83"

	task1 = &todo.Task{
		Id:    uuid1,
		Title: "title",
	}

	updated1 = &todo.Task{
		Id: uuid1,
		Title: "update",
		Completed: true,
	}
)

type repoSuite struct {
	suite.Suite
	repo IRepo
}

func (s *repoSuite) SetupTest() {
	repo, err := CreateIRepo("../../../../../../configs/app.yaml")
	if err != nil {
		panic(err)
	}

	s.repo = repo
}

func TestRepoSuite(t *testing.T) {
	suite.Run(t, new(repoSuite))
}

func (s *repoSuite) Test_impl_Count() {
	type args struct {
	}
	tests := []struct {
		name      string
		args      args
		wantTotal int
		wantErr   bool
	}{
		{
			name:      "count then success",
			args:      args{},
			wantTotal: 0,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			gotTotal, err := s.repo.Count(contextx.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("Count() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotTotal != tt.wantTotal {
				t.Errorf("Count() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}

func (s *repoSuite) Test_impl_Create() {
	type args struct {
		newTask *todo.Task
	}
	tests := []struct {
		name     string
		args     args
		wantTask *todo.Task
		wantErr  bool
	}{
		{
			name:     "create new task then success",
			args:     args{newTask: task1},
			wantTask: task1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			gotTask, err := s.repo.Create(contextx.Background(), tt.args.newTask)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTask, tt.wantTask) {
				t.Errorf("Create() gotTask = %v, want %v", gotTask, tt.wantTask)
			}
		})
	}
}

func (s *repoSuite) Test_impl_GetByID() {
	type args struct {
		id string
	}
	tests := []struct {
		name     string
		args     args
		wantTask *todo.Task
		wantErr  bool
	}{
		{
			name:     "get by id then success",
			args:     args{id: uuid1},
			wantTask: task1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			gotTask, err := s.repo.GetByID(contextx.Background(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTask, tt.wantTask) {
				t.Errorf("GetByID() gotTask = %v, want %v", gotTask, tt.wantTask)
			}
		})
	}
}

func (s *repoSuite) Test_impl_List() {
	type args struct {
		limit  int
		offset int
	}
	tests := []struct {
		name      string
		args      args
		wantTasks []*todo.Task
		wantErr   bool
	}{
		{
			name:      "list by limit and offset then success",
			args:      args{limit: 3, offset: 0},
			wantTasks: []*todo.Task{task1},
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			gotTasks, err := s.repo.List(contextx.Background(), tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTasks, tt.wantTasks) {
				t.Errorf("List() gotTasks = %v, want %v", gotTasks, tt.wantTasks)
			}
		})
	}
}

func (s *repoSuite) Test_impl_Update() {
	type args struct {
		updated *todo.Task
	}
	tests := []struct {
		name     string
		args     args
		wantTask *todo.Task
		wantErr  bool
	}{
		{
			name:     "update then success",
			args:     args{updated: updated1},
			wantTask: updated1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			gotTask, err := s.repo.Update(contextx.Background(), tt.args.updated)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTask, tt.wantTask) {
				t.Errorf("Update() gotTask = %v, want %v", gotTask, tt.wantTask)
			}
		})
	}
}
