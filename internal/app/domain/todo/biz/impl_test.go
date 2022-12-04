package biz

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/todo-app/internal/app/domain/todo/biz/repo"
	"github.com/blackhorseya/todo-app/pkg/contextx"
	tb "github.com/blackhorseya/todo-app/pkg/entity/domain/todo/biz"
	"github.com/blackhorseya/todo-app/pkg/entity/domain/todo/model"
	"github.com/blackhorseya/todo-app/pkg/genx"
	"github.com/blackhorseya/todo-app/test/testdata"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type suiteTester struct {
	suite.Suite
	logger    *zap.Logger
	generator *genx.MockGenerator
	repo      *repo.MockIRepo
	biz       tb.IBiz
}

func (s *suiteTester) SetupTest() {
	s.logger, _ = zap.NewDevelopment()
	s.repo = new(repo.MockIRepo)
	s.generator = new(genx.MockGenerator)
	s.biz = CreateBiz(s.repo, s.generator)
}

func (s *suiteTester) assertExpectation(t *testing.T) {
	s.repo.AssertExpectations(t)
}

func TestAll(t *testing.T) {
	suite.Run(t, new(suiteTester))
}

func (s *suiteTester) Test_impl_GetByID() {
	type args struct {
		id   int64
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *model.Task
		wantErr  bool
	}{
		{
			name: "get by id then error",
			args: args{id: testdata.Task2.Id, mock: func() {
				s.repo.On("GetByID", mock.Anything, testdata.Task2.Id).Return(nil, errors.New("error")).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "get by id then ok",
			args: args{id: testdata.Task2.Id, mock: func() {
				s.repo.On("GetByID", mock.Anything, testdata.Task2.Id).Return(testdata.Task2, nil).Once()
			}},
			wantInfo: testdata.Task2,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.biz.GetByID(contextx.BackgroundWithLogger(s.logger), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("GetByID() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}

			s.assertExpectation(t)
		})
	}
}
