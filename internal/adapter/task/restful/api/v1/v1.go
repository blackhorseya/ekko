package v1

import (
	"github.com/blackhorseya/ekko/internal/adapter/task/restful/api/v1/tasks"
	tb "github.com/blackhorseya/ekko/pkg/entity/domain/task/biz"
	"github.com/gin-gonic/gin"
)

func Handle(g *gin.RouterGroup, biz tb.IBiz) {
	tasks.Handle(g.Group("tasks"), biz)
}
