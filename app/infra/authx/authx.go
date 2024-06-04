package authx

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/blackhorseya/ekko/app/infra/configx"
	"github.com/blackhorseya/ekko/entity/domain/identity/model"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/blackhorseya/ekko/pkg/errorx"
	"github.com/blackhorseya/ekko/pkg/responsex"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CustomClaims is the custom claims.
type CustomClaims struct {
	Email       string   `json:"email,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
}

func (c *CustomClaims) Validate(_ context.Context) error {
	return nil
}

// Authx is a struct that represents the authx.
type Authx struct {
	middleware *jwtmiddleware.JWTMiddleware
}

// NewNil returns a new Authx with nil.
func NewNil() *Authx {
	return &Authx{}
}

// NewAuthx returns a new Authx.
func NewAuthx(app *configx.Application) (*Authx, error) {
	issuerURL, err := url.Parse("https://" + app.Auth0.Domain + "/")
	if err != nil {
		return nil, err
	}

	// create middleware
	jwksProvider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)
	jwtValidator, err := validator.New(
		jwksProvider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		app.Auth0.Audiences,
		validator.WithCustomClaims(func() validator.CustomClaims {
			return &CustomClaims{}
		}),
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		return nil, err
	}
	errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
		contextx.Background().Error("error validating token", zap.Error(err))
	}

	return &Authx{
		middleware: jwtmiddleware.New(jwtValidator.ValidateToken, jwtmiddleware.WithErrorHandler(errorHandler)),
	}, nil
}

// ParseJWT is used to parse the jwt.
func (a *Authx) ParseJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		encounteredError := true
		var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
			encounteredError = false
			c.Request = r

			ctx, err := contextx.FromGin(c)
			if err != nil {
				_ = c.Error(err)
				return
			}

			claims, ok := c.Request.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
			if !ok {
				responsex.Err(c, errorx.Wrap(
					http.StatusUnauthorized,
					401,
					errors.New("claims is not valid"),
				))
				return
			}
			customClaims, ok := claims.CustomClaims.(*CustomClaims)
			if !ok {
				responsex.Err(c, errorx.Wrap(
					http.StatusUnauthorized,
					401,
					errors.New("custom claims is not valid"),
				))
				return
			}

			who := &model.User{
				ID:     claims.RegisteredClaims.Subject,
				Active: true,
				Profile: model.Profile{
					Name:  "",
					Email: customClaims.Email,
				},
			}
			c.Set(contextx.KeyCtx, contextx.WithValue(ctx, contextx.KeyWho, who))

			// continue to the next middleware
			c.Next()
		}

		a.middleware.CheckJWT(handler).ServeHTTP(c.Writer, c.Request)

		if encounteredError {
			responsex.Err(c, errorx.Wrap(
				http.StatusUnauthorized,
				401,
				errors.New("unauthorized"),
			))
			c.Abort()
		}
	}
}
