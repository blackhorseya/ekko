package app

import "github.com/gin-gonic/gin"

// NewGinEngine init a engine of Gin
func NewGinEngine() *gin.Engine {
	engine := gin.New()

	api := engine.Group("/api")
	{
		api.GET("/readiness", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "ok",
			})
		})
	}

	return engine
}
