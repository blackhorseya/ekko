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
	newUser := &um.Profile{
		Id:        testdata.Profile1.Id,
		Username:  testdata.Profile1.Username,
		Password:  fmt.Sprintf("%x", sha256.Sum256([]byte(testdata.Profile1.Password))),
		Token:     "",
		CreatedAt: nil,
		UpdatedAt: nil,
	}

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

func (s *suiteBiz) Test_impl_Login() {
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
			name: "get profile by username then error",
			args: args{username: testdata.Profile1.Username, password: testdata.Profile1.Password, mock: func() {
				// mock get profile by username
				s.repo.EXPECT().GetProfileByUsername(gomock.Any(), testdata.Profile1.Username).Return(nil, errors.New("error")).Times(1)
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "if user not exist then error",
			args: args{username: testdata.Profile1.Username, password: testdata.Profile1.Password, mock: func() {
				// mock get profile by username
				s.repo.EXPECT().GetProfileByUsername(gomock.Any(), testdata.Profile1.Username).Return(nil, nil).Times(1)
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "if password not match then error",
			args: args{username: testdata.Profile1.Username, password: testdata.Profile1.Password, mock: func() {
				// mock get profile by username
				exists := &um.Profile{
					Id:        testdata.Profile1.Id,
					Username:  testdata.Profile1.Username,
					Password:  "not match",
					Token:     testdata.Profile1.Token,
					CreatedAt: testdata.Profile1.CreatedAt,
					UpdatedAt: testdata.Profile1.UpdatedAt,
				}
				s.repo.EXPECT().GetProfileByUsername(gomock.Any(), testdata.Profile1.Username).Return(exists, nil).Times(1)
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "generate token then error",
			args: args{username: testdata.Profile1.Username, password: testdata.Profile1.Password, mock: func() {
				// mock get profile by username
				exists := &um.Profile{
					Id:        testdata.Profile1.Id,
					Username:  testdata.Profile1.Username,
					Password:  fmt.Sprintf("%x", sha256.Sum256([]byte(testdata.Profile1.Password))),
					Token:     testdata.Profile1.Token,
					CreatedAt: testdata.Profile1.CreatedAt,
					UpdatedAt: testdata.Profile1.UpdatedAt,
				}
				s.repo.EXPECT().GetProfileByUsername(gomock.Any(), testdata.Profile1.Username).Return(exists, nil).Times(1)

				// mock generate token
				s.tokenizer.EXPECT().NewToken(exists).Return("", errors.New("error")).Times(1)
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "update token then error",
			args: args{username: testdata.Profile1.Username, password: testdata.Profile1.Password, mock: func() {
				// mock get profile by username
				exists := &um.Profile{
					Id:        testdata.Profile1.Id,
					Username:  testdata.Profile1.Username,
					Password:  fmt.Sprintf("%x", sha256.Sum256([]byte(testdata.Profile1.Password))),
					Token:     testdata.Profile1.Token,
					CreatedAt: testdata.Profile1.CreatedAt,
					UpdatedAt: testdata.Profile1.UpdatedAt,
				}
				s.repo.EXPECT().GetProfileByUsername(gomock.Any(), testdata.Profile1.Username).Return(exists, nil).Times(1)

				// mock generate token
				s.tokenizer.EXPECT().NewToken(exists).Return("new token", nil).Times(1)

				// mock update token
				s.repo.EXPECT().UpdateToken(gomock.Any(), exists, "new token").Return(nil, errors.New("error")).Times(1)
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "login success",
			args: args{username: testdata.Profile1.Username, password: testdata.Profile1.Password, mock: func() {
				// mock get profile by username
				exists := &um.Profile{
					Id:        testdata.Profile1.Id,
					Username:  testdata.Profile1.Username,
					Password:  fmt.Sprintf("%x", sha256.Sum256([]byte(testdata.Profile1.Password))),
					Token:     testdata.Profile1.Token,
					CreatedAt: testdata.Profile1.CreatedAt,
					UpdatedAt: testdata.Profile1.UpdatedAt,
				}
				s.repo.EXPECT().GetProfileByUsername(gomock.Any(), testdata.Profile1.Username).Return(exists, nil).Times(1)

				// mock generate token
				s.tokenizer.EXPECT().NewToken(exists).Return("new token", nil).Times(1)

				// mock update token
				s.repo.EXPECT().UpdateToken(gomock.Any(), exists, "new token").Return(testdata.Profile1, nil).Times(1)
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

			gotInfo, err := s.biz.Login(contextx.BackgroundWithLogger(s.logger), tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("Login() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}
		})
	}
}

func (s *suiteBiz) Test_impl_Logout() {
	type args struct {
		who  *um.Profile
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "who is nil then error",
			args:    args{who: nil},
			wantErr: true,
		},
		{
			name: "update token to empty then error",
			args: args{who: testdata.Profile1, mock: func() {
				s.repo.EXPECT().UpdateToken(gomock.Any(), testdata.Profile1, "").Return(nil, errors.New("error")).Times(1)
			}},
			wantErr: true,
		},
		{
			name: "ok",
			args: args{who: testdata.Profile1, mock: func() {
				s.repo.EXPECT().UpdateToken(gomock.Any(), testdata.Profile1, "").Return(testdata.Profile1, nil).Times(1)
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if err := s.biz.Logout(contextx.BackgroundWithLogger(s.logger), tt.args.who); (err != nil) != tt.wantErr {
				t.Errorf("Logout() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
