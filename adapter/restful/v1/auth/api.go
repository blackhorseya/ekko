package auth

import (
	"github.com/blackhorseya/ekko/pkg/authx"
	"github.com/gin-gonic/gin"
)

// Handler is the api handler.
func Handler(g *gin.RouterGroup, authenticator *authx.Authenticator) {
	g.GET("/login", func(c *gin.Context) {
		// todo: 2024/3/22|sean|implement me
		panic("implement me")
	})

	g.GET("/callback", func(c *gin.Context) {
		// todo: 2024/3/22|sean|implement me
		panic("implement me")
	})

	g.GET("/me", func(c *gin.Context) {
		// todo: 2024/3/22|sean|implement me
		panic("implement me")
	})

	g.GET("/logout", func(c *gin.Context) {
		// todo: 2024/3/22|sean|implement me
		panic("implement me")
	})
}
