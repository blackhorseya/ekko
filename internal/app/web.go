package app

import (
	"github.com/blackhorseya/todo-app/internal/app/config"
	"github.com/blackhorseya/todo-app/internal/app/middlewares"
	"github.com/blackhorseya/todo-app/internal/app/router"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// NewGinEngine init a engine of Gin
func NewGinEngine(r router.IRouter, config *config.Config) *gin.Engine {
	gin.SetMode(config.RunMode)

	app := gin.Default()

	// logger
	app.Use(middlewares.LoggerMiddleware())

	// register route to Gin engine
	_ = r.Register(app)

	// swagger
	if mode := gin.Mode(); mode == gin.DebugMode {
		app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return app
}
