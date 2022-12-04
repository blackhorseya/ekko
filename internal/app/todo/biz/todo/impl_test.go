package todo

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/todo-app/internal/app/todo/biz/todo/repo"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/ticket"
	"github.com/blackhorseya/todo-app/pb"
	"github.com/blackhorseya/todo-app/test/testdata"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type bizSuite struct {
	suite.Suite
	logger *zap.Logger
	mock   *repo.MockITodoRepo
	biz    ITodoBiz
}

func (s *bizSuite) SetupTest() {
	s.logger, _ = zap.NewDevelopment()

	s.mock = new(repo.MockITodoRepo)
	biz, err := CreateIBiz(s.mock)
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
		id   int64
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantTask *ticket.Task
		wantErr  bool
	}{
		{
			name: "get by id then error",
			args: args{id: testdata.Task1.ID, mock: func() {
				s.mock.On("GetByID", mock.Anything, testdata.Task1.ID).Return(nil, errors.New("error")).Once()
			}},
			wantTask: nil,
			wantErr:  true,
		},
		{
			name: "get by id not found then error",
			args: args{id: testdata.Task1.ID, mock: func() {
				s.mock.On("GetByID", mock.Anything, testdata.Task1.ID).Return(nil, nil).Once()
			}},
			wantTask: nil,
			wantErr:  true,
		},
		{
			name: "get by id then success",
			args: args{id: testdata.Task1.ID, mock: func() {
				s.mock.On("GetByID", mock.Anything, testdata.Task1.ID).Return(testdata.Task1, nil).Once()
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

			gotTask, err := s.biz.GetByID(contextx.BackgroundWithLogger(s.logger), tt.args.id)
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
		wantTask *ticket.Task
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
				s.mock.On("Create", mock.Anything, mock.Anything).Return(nil, errors.New("error")).Once()
			}},
			wantTask: nil,
			wantErr:  true,
		},
		{
			name: "create task then success",
			args: args{title: "task 1", mock: func() {
				s.mock.On("Create", mock.Anything, mock.Anything).Return(testdata.Task1, nil).Once()
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

			gotTask, err := s.biz.Create(contextx.BackgroundWithLogger(s.logger), tt.args.title)
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
		id     int64
		status pb.TaskStatus
		mock   func()
	}
	tests := []struct {
		name     string
		args     args
		wantTask *ticket.Task
		wantErr  bool
	}{
		{
			name: "get task by id then error",
			args: args{id: testdata.Task1.ID, mock: func() {
				s.mock.On("GetByID", mock.Anything, testdata.Task1.ID).Return(nil, errors.New("error")).Once()
			}},
			wantTask: nil,
			wantErr:  true,
		},
		{
			name: "get task by id not found then error",
			args: args{id: testdata.Task1.ID, mock: func() {
				s.mock.On("GetByID", mock.Anything, testdata.Task1.ID).Return(nil, nil).Once()
			}},
			wantTask: nil,
			wantErr:  true,
		},
		{
			name: "update status then error",
			args: args{id: testdata.Task1.ID, status: pb.TaskStatus_TASK_STATUS_DONE, mock: func() {
				s.mock.On("GetByID", mock.Anything, testdata.Task1.ID).Return(testdata.Task1, nil).Once()

				updated := testdata.Task1
				s.mock.On("Update", mock.Anything, updated).Return(nil, errors.New("error")).Once()
			}},
			wantTask: nil,
			wantErr:  true,
		},
		{
			name: "update status then success",
			args: args{id: testdata.Task1.ID, status: pb.TaskStatus_TASK_STATUS_DONE, mock: func() {
				s.mock.On("GetByID", mock.Anything, testdata.Task1.ID).Return(testdata.Task1, nil).Once()

				updated := testdata.Task1
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

			gotTask, err := s.biz.UpdateStatus(contextx.BackgroundWithLogger(s.logger), tt.args.id, tt.args.status)
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

func (s *bizSuite) Test_impl_ChangeTitle() {
	type args struct {
		id    int64
		title string
		mock  func()
	}
	tests := []struct {
		name     string
		args     args
		wantTask *ticket.Task
		wantErr  bool
	}{
		{
			name:     "empty title then error",
			args:     args{id: testdata.Task1.ID, title: ""},
			wantTask: nil,
			wantErr:  true,
		},
		{
			name: "get task by id then error",
			args: args{id: testdata.Task1.ID, title: testdata.Task1.Title, mock: func() {
				s.mock.On("GetByID", mock.Anything, testdata.Task1.ID).Return(nil, errors.New("error")).Once()
			}},
			wantTask: nil,
			wantErr:  true,
		},
		{
			name: "get task by id not found then error",
			args: args{id: testdata.Task1.ID, title: testdata.Task1.Title, mock: func() {
				s.mock.On("GetByID", mock.Anything, testdata.Task1.ID).Return(nil, nil).Once()
			}},
			wantTask: nil,
			wantErr:  true,
		},
		{
			name: "change task title by id then error",
			args: args{id: testdata.Task1.ID, title: testdata.Task1.Title, mock: func() {
				s.mock.On("GetByID", mock.Anything, testdata.Task1.ID).Return(testdata.Task1, nil).Once()

				updated := testdata.Task1
				updated.Title = testdata.Task1.Title
				s.mock.On("Update", mock.Anything, updated).Return(nil, errors.New("error")).Once()
			}},
			wantTask: nil,
			wantErr:  true,
		},
		{
			name: "change task title by id then success",
			args: args{id: testdata.Task1.ID, title: testdata.Task1.Title, mock: func() {
				s.mock.On("GetByID", mock.Anything, testdata.Task1.ID).Return(testdata.Task1, nil).Once()

				updated := testdata.Task1
				updated.Title = testdata.Task1.Title
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

			gotTask, err := s.biz.ChangeTitle(contextx.BackgroundWithLogger(s.logger), tt.args.id, tt.args.title)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChangeTitle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTask, tt.wantTask) {
				t.Errorf("ChangeTitle() gotTask = %v, want %v", gotTask, tt.wantTask)
			}

			s.mock.AssertExpectations(t)
		})
	}
}

func (s *bizSuite) Test_impl_Delete() {
	type args struct {
		id   int64
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "delete task by id then error",
			args: args{id: testdata.Task1.ID, mock: func() {
				s.mock.On("Remove", mock.Anything, testdata.Task1.ID).Return(errors.New("error")).Once()
			}},
			wantErr: true,
		},
		{
			name: "delete task by id then success",
			args: args{id: testdata.Task1.ID, mock: func() {
				s.mock.On("Remove", mock.Anything, testdata.Task1.ID).Return(nil).Once()
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if err := s.biz.Delete(contextx.BackgroundWithLogger(s.logger), tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}

			s.mock.AssertExpectations(t)
		})
	}
}

func (s *bizSuite) Test_impl_List() {
	type args struct {
		page int
		size int
		mock func()
	}
	tests := []struct {
		name      string
		args      args
		wantTasks []*ticket.Task
		wantTotal int
		wantErr   bool
	}{
		{
			name:      "invalid page then error",
			args:      args{page: -1, size: 10},
			wantTasks: nil,
			wantTotal: 0,
			wantErr:   true,
		},
		{
			name:      "invalid size then error",
			args:      args{page: 0, size: -10},
			wantTasks: nil,
			wantTotal: 0,
			wantErr:   true,
		},
		{
			name: "list tasks then error",
			args: args{page: 1, size: 10, mock: func() {
				s.mock.On("List", mock.Anything, repo.QueryTodoCondition{Limit: 10, Offset: 0}).Return(nil, errors.New("error")).Once()
			}},
			wantTasks: nil,
			wantTotal: 0,
			wantErr:   true,
		},
		{
			name: "list tasks not found then error",
			args: args{page: 2, size: 5, mock: func() {
				s.mock.On("List", mock.Anything, repo.QueryTodoCondition{Limit: 5, Offset: 5}).Return(nil, nil).Once()
			}},
			wantTasks: nil,
			wantTotal: 0,
			wantErr:   true,
		},
		{
			name: "count all tasks then error",
			args: args{page: 1, size: 0, mock: func() {
				s.mock.On("List", mock.Anything, repo.QueryTodoCondition{Limit: 0, Offset: 0}).Return([]*ticket.Task{testdata.Task1}, nil).Once()

				s.mock.On("Count", mock.Anything, mock.Anything).Return(0, errors.New("error")).Once()
			}},
			wantTasks: nil,
			wantTotal: 0,
			wantErr:   true,
		},
		{
			name: "list and count tasks then success",
			args: args{page: 2, size: 1, mock: func() {
				s.mock.On("List", mock.Anything, repo.QueryTodoCondition{Limit: 1, Offset: 1}).Return([]*ticket.Task{testdata.Task1}, nil).Once()

				s.mock.On("Count", mock.Anything, mock.Anything).Return(10, nil).Once()
			}},
			wantTasks: []*ticket.Task{testdata.Task1},
			wantTotal: 10,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotTasks, gotTotal, err := s.biz.List(contextx.BackgroundWithLogger(s.logger), tt.args.page, tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTasks, tt.wantTasks) {
				t.Errorf("List() gotTasks = %v, want %v", gotTasks, tt.wantTasks)
			}
			if gotTotal != tt.wantTotal {
				t.Errorf("List() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}

			s.mock.AssertExpectations(t)
		})
	}
}
