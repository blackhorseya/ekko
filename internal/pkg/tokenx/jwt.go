package tokenx

import (
	"strconv"
	"time"

	um "github.com/blackhorseya/ekko/pkg/entity/domain/user/model"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// Options declare options for tokenx
type Options struct {
	Issuer    string `json:"issuer" yaml:"issuer"`
	Signature string `json:"signature" yaml:"signature"`
}

// NewOptions create a new options
func NewOptions(v *viper.Viper) (*Options, error) {
	var (
		err error
		o   = new(Options)
	)

	if err = v.UnmarshalKey("jwt", o); err != nil {
		return nil, errors.Wrap(err, "unmarshal jwt option error")
	}

	return o, nil
}

type jwtx struct {
	opts *Options
}

// NewJwtx create a new jwtx
func NewJwtx(opts *Options) Tokenizer {
	return &jwtx{
		opts: opts,
	}
}

func (j *jwtx) NewToken(who *um.Profile) (token string, err error) {
	if who == nil {
		return "", errors.New("user is nil")
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenClaims{
		StandardClaims: jwt.StandardClaims{
			Audience:  "",
			ExpiresAt: 0,
			Id:        "",
			IssuedAt:  time.Now().UTC().Unix(),
			Issuer:    j.opts.Issuer,
			NotBefore: 0,
			Subject:   strconv.FormatInt(who.Id, 10),
		},
	})
	ret, err := claims.SignedString([]byte(j.opts.Signature))
	if err != nil {
		return "", err
	}

	return ret, nil
}

func (j *jwtx) ValidateToken(signedToken string) (claims *TokenClaims, err error) {
	token, err := jwt.ParseWithClaims(signedToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.opts.Signature), nil
	})
	if err != nil {
		return nil, err
	}

	ret, ok := token.Claims.(*TokenClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return ret, nil
}
