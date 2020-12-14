package health

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/blackhorseya/todo-app/internal/app/biz/health/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type handlerTestSuite struct {
	suite.Suite
	r       *gin.Engine
	biz     *mocks.Biz
	handler IHandler
}

func (s *handlerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)

	s.r = gin.New()
	s.biz = new(mocks.Biz)
	handler, err := CreateHealthHandler(s.biz)
	if err != nil {
		panic(err)
	}

	s.handler = handler
}

func (s *handlerTestSuite) TearDownTest() {
	s.biz.AssertExpectations(s.T())
}

func (s *handlerTestSuite) Test_impl_Readiness() {
	s.r.GET("/api/readiness", s.handler.Readiness)

	type args struct {
	}
	tests := []struct {
		name     string
		args     args
		mockFunc func()
		wantCode int
	}{
		{
			name: "readiness then 200 ok",
			args: args{},
			mockFunc: func() {
				s.biz.On("Readiness").Return(true, nil).Once()
			},
			wantCode: http.StatusOK,
		},
		{
			name: "readiness then 500 fail",
			args: args{},
			mockFunc: func() {
				s.biz.On("Readiness").Return(
					false, errors.New("test error")).Once()
			},
			wantCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		tt.mockFunc()
		uri := "/api/readiness"
		req := httptest.NewRequest(http.MethodGet, uri, nil)
		w := httptest.NewRecorder()

		s.r.ServeHTTP(w, req)

		got := w.Result()
		defer got.Body.Close()

		s.EqualValuesf(tt.wantCode, got.StatusCode, "Readiness() code = [%v], wantCode = [%v]", got.StatusCode, tt.wantCode)

		s.TearDownTest()
	}
}

func TestTaskHandler(t *testing.T) {
	suite.Run(t, new(handlerTestSuite))
}
