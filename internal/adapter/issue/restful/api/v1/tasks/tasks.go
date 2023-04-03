package tasks

import (
	ib "github.com/blackhorseya/ekko/pkg/entity/domain/issue/biz"
	"github.com/gin-gonic/gin"
)

type impl struct {
	biz ib.IBiz
}

func Handle(g *gin.RouterGroup, biz ib.IBiz) {
	i := &impl{biz: biz}

	g.GET(":id", i.GetByID)
	g.GET("", i.List)
	g.POST("", i.Create)
	g.PATCH(":id/status", i.UpdateStatus)
	g.DELETE(":id", i.Delete)
}
