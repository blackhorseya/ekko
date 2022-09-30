package todo

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/blackhorseya/gocommon/pkg/ginhttp"
	todoBiz "github.com/blackhorseya/todo-app/internal/app/todo/biz/todo"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/er"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/todo"
	"github.com/blackhorseya/todo-app/pb"
	"github.com/blackhorseya/todo-app/test/testdata"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type handlerSuite struct {
	suite.Suite
	r       *gin.Engine
	mock    *todoBiz.MockIBiz
	handler IHandler
}

func (s *handlerSuite) SetupTest() {
	logger, _ := zap.NewDevelopment()

	gin.SetMode(gin.TestMode)
	s.r = gin.New()
	s.r.Use(ginhttp.AddContextx())
	s.r.Use(ginhttp.HandleError())

	s.mock = new(todoBiz.MockIBiz)
	handler, err := CreateIHandler(s.r, logger, s.mock)
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
			name:     "id is invalid then 400",
			args:     args{id: "xxx"},
			wantCode: 400,
		},
		{
			name: "get by id then 500",
			args: args{id: testdata.TaskOID1.Hex(), mock: func() {
				s.mock.On("GetByID", mock.Anything, testdata.TaskOID1).Return(nil, er.ErrGetTask).Once()
			}},
			wantCode: 500,
		},
		{
			name: "get by id then 200",
			args: args{id: testdata.TaskOID1.Hex(), mock: func() {
				s.mock.On("GetByID", mock.Anything, testdata.TaskOID1).Return(testdata.Task1, nil).Once()
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
				s.mock.On("List", mock.Anything, 0, 3).Return([]*todo.Task{testdata.Task1}, 10, nil).Once()
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

func (s *handlerSuite) Test_impl_Create() {
	type args struct {
		title string
		mock  func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name: "create task by title then 500",
			args: args{title: testdata.Task1.Title, mock: func() {
				s.mock.On("Create", mock.Anything, testdata.Task1.Title).Return(nil, er.ErrCreateTask).Once()
			}},
			wantCode: 500,
		},
		{
			name: "create task by title then 201",
			args: args{title: testdata.Task1.Title, mock: func() {
				s.mock.On("Create", mock.Anything, testdata.Task1.Title).Return(testdata.Task1, nil).Once()
			}},
			wantCode: 201,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := "/api/v1/tasks"
			val := url.Values{}
			val.Add("title", tt.args.title)
			req := httptest.NewRequest(http.MethodPost, uri, strings.NewReader(val.Encode()))
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "Create() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}

func (s *handlerSuite) Test_impl_UpdateStatus() {
	type args struct {
		id     string
		status string
		mock   func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name:     "missing id then 400",
			args:     args{id: ""},
			wantCode: 400,
		},
		{
			name:     "invalid id then 400",
			args:     args{id: "xxx"},
			wantCode: 400,
		},
		{
			name: "update status then 500",
			args: args{id: testdata.Task1.ID.Hex(), status: "3", mock: func() {
				s.mock.On("UpdateStatus", mock.Anything, testdata.Task1.ID, pb.TaskStatus_TASK_STATUS_DONE).Return(nil, er.ErrUpdateStatusTask).Once()
			}},
			wantCode: 500,
		},
		{
			name: "update status then 200",
			args: args{id: testdata.Task1.ID.Hex(), status: "3", mock: func() {
				s.mock.On("UpdateStatus", mock.Anything, testdata.Task1.ID, pb.TaskStatus_TASK_STATUS_DONE).Return(testdata.Task1, nil).Once()
			}},
			wantCode: 200,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/tasks/%v/status", tt.args.id)
			val := url.Values{}
			val.Add("status", tt.args.status)
			req := httptest.NewRequest(http.MethodPatch, uri, strings.NewReader(val.Encode()))
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "UpdateStatus() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}

func (s *handlerSuite) Test_impl_ChangeTitle() {
	type args struct {
		id    string
		title string
		mock  func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name:     "missing id then 400",
			args:     args{id: "", title: "title"},
			wantCode: 400,
		},
		{
			name:     "invalid id then 400",
			args:     args{id: "xxx", title: "title"},
			wantCode: 400,
		},
		{
			name: "change title then 500",
			args: args{id: testdata.Task1.ID.Hex(), title: "title", mock: func() {
				s.mock.On("ChangeTitle", mock.Anything, testdata.Task1.ID, "title").Return(nil, er.ErrChangeTitleTask).Once()
			}},
			wantCode: 500,
		},
		{
			name: "change title then 200",
			args: args{id: testdata.Task1.ID.Hex(), title: "title", mock: func() {
				s.mock.On("ChangeTitle", mock.Anything, testdata.Task1.ID, "title").Return(testdata.Task1, nil).Once()
			}},
			wantCode: 200,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/tasks/%v/title", tt.args.id)
			val := url.Values{}
			val.Add("title", tt.args.title)
			req := httptest.NewRequest(http.MethodPatch, uri, strings.NewReader(val.Encode()))
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "ChangeTitle() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.mock.AssertExpectations(t)
		})
	}
}

func (s *handlerSuite) Test_impl_Delete() {
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
			name:     "missing id then 404",
			args:     args{id: ""},
			wantCode: 404,
		},
		{
			name:     "invalid id then 400",
			args:     args{id: "xxx"},
			wantCode: 400,
		},
		{
			name: "delete task by id then 500",
			args: args{id: testdata.Task1.ID.Hex(), mock: func() {
				s.mock.On("Delete", mock.Anything, testdata.Task1.ID).Return(er.ErrDeleteTask).Once()
			}},
			wantCode: 500,
		},
		{
			name: "delete task by id then 200",
			args: args{id: testdata.Task1.ID.Hex(), mock: func() {
				s.mock.On("Delete", mock.Anything, testdata.Task1.ID).Return(nil).Once()
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
			req := httptest.NewRequest(http.MethodDelete, uri, nil)
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "Delete() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}
