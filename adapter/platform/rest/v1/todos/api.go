package todos

import (
	"github.com/blackhorseya/ekko/adapter/platform/wirex"
	"github.com/gin-gonic/gin"
)

type impl struct {
	injector *wirex.Injector
}

// Handle is used to handle the todos restful api.
func Handle(g *gin.RouterGroup, injector *wirex.Injector) {
	instance := &impl{
		injector: injector,
	}

	g.GET("", instance.GetList)
	g.POST("", instance.Post)
}

// GetList is used to get todo list.
// @Summary Get todo list.
// @Description get todo list.
// @Tags todos
// @Accept json
// @Produce json
// @Success 200 {object} responsex.Response
// @Failure 500 {object} responsex.Response
// @Router /v1/todos [get]
func (i *impl) GetList(c *gin.Context) {
	// todo: 2024/6/3|sean|implement me
	panic("implement me")
}

type PostPayload struct {
}

// Post is used to create a todo.
// @Summary Create a todo.
// @Description create a todo.
// @Tags todos
// @Accept json
// @Produce json
// @Param payload body PostPayload true "payload"
// @Success 200 {object} responsex.Response
// @Failure 500 {object} responsex.Response
// @Router /v1/todos [post]
func (i *impl) Post(c *gin.Context) {
	// todo: 2024/6/3|sean|implement me
	panic("implement me")
}
