package todos

import (
	"net/http"
	"strconv"

	"github.com/blackhorseya/ekko/adapter/platform/wirex"
	"github.com/blackhorseya/ekko/entity/domain/todo/biz"
	_ "github.com/blackhorseya/ekko/entity/domain/todo/model" // import model
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/blackhorseya/ekko/pkg/errorx"
	"github.com/blackhorseya/ekko/pkg/responsex"
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

// GetListQuery defines the list query.
type GetListQuery struct {
	Page int `form:"page" default:"1" minimum:"1"`
	Size int `form:"size" default:"10" minimum:"1" maximum:"100"`
}

// GetList is used to get todo list.
// @Summary Get todo list.
// @Description get todo list.
// @Tags todos
// @Accept json
// @Produce json
// @Param query query GetListQuery false "query string"
// @Success 200 {object} responsex.Response{data=[]model.Todo}
// @Failure 500 {object} responsex.Response
// @Header 200 {int} X-Total-Count "total count"
// @Header 200 {int} X-Page "page"
// @Header 200 {int} X-Page-Size "page size"
// @Router /v1/todos [get]
func (i *impl) GetList(c *gin.Context) {
	ctx, err := contextx.FromGin(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	var query GetListQuery
	err = c.ShouldBindQuery(&query)
	if err != nil {
		responsex.Err(c, errorx.Wrap(http.StatusBadRequest, 400, err))
		return
	}

	ret, total, err := i.injector.Todo.ListTodo(ctx, biz.ListTodoOptions{
		Page: query.Page,
		Size: query.Size,
	})
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.Header("X-Total-Count", strconv.Itoa(total))
	c.Header("X-Page", strconv.Itoa(query.Page))
	c.Header("X-Page-Size", strconv.Itoa(query.Size))
	responsex.OK(c, ret)
}

// PostPayload defines the post payload.
type PostPayload struct {
	Title string `json:"title" binding:"required" example:"example"`
}

// Post is used to create a todo.
// @Summary Create a todo.
// @Description create a todo.
// @Tags todos
// @Accept json
// @Produce json
// @Param payload body PostPayload true "payload"
// @Success 200 {object} responsex.Response{data=model.Todo}
// @Failure 500 {object} responsex.Response
// @Router /v1/todos [post]
func (i *impl) Post(c *gin.Context) {
	ctx, err := contextx.FromGin(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	var payload PostPayload
	err = c.ShouldBindJSON(&payload)
	if err != nil {
		responsex.Err(c, errorx.Wrap(http.StatusBadRequest, 400, err))
		return
	}

	ret, err := i.injector.Todo.CreateTodo(ctx, payload.Title)
	if err != nil {
		_ = c.Error(err)
		return
	}

	responsex.OK(c, ret)
}
