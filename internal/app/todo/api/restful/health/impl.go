package health

import (
	"net/http"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/gocommon/pkg/response"
	"github.com/blackhorseya/todo-app/internal/app/todo/biz/health"
	"github.com/gin-gonic/gin"

	// import entity
	_ "github.com/blackhorseya/gocommon/pkg/er"
)

type impl struct {
	biz health.IHealthBiz
}

// NewImpl is a constructor of implement health api handler
func NewImpl(e *gin.Engine, biz health.IHealthBiz) IHandler {
	ret := &impl{biz: biz}

	e.GET("/api/readiness", ret.Readiness)
	e.GET("/api/liveness", ret.Liveness)

	return ret
}

// Readiness to know when an application is ready to start accepting traffic
// @Summary Readiness
// @Description Show application was ready to start accepting traffic
// @Tags Health
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response
// @Failure 500 {object} er.APPError
// @Router /readiness [get]
func (i *impl) Readiness(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	_, err := i.biz.Readiness(ctx)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData("success"))
}

// Liveness to know when to restart an application
// @Summary Liveness
// @Description to know when to restart an application
// @Tags Health
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response
// @Failure 500 {object} er.APPError
// @Router /liveness [get]
func (i *impl) Liveness(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	_, err := i.biz.Liveness(ctx)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData("success"))
}
