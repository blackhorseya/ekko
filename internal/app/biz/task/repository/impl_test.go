// +build integration

package repository

import (
	"testing"

	"github.com/blackhorseya/todo-app/internal/app/entities"
	"github.com/stretchr/testify/suite"
)

type repoTestSuite struct {
	suite.Suite
	taskRepo TaskRepo
}

func (s *repoTestSuite) SetupTest() {
	repo, err := CreateTaskRepo("../../../../../configs/app.yaml")
	if err != nil {
		panic(err)
	}
	s.taskRepo = repo
}

func (s *repoTestSuite) Test_impl_CreateTask() {
	type args struct {
		newTask *entities.Task
	}
	tests := []struct {
		name     string
		args     args
		wantTask *entities.Task
		wantErr  string
	}{
		{
			name: "task then task nil",
			args: args{&entities.Task{
				Title: "test",
			}},
			wantTask: &entities.Task{
				Title: "test",
			},
			wantErr: "",
		},
		{
			name:     "nil then nil error",
			args:     args{nil},
			wantTask: nil,
			wantErr:  "cannot transform type *entities.Task to a BSON Document: WriteNull can only write while positioned on a Element or Value but is positioned on a TopLevel",
		},
	}
	for _, tt := range tests {
		gotTask, err := s.taskRepo.CreateTask(tt.args.newTask)
		if err != nil {
			s.EqualErrorf(err, tt.wantErr, "CreateTask() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		s.EqualValuesf(tt.wantTask.Title, gotTask.Title, "CreateTask() gotTask = %v, want %v", gotTask, tt.wantTask)
	}
}

func TestTaskRepo(t *testing.T) {
	suite.Run(t, new(repoTestSuite))
}
