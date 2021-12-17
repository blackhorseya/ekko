package todo

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/todo-app/internal/app/todo/biz/todo/repo/mocks"
	"github.com/blackhorseya/todo-app/internal/pkg/base/contextx"
	"github.com/blackhorseya/todo-app/pb"
	"github.com/bwmarrin/snowflake"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

var (
	uuid1 = int64(1)

	task1 = &pb.Task{
		Id:    uuid1,
		Title: "title",
	}

	updated1 = &pb.Task{
		Id:        uuid1,
		Completed: true,
	}

	updated2 = &pb.Task{
		Id:    uuid1,
		Title: "title",
	}
)

type bizSuite struct {
	suite.Suite
	mock *mocks.IRepo
	biz  IBiz
}

func (s *bizSuite) SetupTest() {
	logger, _ := zap.NewDevelopment()
	node, _ := snowflake.NewNode(1)

	s.mock = new(mocks.IRepo)
	biz, err := CreateIBiz(logger, s.mock, node)
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
		id   int64
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantTask *pb.Task
		wantErr  bool
	}{
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
		wantTasks []*pb.Task
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
				s.mock.On("List", mock.Anything, 3, 0).Return([]*pb.Task{task1}, nil).Once()
				s.mock.On("Count", mock.Anything).Return(0, errors.New("error")).Once()
			}},
			wantTasks: nil,
			wantTotal: 0,
			wantErr:   true,
		},
		{
			name: "start 0 end 2 then success",
			args: args{start: 0, end: 2, mock: func() {
				s.mock.On("List", mock.Anything, 3, 0).Return([]*pb.Task{task1}, nil).Once()
				s.mock.On("Count", mock.Anything).Return(10, nil).Once()
			}},
			wantTasks: []*pb.Task{task1},
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
		wantTask *pb.Task
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
				s.mock.On("Create", mock.Anything, mock.Anything).Return(nil, errors.New("error")).Once()
			}},
			wantTask: nil,
			wantErr:  true,
		},
		{
			name: "create then success",
			args: args{title: "title", mock: func() {
				s.mock.On("Create", mock.Anything, mock.Anything).Return(task1, nil).Once()
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
			name: "uuid remove then error",
			args: args{id: uuid1, mock: func() {
				s.mock.On("Remove", mock.Anything, uuid1).Return(errors.New("error")).Once()
			}},
			wantErr: true,
		},
		{
			name: "uuid remove then success",
			args: args{id: uuid1, mock: func() {
				s.mock.On("Remove", mock.Anything, uuid1).Return(nil).Once()
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if err := s.biz.Delete(contextx.Background(), tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}

			s.TearDownTest()
		})
	}
}

func (s *bizSuite) Test_impl_UpdateStatus() {
	type args struct {
		id     int64
		status bool
		mock   func()
	}
	tests := []struct {
		name     string
		args     args
		wantTask *pb.Task
		wantErr  bool
	}{
		{
			name: "get by id then error",
			args: args{id: uuid1, status: true, mock: func() {
				s.mock.On("GetByID", mock.Anything, uuid1).Return(nil, errors.New("error")).Once()
			}},
			wantTask: nil,
			wantErr:  true,
		},
		{
			name: "get by id then not found error",
			args: args{id: uuid1, status: true, mock: func() {
				s.mock.On("GetByID", mock.Anything, uuid1).Return(nil, nil).Once()
			}},
			wantTask: nil,
			wantErr:  true,
		},
		{
			name: "update status then error",
			args: args{id: uuid1, status: true, mock: func() {
				s.mock.On("GetByID", mock.Anything, uuid1).Return(task1, nil).Once()
				s.mock.On("Update", mock.Anything, task1).Return(nil, errors.New("error")).Once()
			}},
			wantTask: nil,
			wantErr:  true,
		},
		{
			name: "update status then success",
			args: args{id: uuid1, status: true, mock: func() {
				s.mock.On("GetByID", mock.Anything, uuid1).Return(task1, nil).Once()
				s.mock.On("Update", mock.Anything, task1).Return(updated1, nil).Once()
			}},
			wantTask: updated1,
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

			s.TearDownTest()
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
		wantTask *pb.Task
		wantErr  bool
	}{
		{
			name:     "missing title then error",
			args:     args{id: uuid1, title: ""},
			wantTask: nil,
			wantErr:  true,
		},
		{
			name: "get by id then error",
			args: args{id: uuid1, title: "title", mock: func() {
				s.mock.On("GetByID", mock.Anything, uuid1).Return(nil, errors.New("error")).Once()
			}},
			wantTask: nil,
			wantErr:  true,
		},
		{
			name: "get by id then not found error",
			args: args{id: uuid1, title: "title", mock: func() {
				s.mock.On("GetByID", mock.Anything, uuid1).Return(nil, nil).Once()
			}},
			wantTask: nil,
			wantErr:  true,
		},
		{
			name: "change title then error",
			args: args{id: uuid1, title: "title", mock: func() {
				s.mock.On("GetByID", mock.Anything, uuid1).Return(task1, nil).Once()
				s.mock.On("Update", mock.Anything, updated2).Return(nil, errors.New("error")).Once()
			}},
			wantTask: nil,
			wantErr:  true,
		},
		{
			name: "change title then success",
			args: args{id: uuid1, title: "title", mock: func() {
				s.mock.On("GetByID", mock.Anything, uuid1).Return(task1, nil).Once()
				s.mock.On("Update", mock.Anything, updated2).Return(updated2, nil).Once()
			}},
			wantTask: updated2,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotTask, err := s.biz.ChangeTitle(contextx.Background(), tt.args.id, tt.args.title)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChangeTitle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTask, tt.wantTask) {
				t.Errorf("ChangeTitle() gotTask = %v, want %v", gotTask, tt.wantTask)
			}
		})
	}
}
