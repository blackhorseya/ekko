package app

import (
	"github.com/blackhorseya/todo-app/internal/app/router"
	"github.com/gin-gonic/gin"
)

// NewGinEngine init a engine of Gin
func NewGinEngine(r router.IRouter) *gin.Engine {
	app := gin.Default()

	// register route to Gin engine
	_ = r.Register(app)

	return app
}
