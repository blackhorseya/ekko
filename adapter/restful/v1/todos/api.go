package todos

import (
	"github.com/blackhorseya/ekko/entity/domain/workflow/biz"
	"github.com/gin-gonic/gin"
)

type impl struct {
	workflow biz.IWorkflowBiz
}

// Handle is the api handler.
func Handle(g *gin.RouterGroup) {
	instance := &impl{
		workflow: nil,
	}

	g.GET("", instance.GetList)
}
