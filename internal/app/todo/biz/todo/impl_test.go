package todo

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/todo-app/internal/app/todo/biz/todo/repo/mocks"
	"github.com/blackhorseya/todo-app/internal/pkg/base/contextx"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/todo"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

var (
	uuid1 = "43fa0832-fd3a-4ba7-a3c7-8b4a36506a83"

	task1 = &todo.Task{
		Id: uuid1,
	}
)

type bizSuite struct {
	suite.Suite
	mock *mocks.IRepo
	biz  IBiz
}

func (s *bizSuite) SetupTest() {
	logger, _ := zap.NewDevelopment()

	s.mock = new(mocks.IRepo)
	biz, err := CreateIBiz(logger, s.mock)
	if err != nil {
		panic(err)
	}

	s.biz = biz
}

func (s *bizSuite) TearDownTest() {
	s.mock.AssertExpectations(s.T())
}

func TestBizSuite(t *testing.T) {
	suite.Run(t, new(bizSuite))
}

func (s *bizSuite) Test_impl_GetByID() {
	type args struct {
		id   string
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantTask *todo.Task
		wantErr  bool
	}{
		{
			name:     "missing id then error",
			args:     args{id: ""},
			wantTask: nil,
			wantErr:  true,
		},
		{
			name:     "id is not a uuid then error",
			args:     args{id: "id"},
			wantTask: nil,
			wantErr:  true,
		},
		{
			name: "get by id then error",
			args: args{id: uuid1, mock: func() {
				s.mock.On("GetByID", mock.Anything, uuid1).Return(nil, errors.New("error")).Once()
			}},
			wantTask: nil,
			wantErr:  true,
		},
		{
			name: "get by id then not found error",
			args: args{id: uuid1, mock: func() {
				s.mock.On("GetByID", mock.Anything, uuid1).Return(nil, nil).Once()
			}},
			wantTask: nil,
			wantErr:  true,
		},
		{
			name: "get by id then success",
			args: args{id: uuid1, mock: func() {
				s.mock.On("GetByID", mock.Anything, uuid1).Return(task1, nil).Once()
			}},
			wantTask: task1,
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

			s.TearDownTest()
		})
	}
}

func (s *bizSuite) Test_impl_List() {
	type args struct {
		start int
		end   int
		mock  func()
	}
	tests := []struct {
		name      string
		args      args
		wantTasks []*todo.Task
		wantTotal int
		wantErr   bool
	}{
		{
			name:      "start < 0 then error",
			args:      args{start: -1, end: 3},
			wantTasks: nil,
			wantTotal: 0,
			wantErr:   true,
		},
		{
			name:      "end < 0 then error",
			args:      args{start: 0, end: -1},
			wantTasks: nil,
			wantTotal: 0,
			wantErr:   true,
		},
		{
			name: "start 0 end 2 list then error",
			args: args{start: 0, end: 2, mock: func() {
				s.mock.On("List", mock.Anything, 3, 0).Return(nil, errors.New("error")).Once()
			}},
			wantTasks: nil,
			wantTotal: 0,
			wantErr:   true,
		},
		{
			name: "start 0 end 2 list then not found error",
			args: args{start: 0, end: 2, mock: func() {
				s.mock.On("List", mock.Anything, 3, 0).Return(nil, nil).Once()
			}},
			wantTasks: nil,
			wantTotal: 0,
			wantErr:   true,
		},
		{
			name: "count then error",
			args: args{start: 0, end: 2, mock: func() {
				s.mock.On("List", mock.Anything, 3, 0).Return([]*todo.Task{task1}, nil).Once()
				s.mock.On("Count", mock.Anything).Return(0, errors.New("error")).Once()
			}},
			wantTasks: nil,
			wantTotal: 0,
			wantErr:   true,
		},
		{
			name: "start 0 end 2 then success",
			args: args{start: 0, end: 2, mock: func() {
				s.mock.On("List", mock.Anything, 3, 0).Return([]*todo.Task{task1}, nil).Once()
				s.mock.On("Count", mock.Anything).Return(10, nil).Once()
			}},
			wantTasks: []*todo.Task{task1},
			wantTotal: 10,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotTasks, gotTotal, err := s.biz.List(contextx.Background(), tt.args.start, tt.args.end)
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

			s.TearDownTest()
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
			name:     "missing title then error",
			args:     args{title: ""},
			wantTask: nil,
			wantErr:  true,
		},
		{
			name: "create then error",
			args: args{title: "title", mock: func() {
				s.mock.On("Create", mock.Anything, "title").Return(nil, errors.New("error")).Once()
			}},
			wantTask: nil,
			wantErr:  true,
		},
		{
			name: "create then success",
			args: args{title: "title", mock: func() {
				s.mock.On("Create", mock.Anything, "title").Return(task1, nil).Once()
			}},
			wantTask: task1,
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

			s.TearDownTest()
		})
	}
}