package router

import "github.com/gin-gonic/gin"

// RegisterAPI register api group route
func (r *Router) RegisterAPI(app *gin.Engine) {
	api := app.Group("/api")
	{
		api.GET("/readiness", r.HealthAPI.Readiness)
		api.GET("/liveness", r.HealthAPI.Liveness)
	}
}
