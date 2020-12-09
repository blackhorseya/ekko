package router

import (
	"fmt"

	// import docs for swagger
	_ "github.com/blackhorseya/todo-app/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// RegisterAPI register api group route
func (r *Router) RegisterAPI(app *gin.Engine) {
	if mode := gin.Mode(); mode == gin.DebugMode {
		url := ginSwagger.URL(fmt.Sprintf("http://localhost:%d/docs/doc.json", r.C.HTTP.Port))
		app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}

	api := app.Group("/api")
	{
		api.GET("/readiness", r.HealthAPI.Readiness)
		api.GET("/liveness", r.HealthAPI.Liveness)
	}
}
