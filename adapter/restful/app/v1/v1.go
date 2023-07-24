package v1

import (
	"github.com/blackhorseya/ekko/adapter/restful/app/v1/tasks"
	issueB "github.com/blackhorseya/ekko/entity/domain/issue/biz"
	"github.com/gin-gonic/gin"
)

func Handle(g *gin.RouterGroup, issue issueB.IBiz) {
	tasks.Handle(g.Group("/tasks"), issue)
}
