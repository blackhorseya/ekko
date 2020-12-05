package router

import "github.com/gin-gonic/gin"

// RegisterAPI register api group route
func (r *Router) RegisterAPI(app *gin.Engine) {
	api := app.Group("/api")
	{
		api.GET("/readiness", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "ok",
			})
		})
	}
}
