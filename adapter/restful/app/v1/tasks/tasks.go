package tasks

import (
	issueB "github.com/blackhorseya/ekko/entity/domain/issue/biz"
	"github.com/gin-gonic/gin"
)

type impl struct {
	biz issueB.IBiz
}

func Handle(g *gin.RouterGroup, biz issueB.IBiz) {
	i := &impl{biz: biz}

	g.GET(":id", i.GetByID)
	g.GET("", i.List)
	g.POST("", i.Create)
	g.PATCH(":id/status", i.UpdateStatus)
	g.DELETE(":id", i.Delete)
}
