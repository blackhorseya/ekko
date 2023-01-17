package tasks

import (
	tb "github.com/blackhorseya/todo-app/pkg/entity/domain/task/biz"
	"github.com/gin-gonic/gin"
)

type impl struct {
	biz tb.IBiz
}

func Handle(g *gin.RouterGroup, biz tb.IBiz) {
	i := &impl{biz: biz}

	g.GET(":id", i.GetByID)
	g.GET("", i.List)
	g.POST("", i.Create)
	g.PATCH(":id/status", i.UpdateStatus)
	g.PATCH(":id/title", i.ChangeTitle)
	g.DELETE(":id", i.Delete)
}
