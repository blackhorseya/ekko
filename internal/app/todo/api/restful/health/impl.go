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

func (i *impl) Readiness(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	_, err := i.biz.Readiness(ctx)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData("success"))
}

func (i *impl) Liveness(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	_, err := i.biz.Liveness(ctx)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData("success"))
}
