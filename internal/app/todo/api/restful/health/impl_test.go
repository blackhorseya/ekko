package health

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/blackhorseya/gocommon/pkg/ginhttp"
	"github.com/blackhorseya/todo-app/internal/app/todo/biz/health"
	"github.com/blackhorseya/todo-app/internal/pkg/errorx"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type handlerSuite struct {
	suite.Suite
	r       *gin.Engine
	mock    *health.MockIHealthBiz
	handler IHandler
}

func (s *handlerSuite) SetupTest() {
	logger, _ := zap.NewDevelopment()

	gin.SetMode(gin.TestMode)
	s.r = gin.New()
	s.r.Use(ginhttp.AddContextxWithLogger(logger))
	s.r.Use(ginhttp.HandleError())

	s.mock = new(health.MockIHealthBiz)
	if handler, err := CreateIHandler(s.r, s.mock); err != nil {
		panic(err)
	} else {
		s.handler = handler
	}
}

func (s *handlerSuite) TearDownTest() {
	s.mock.AssertExpectations(s.T())
}

func TestHandlerSuite(t *testing.T) {
	suite.Run(t, new(handlerSuite))
}

func (s *handlerSuite) Test_impl_Readiness() {
	type args struct {
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name: "readiness then error",
			args: args{mock: func() {
				s.mock.On("Readiness", mock.Anything).Return(false, errorx.ErrPing).Once()
			}},
			wantCode: 500,
		},
		{
			name: "readiness then success",
			args: args{mock: func() {
				s.mock.On("Readiness", mock.Anything).Return(true, nil).Once()
			}},
			wantCode: 200,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := "/api/readiness"
			req := httptest.NewRequest(http.MethodGet, uri, nil)
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "Readiness() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}

func (s *handlerSuite) Test_impl_Liveness() {
	type args struct {
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name: "liveness then error",
			args: args{mock: func() {
				s.mock.On("Liveness", mock.Anything).Return(false, errorx.ErrPing).Once()
			}},
			wantCode: 500,
		},
		{
			name: "liveness then success",
			args: args{mock: func() {
				s.mock.On("Liveness", mock.Anything).Return(true, nil).Once()
			}},
			wantCode: 200,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := "/api/liveness"
			req := httptest.NewRequest(http.MethodGet, uri, nil)
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "Liveness() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}
