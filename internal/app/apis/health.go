package apis

import "github.com/gin-gonic/gin"

type Health struct {
}

// Readiness to know when an application is ready to start accepting traffic
func (h *Health) Readiness(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": "ok",
		"message": "application is ready",
	})
}

// Liveness to know when to restart an application
func (h *Health) Liveness(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": "ok",
		"message": "application is live",
	})
}
