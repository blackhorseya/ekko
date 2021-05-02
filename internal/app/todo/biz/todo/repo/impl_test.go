package repo

import (
	"testing"

	"github.com/blackhorseya/todo-app/internal/pkg/base/contextx"
	"github.com/stretchr/testify/suite"
)

type repoSuite struct {
	suite.Suite
	repo IRepo
}

func (s *repoSuite) SetupTest() {
	repo, err := CreateIRepo("../../../../../../configs/app.yaml")
	if err != nil {
		panic(err)
	}

	s.repo = repo
}

func TestRepoSuite(t *testing.T) {
	suite.Run(t, new(repoSuite))
}

func (s *repoSuite) Test_impl_Count() {
	type args struct {
	}
	tests := []struct {
		name      string
		args      args
		wantTotal int
		wantErr   bool
	}{
		{
			name:      "count then success",
			args:      args{},
			wantTotal: 0,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			gotTotal, err := s.repo.Count(contextx.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("Count() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotTotal != tt.wantTotal {
				t.Errorf("Count() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}
