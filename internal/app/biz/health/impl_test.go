package health

import (
	"testing"

	"github.com/blackhorseya/todo-app/internal/app/biz/health/repository/mocks"
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

func Test_impl_Readiness(t *testing.T) {
	// todo: 2020-12-10|10:13|doggy|test it using testify and mock mongo.client
	tests := []struct {
		name    string
		wantOk  bool
		wantErr bool
	}{
		{
			name:    "readiness then true nil",
			wantOk:  true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &impl{}
			gotOk, err := i.Readiness()
			if (err != nil) != tt.wantErr {
				t.Errorf("Readiness() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOk != tt.wantOk {
				t.Errorf("Readiness() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
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
