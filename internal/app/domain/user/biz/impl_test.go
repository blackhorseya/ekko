package biz

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/blackhorseya/ekko/internal/app/domain/user/biz/repo"
	"github.com/blackhorseya/ekko/internal/pkg/tokenx"
	"github.com/blackhorseya/ekko/pkg/contextx"
	ub "github.com/blackhorseya/ekko/pkg/entity/domain/user/biz"
	um "github.com/blackhorseya/ekko/pkg/entity/domain/user/model"
	"github.com/blackhorseya/ekko/pkg/genx"
	"github.com/blackhorseya/ekko/test/testdata"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type suiteBiz struct {
	suite.Suite

	logger *zap.Logger
	ctrl   *gomock.Controller

	repo      *repo.MockIRepo
	node      *genx.MockGenerator
	tokenizer *tokenx.MockTokenizer
	biz       ub.IBiz
}

func (s *suiteBiz) SetupTest() {
	s.logger, _ = zap.NewDevelopment()
	s.ctrl = gomock.NewController(s.T())

	s.repo = repo.NewMockIRepo(s.ctrl)
	s.node = genx.NewMockGenerator(s.ctrl)
	s.tokenizer = tokenx.NewMockTokenizer(s.ctrl)

	s.biz = CreateBiz(s.repo, s.node, s.tokenizer)
}

func (s *suiteBiz) TearDownTest() {
	s.ctrl.Finish()
}

func TestAll(t *testing.T) {
	suite.Run(t, new(suiteBiz))
}

func (s *suiteBiz) Test_impl_Signup() {
	type args struct {
		username string
		password string
		mock     func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *um.Profile
		wantErr  bool
	}{
		{
			name:     "username empty then error",
			args:     args{username: "", password: testdata.Profile1.Password},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name:     "password empty then error",
			args:     args{username: testdata.Profile1.Username, password: ""},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "register failed then error",
			args: args{username: testdata.Profile1.Username, password: testdata.Profile1.Password, mock: func() {
				// mock generate id
				s.node.EXPECT().Int64().Return(testdata.Profile1.Id).Times(1)

				// mock register
				newUser := &um.Profile{
					Id:        testdata.Profile1.Id,
					Username:  testdata.Profile1.Username,
					Password:  fmt.Sprintf("%x", sha256.Sum256([]byte(testdata.Profile1.Password))),
					Token:     "",
					CreatedAt: nil,
					UpdatedAt: nil,
				}
				s.repo.EXPECT().Register(gomock.Any(), newUser).Return(nil, errors.New("error")).Times(1)
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "register success then return info",
			args: args{username: testdata.Profile1.Username, password: testdata.Profile1.Password, mock: func() {
				// mock generate id
				s.node.EXPECT().Int64().Return(testdata.Profile1.Id).Times(1)

				// mock register
				newUser := &um.Profile{
					Id:        testdata.Profile1.Id,
					Username:  testdata.Profile1.Username,
					Password:  fmt.Sprintf("%x", sha256.Sum256([]byte(testdata.Profile1.Password))),
					Token:     "",
					CreatedAt: nil,
					UpdatedAt: nil,
				}
				s.repo.EXPECT().Register(gomock.Any(), newUser).Return(testdata.Profile1, nil).Times(1)
			}},
			wantInfo: testdata.Profile1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.biz.Signup(contextx.BackgroundWithLogger(s.logger), tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Signup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("Signup() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}
		})
	}
}
