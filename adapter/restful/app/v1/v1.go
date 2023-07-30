package v1

import (
	"github.com/blackhorseya/ekko/adapter/restful/app/v1/tasks"
	taskB "github.com/blackhorseya/ekko/entity/domain/task/biz"
	"github.com/gin-gonic/gin"
)

func Handle(g *gin.RouterGroup, taskB taskB.IBiz) {
	tasks.Handle(g.Group("/tasks"), taskB)
}
