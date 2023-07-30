package v1

import (
	"github.com/blackhorseya/ekko/adapter/restful/app/v1/tasks"
	issueB "github.com/blackhorseya/ekko/entity/domain/issue/biz"
	taskB "github.com/blackhorseya/ekko/entity/domain/task/biz"
	"github.com/gin-gonic/gin"
)

func Handle(g *gin.RouterGroup, issue issueB.IBiz, taskB taskB.IBiz) {
	tasks.Handle(g.Group("/tasks"), issue, taskB)
}
