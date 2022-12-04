package api

import (
	"net/http"

	_ "github.com/blackhorseya/todo-app/api/docs" // import swagger spec
	v1 "github.com/blackhorseya/todo-app/cmd/restful/todo/api/v1"
	"github.com/blackhorseya/todo-app/internal/pkg/errorx"
	"github.com/blackhorseya/todo-app/pkg/contextx"
	tb "github.com/blackhorseya/todo-app/pkg/entity/domain/todo/biz"
	"github.com/blackhorseya/todo-app/pkg/response"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Handle(g *gin.RouterGroup, biz tb.IBiz) {
	i := &impl{biz: biz}

	if gin.Mode() != gin.ReleaseMode {
		g.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	g.GET("readiness", i.Readiness)
	g.GET("liveness", i.Liveness)

	v1.Handle(g.Group("v1"), biz)
}

type impl struct {
	biz tb.IBiz
}

func (i *impl) Readiness(c *gin.Context) {
	ctx, ok := c.MustGet(string(contextx.KeyCtx)).(contextx.Contextx)
	if !ok {
		_ = c.Error(errorx.ErrContextx)
		return
	}

	err := i.biz.Readiness(ctx)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK)
}

func (i *impl) Liveness(c *gin.Context) {
	ctx, ok := c.MustGet(string(contextx.KeyCtx)).(contextx.Contextx)
	if !ok {
		_ = c.Error(errorx.ErrContextx)
		return
	}

	err := i.biz.Liveness(ctx)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK)
}
