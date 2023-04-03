package api

import (
	"net/http"

	_ "github.com/blackhorseya/ekko/api/docs" // import swagger spec
	"github.com/blackhorseya/ekko/internal/adapter/task/restful/api/v1"
	"github.com/blackhorseya/ekko/internal/pkg/errorx"
	"github.com/blackhorseya/ekko/pkg/contextx"
	tb "github.com/blackhorseya/ekko/pkg/entity/domain/task/biz"
	"github.com/blackhorseya/ekko/pkg/response"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Handle(g *gin.RouterGroup, biz tb.IBiz) {
	i := &impl{biz: biz}

	g.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	g.GET("readiness", i.Readiness)
	g.GET("liveness", i.Liveness)

	v1.Handle(g.Group("v1"), biz)
}

type impl struct {
	biz tb.IBiz
}

func (i *impl) Readiness(c *gin.Context) {
	_, ok := c.MustGet(string(contextx.KeyCtx)).(contextx.Contextx)
	if !ok {
		_ = c.Error(errorx.ErrContextx)
		return
	}

	c.JSON(http.StatusOK, response.OK)
}

func (i *impl) Liveness(c *gin.Context) {
	_, ok := c.MustGet(string(contextx.KeyCtx)).(contextx.Contextx)
	if !ok {
		_ = c.Error(errorx.ErrContextx)
		return
	}

	c.JSON(http.StatusOK, response.OK)
}
