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

	g.Group(":id", i.GetByID)
}
