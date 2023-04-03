package v1

import (
	"github.com/blackhorseya/ekko/internal/adapter/issue/restful/api/v1/tasks"
	ib "github.com/blackhorseya/ekko/pkg/entity/domain/issue/biz"
	"github.com/gin-gonic/gin"
)

func Handle(g *gin.RouterGroup, biz ib.IBiz) {
	tasks.Handle(g.Group("tasks"), biz)
}
