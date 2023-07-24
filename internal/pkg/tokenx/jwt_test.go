package tokenx

import (
	"reflect"
	"testing"

	userM "github.com/blackhorseya/ekko/entity/domain/user/model"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

const (
	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2ODE2MTQ2NTIsImlzcyI6ImVra28iLCJzdWIiOiIxIn0.44DJDPFMi9I8-axFtXrZ2Is98mU_VHwEUV1gEw3B_3c"
)

type jwtSuite struct {
	suite.Suite

	logger *zap.Logger
	tokenx Tokenizer
}

func (s *jwtSuite) SetupTest() {
	s.logger, _ = zap.NewDevelopment()

	opts := &Options{
		Issuer:    "ekko",
		Signature: "changeme",
	}
	s.tokenx = CreateTokenizer(opts)
}

func TestJwt(t *testing.T) {
	suite.Run(t, new(jwtSuite))
}

func (s *jwtSuite) Test_jwtx_NewToken() {
	type args struct {
		who *userM.Profile
	}
	tests := []struct {
		name      string
		args      args
		wantToken string
		wantErr   bool
	}{
		{
			name:      "nil user then error",
			args:      args{who: nil},
			wantToken: "",
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			gotToken, err := s.tokenx.NewToken(tt.args.who)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotToken != tt.wantToken {
				t.Errorf("NewToken() gotToken = %v, want %v", gotToken, tt.wantToken)
			}
		})
	}
}

func (s *jwtSuite) Test_jwtx_ValidateToken() {
	type args struct {
		token string
	}
	tests := []struct {
		name       string
		args       args
		wantClaims *TokenClaims
		wantErr    bool
	}{
		{
			name:       "valid token then success",
			args:       args{token: "invalid token"},
			wantClaims: nil,
			wantErr:    true,
		},
		{
			name: "valid token then success",
			args: args{token: token},
			wantClaims: &TokenClaims{StandardClaims: jwt.StandardClaims{
				Audience:  "",
				ExpiresAt: 0,
				Id:        "",
				IssuedAt:  1681614652,
				Issuer:    "ekko",
				NotBefore: 0,
				Subject:   "1",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			gotClaims, err := s.tokenx.ValidateToken(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotClaims, tt.wantClaims) {
				t.Errorf("ValidateToken() gotClaims = %v, want %v", gotClaims, tt.wantClaims)
			}
		})
	}
}
