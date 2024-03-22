package v1

import (
	"github.com/blackhorseya/ekko/adapter/restful/v1/auth"
	"github.com/blackhorseya/ekko/pkg/authx"
	"github.com/gin-gonic/gin"
)

// Handler is the api handler.
func Handler(g *gin.RouterGroup, authenticator *authx.Authenticator) {
	auth.Handler(g.Group("/auth"), authenticator)
}
