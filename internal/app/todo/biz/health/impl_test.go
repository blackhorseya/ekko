package health

import (
	"errors"
	"testing"
	"time"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/todo-app/internal/app/todo/biz/health/repo"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type bizTestSuite struct {
	suite.Suite
	logger *zap.Logger
	mock   *repo.MockIRepo
	biz    IBiz
}

func (s *bizTestSuite) SetupTest() {
	s.logger, _ = zap.NewDevelopment()

	s.mock = new(repo.MockIRepo)
	biz, err := CreateIBiz(s.mock)
	if err != nil {
		panic(err)
	}

	s.biz = biz
}

func (s *bizTestSuite) TearDownTest() {
	s.mock.AssertExpectations(s.T())
}

func TestSuiteTest(t *testing.T) {
	suite.Run(t, new(bizTestSuite))
}

func (s *bizTestSuite) Test_impl_Readiness() {
	tests := []struct {
		name     string
		wantOk   bool
		wantErr  string
		mockFunc func()
	}{
		{
			name:    "no error then true nil",
			wantOk:  true,
			wantErr: "",
			mockFunc: func() {
				s.mock.On("Ping", mock.Anything, 2*time.Second).Return(nil).Once()
			},
		},
		{
			name:    "has error then false error",
			wantOk:  false,
			wantErr: "db ping is failure",
			mockFunc: func() {
				s.mock.On("Ping", mock.Anything, 2*time.Second).Return(errors.New("db ping is failure")).Once()
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			tt.mockFunc()
			gotOk, err := s.biz.Readiness(contextx.BackgroundWithLogger(s.logger))
			if err != nil {
				s.EqualErrorf(err, tt.wantErr, "Readiness() error = %v, wantErr %v", err, tt.wantErr)
			}
			s.EqualValuesf(tt.wantOk, gotOk, "Readiness() gotOk = %v, want %v", gotOk, tt.wantOk)
			s.TearDownTest()
		})
	}
}

func (s *bizTestSuite) Test_impl_Liveness() {
	tests := []struct {
		name     string
		wantOk   bool
		wantErr  string
		mockFunc func()
	}{
		{
			name:    "no error then true nil",
			wantOk:  true,
			wantErr: "",
			mockFunc: func() {
				s.mock.On("Ping", mock.Anything, 5*time.Second).Return(nil).Once()
			},
		},
		{
			name:    "has error then false error",
			wantOk:  false,
			wantErr: "db ping is failure",
			mockFunc: func() {
				s.mock.On("Ping", mock.Anything, 5*time.Second).Return(errors.New("db ping is failure")).Once()
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			tt.mockFunc()
			gotOk, err := s.biz.Liveness(contextx.BackgroundWithLogger(s.logger))
			if err != nil {
				s.EqualErrorf(err, tt.wantErr, "Liveness() error = %v, wantErr %v", err, tt.wantErr)
			}
			s.EqualValuesf(tt.wantOk, gotOk, "Liveness() gotOk = %v, want %v", gotOk, tt.wantOk)
			s.TearDownTest()
		})
	}
}
