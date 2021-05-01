package health

import (
	"net/http"

	health2 "github.com/blackhorseya/todo-app/internal/app/todo/biz/health"
	"github.com/gin-gonic/gin"
)

type impl struct {
	HealthBiz health2.Biz
}

// NewImpl is a constructor of implement health api handler
func NewImpl(healthBiz health2.Biz) IHandler {
	return &impl{HealthBiz: healthBiz}
}

// Readiness to know when an application is ready to start accepting traffic
// @Summary Readiness
// @Description Show application was ready to start accepting traffic
// @Tags Health
// @Accept application/json
// @Produce application/json
// @Success 200 {object} string
// @Failure 500 {object} string
// @Router /readiness [get]
func (h *impl) Readiness(ctx *gin.Context) {

	ok, _ := h.HealthBiz.Readiness()
	if !ok {
	}

	ctx.Status(http.StatusOK)
}

// Liveness to know when to restart an application
// @Summary Liveness
// @Description to know when to restart an application
// @Tags Health
// @Accept application/json
// @Produce application/json
// @Success 200 {object} string
// @Failure 500 {object} string
// @Router /liveness [get]
func (h *impl) Liveness(ctx *gin.Context) {
	ok, _ := h.HealthBiz.Liveness()
	if !ok {
	}

	ctx.Status(http.StatusOK)
}
