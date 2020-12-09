package app

import (
	"github.com/blackhorseya/todo-app/internal/app/config"
	"github.com/blackhorseya/todo-app/internal/app/router"
	"github.com/gin-gonic/gin"
)

// NewGinEngine init a engine of Gin
func NewGinEngine(r router.IRouter, config *config.Config) *gin.Engine {
	gin.SetMode(config.RunMode)
	app := gin.Default()

	// register route to Gin engine
	_ = r.Register(app)

	return app
}
