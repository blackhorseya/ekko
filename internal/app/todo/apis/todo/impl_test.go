package todo

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/blackhorseya/todo-app/internal/app/todo/biz/todo/mocks"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/er"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/todo"
	"github.com/blackhorseya/todo-app/internal/pkg/infra/transports/http/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

var (
	uuid1 = "43fa0832-fd3a-4ba7-a3c7-8b4a36506a83"

	task1 = &todo.Task{
		Id:    uuid1,
		Title: "title",
	}

	updated1 = &todo.Task{
		Id:        uuid1,
		Completed: true,
	}

	updated2 = &todo.Task{
		Id:    uuid1,
		Title: "title",
	}
)

type handlerSuite struct {
	suite.Suite
	r       *gin.Engine
	mock    *mocks.IBiz
	handler IHandler
}

func (s *handlerSuite) SetupTest() {
	logger, _ := zap.NewDevelopment()

	gin.SetMode(gin.TestMode)
	s.r = gin.New()
	s.r.Use(middlewares.ContextMiddleware())
	s.r.Use(middlewares.ErrorMiddleware())

	s.mock = new(mocks.IBiz)
	handler, err := CreateIHandler(logger, s.mock)
	if err != nil {
		panic(err)
	}

	s.handler = handler
}

func (s *handlerSuite) TearDownTest() {
	s.mock.AssertExpectations(s.T())
}

func TestHandlerSuite(t *testing.T) {
	suite.Run(t, new(handlerSuite))
}

func (s *handlerSuite) Test_impl_GetByID() {
	s.r.GET("/api/v1/tasks/:id", s.handler.GetByID)

	type args struct {
		id   string
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name:     "missing id then error",
			args:     args{id: ""},
			wantCode: 404,
		},
		{
			name:     "id is not a uuid then error",
			args:     args{id: "id"},
			wantCode: 400,
		},
		{
			name: "get by id then error",
			args: args{id: uuid1, mock: func() {
				s.mock.On("GetByID", mock.Anything, uuid1).Return(nil, er.ErrGetTask).Once()
			}},
			wantCode: 500,
		},
		{
			name: "get by id then not found error",
			args: args{id: uuid1, mock: func() {
				s.mock.On("GetByID", mock.Anything, uuid1).Return(nil, er.ErrTaskNotExists).Once()
			}},
			wantCode: 404,
		},
		{
			name: "get by id then success",
			args: args{id: uuid1, mock: func() {
				s.mock.On("GetByID", mock.Anything, uuid1).Return(task1, nil).Once()
			}},
			wantCode: 200,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/tasks/%v", tt.args.id)
			req := httptest.NewRequest(http.MethodGet, uri, nil)
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "GetByID() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}

func (s *handlerSuite) Test_impl_List() {
	s.r.GET("/api/v1/tasks", s.handler.List)

	type args struct {
		start string
		end   string
		mock  func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name:     "start not a integer then error",
			args:     args{start: "start", end: "3"},
			wantCode: 400,
		},
		{
			name:     "end not a integer then error",
			args:     args{start: "0", end: "end"},
			wantCode: 400,
		},
		{
			name: "list then error",
			args: args{start: "0", end: "3", mock: func() {
				s.mock.On("List", mock.Anything, 0, 3).Return(nil, 0, er.ErrListTasks).Once()
			}},
			wantCode: 500,
		},
		{
			name: "list then not found error",
			args: args{start: "0", end: "3", mock: func() {
				s.mock.On("List", mock.Anything, 0, 3).Return(nil, 0, er.ErrTaskNotExists).Once()
			}},
			wantCode: 404,
		},
		{
			name: "list then success",
			args: args{start: "0", end: "3", mock: func() {
				s.mock.On("List", mock.Anything, 0, 3).Return([]*todo.Task{task1}, 10, nil).Once()
			}},
			wantCode: 200,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/tasks?start=%v&end=%v", tt.args.start, tt.args.end)
			req := httptest.NewRequest(http.MethodGet, uri, nil)
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "List() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}
