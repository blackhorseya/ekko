package health

import (
	"net/http"

	"github.com/blackhorseya/todo-app/internal/app/biz/health"
	"github.com/gin-gonic/gin"
)

type impl struct {
	HealthBiz health.Biz
}

// NewImpl is a constructor of implement health api handler
func NewImpl(healthBiz health.Biz) IHandler {
	return &impl{HealthBiz: healthBiz}
}

// Readiness to know when an application is ready to start accepting traffic
// @Summary Readiness
// @Description Show application was ready to start accepting traffic
// @Tags Health
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "success"
// @Router /readiness [get]
func (h *impl) Readiness(ctx *gin.Context) {
	var (
		code    = http.StatusOK
		status  = "ok"
		message = "application has been ready"
	)

	ok, _ := h.HealthBiz.Readiness()
	if !ok {
		code = http.StatusInternalServerError
		status = "fail"
		message = "application has failed"
	}

	ctx.JSON(code, gin.H{
		"status":  status,
		"message": message,
	})
}

// Liveness to know when to restart an application
// @Summary Liveness
// @Description to know when to restart an application
// @Tags Health
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "success"
// @Router /liveness [get]
func (h *impl) Liveness(ctx *gin.Context) {
	var (
		code    = http.StatusOK
		status  = "ok"
		message = "application was alive"
	)

	ok, _ := h.HealthBiz.Liveness()
	if !ok {
		code = http.StatusInternalServerError
		status = "fail"
		message = "application has failed"
	}

	ctx.JSON(code, gin.H{
		"status":  status,
		"message": message,
	})
}
