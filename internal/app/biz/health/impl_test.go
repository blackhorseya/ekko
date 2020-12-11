package health

import (
	"errors"
	"testing"

	"github.com/blackhorseya/todo-app/internal/app/biz/health/repository/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type bizTestSuite struct {
	suite.Suite
	healthRepo *mocks.HealthRepo
	healthBiz  Biz
}

func (s *bizTestSuite) SetupTest() {
	s.healthRepo = new(mocks.HealthRepo)
	biz, err := CreateHealthBiz(s.healthRepo)
	if err != nil {
		panic(err)
	}

	s.healthBiz = biz
}

func (s *bizTestSuite) TearDownTest() {
	s.healthRepo.AssertExpectations(s.T())
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
				s.healthRepo.On("Ping", mock.AnythingOfType("time.Duration")).Return(nil).Once()
			},
		},
		{
			name:    "has error then false error",
			wantOk:  false,
			wantErr: "test error",
			mockFunc: func() {
				s.healthRepo.On("Ping", mock.AnythingOfType("time.Duration")).Return(
					errors.New("test error")).Once()
			},
		},
	}
	for _, tt := range tests {
		tt.mockFunc()
		gotOk, err := s.healthBiz.Readiness()
		if err != nil {
			s.EqualErrorf(err, tt.wantErr, "Readiness() error = %v, wantErr %v", err, tt.wantErr)
		}
		s.EqualValuesf(tt.wantOk, gotOk, "Readiness() gotOk = %v, want %v", gotOk, tt.wantOk)
		s.TearDownTest()
	}
}

func Test_impl_Liveness(t *testing.T) {
	// todo: 2020-12-10|10:18|doggy|test it using testify and mock mongo.client
	tests := []struct {
		name    string
		wantOk  bool
		wantErr bool
	}{
		{
			name:    "liveness then true nil",
			wantOk:  true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &impl{}
			gotOk, err := i.Liveness()
			if (err != nil) != tt.wantErr {
				t.Errorf("Liveness() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOk != tt.wantOk {
				t.Errorf("Liveness() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestHealthBiz(t *testing.T) {
	suite.Run(t, new(bizTestSuite))
}
