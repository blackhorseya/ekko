//go:generate mockgen -destination=./mock_${GOFILE} -package=tokenx -source=${GOFILE}
package tokenx

import (
	userM "github.com/blackhorseya/ekko/entity/domain/user/model"
	"github.com/golang-jwt/jwt"
	"github.com/google/wire"
)

// TokenClaims declare token claims
type TokenClaims struct {
	jwt.StandardClaims
}

// Tokenizer declare token service
type Tokenizer interface {
	// NewToken serve caller to given a user and generate a token
	NewToken(who *userM.Profile) (token string, err error)

	// ValidateToken serve caller to given a signed token and validate it
	ValidateToken(signedToken string) (claims *TokenClaims, err error)
}

// JwtSet declare tokenx wire set
var JwtSet = wire.NewSet(NewOptions, NewJwtx)
