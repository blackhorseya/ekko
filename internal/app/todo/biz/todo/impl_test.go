package todo

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/todo-app/internal/app/todo/biz/todo/repo/mocks"
	"github.com/blackhorseya/todo-app/internal/pkg/base/contextx"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/todo"
	"github.com/blackhorseya/todo-app/test/testdata"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type bizSuite struct {
	suite.Suite
	mock *mocks.IRepo
	biz  IBiz
}

func (s *bizSuite) SetupTest() {
	logger := zap.NewNop()

	s.mock = new(mocks.IRepo)
	biz, err := CreateIBiz(logger, s.mock)
	if err != nil {
		panic(err)
	}
	s.biz = biz
}

func TestBizSuite(t *testing.T) {
	suite.Run(t, new(bizSuite))
}

func (s *bizSuite) Test_impl_GetByID() {
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
			name:     "nil object id then error",
			args:     args{id: primitive.NilObjectID},
			wantTask: nil,
			wantErr:  true,
		},
		{
			name: "get by id then error",
			args: args{id: testdata.TaskOID1, mock: func() {
				s.mock.On("GetByID", mock.Anything, testdata.TaskOID1).Return(nil, errors.New("error")).Once()
			}},
			wantTask: nil,
			wantErr:  true,
		},
		{
			name: "get by id not found then error",
			args: args{id: testdata.TaskOID1, mock: func() {
				s.mock.On("GetByID", mock.Anything, testdata.TaskOID1).Return(nil, nil).Once()
			}},
			wantTask: nil,
			wantErr:  true,
		},
		{
			name: "get by id then success",
			args: args{id: testdata.TaskOID1, mock: func() {
				s.mock.On("GetByID", mock.Anything, testdata.TaskOID1).Return(testdata.Task1, nil).Once()
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

			gotTask, err := s.biz.GetByID(contextx.Background(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTask, tt.wantTask) {
				t.Errorf("GetByID() gotTask = %v, want %v", gotTask, tt.wantTask)
			}

			s.mock.AssertExpectations(t)
		})
	}
}

func (s *bizSuite) Test_impl_Create() {
	type args struct {
		title string
		mock  func()
	}
	tests := []struct {
		name     string
		args     args
		wantTask *todo.Task
		wantErr  bool
	}{
		{
			name:     "empty title then error",
			args:     args{title: ""},
			wantTask: nil,
			wantErr:  true,
		},
		{
			name: "create task then error",
			args: args{title: "task 1", mock: func() {
				s.mock.On("Create", mock.Anything, testdata.TaskCreate1).Return(nil, errors.New("error")).Once()
			}},
			wantTask: nil,
			wantErr:  true,
		},
		{
			name: "create task then success",
			args: args{title: "task 1", mock: func() {
				s.mock.On("Create", mock.Anything, testdata.TaskCreate1).Return(testdata.Task1, nil).Once()
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

			gotTask, err := s.biz.Create(contextx.Background(), tt.args.title)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTask, tt.wantTask) {
				t.Errorf("Create() gotTask = %v, want %v", gotTask, tt.wantTask)
			}

			s.mock.AssertExpectations(t)
		})
	}
}

func (s *bizSuite) Test_impl_UpdateStatus() {
	type args struct {
		id     primitive.ObjectID
		status bool
		mock   func()
	}
	tests := []struct {
		name     string
		args     args
		wantTask *todo.Task
		wantErr  bool
	}{
		{
			name:     "nil object id then error",
			args:     args{id: primitive.NilObjectID},
			wantTask: nil,
			wantErr:  true,
		},
		{
			name: "get task by id then error",
			args: args{id: testdata.TaskOID1, mock: func() {
				s.mock.On("GetByID", mock.Anything, testdata.TaskOID1).Return(nil, errors.New("error")).Once()
			}},
			wantTask: nil,
			wantErr:  true,
		},
		{
			name: "get task by id not found then error",
			args: args{id: testdata.TaskOID1, mock: func() {
				s.mock.On("GetByID", mock.Anything, testdata.TaskOID1).Return(nil, nil).Once()
			}},
			wantTask: nil,
			wantErr:  true,
		},
		{
			name: "update status then error",
			args: args{id: testdata.TaskOID1, status: true, mock: func() {
				s.mock.On("GetByID", mock.Anything, testdata.TaskOID1).Return(testdata.Task1, nil).Once()

				updated := testdata.Task1
				updated.Completed = true
				s.mock.On("Update", mock.Anything, updated).Return(nil, errors.New("error")).Once()
			}},
			wantTask: nil,
			wantErr:  true,
		},
		{
			name: "update status then success",
			args: args{id: testdata.TaskOID1, status: true, mock: func() {
				s.mock.On("GetByID", mock.Anything, testdata.TaskOID1).Return(testdata.Task1, nil).Once()

				updated := testdata.Task1
				updated.Completed = true
				s.mock.On("Update", mock.Anything, updated).Return(testdata.Task1, nil).Once()
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

			gotTask, err := s.biz.UpdateStatus(contextx.Background(), tt.args.id, tt.args.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTask, tt.wantTask) {
				t.Errorf("UpdateStatus() gotTask = %v, want %v", gotTask, tt.wantTask)
			}

			s.mock.AssertExpectations(t)
		})
	}
}
