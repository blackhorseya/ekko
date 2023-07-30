package tasks

import (
	taskB "github.com/blackhorseya/ekko/entity/domain/task/biz"
	"github.com/gin-gonic/gin"
)

type impl struct {
	task taskB.IBiz
}

func Handle(g *gin.RouterGroup, task taskB.IBiz) {
	i := &impl{
		task: task,
	}

	g.GET(":id", i.GetByID)
	g.GET("", i.List)
	g.POST("", i.Create)
	g.PATCH(":id/status", i.UpdateStatus)
	g.DELETE(":id", i.Delete)
}
