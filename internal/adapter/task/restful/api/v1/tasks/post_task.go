package tasks

import (
	"net/http"
	"strings"

	"github.com/blackhorseya/todo-app/internal/pkg/errorx"
	"github.com/blackhorseya/todo-app/pkg/contextx"
	"github.com/blackhorseya/todo-app/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	_formTitle = "title"
)

// Create
// @Summary Create a task
// @Description Create a task
// @Tags Tasks
// @Accept application/x-www-form-urlencoded
// @Produce application/json
// @Param title formData string true "title"
// @Success 201 {object} response.Response{data=model.Task}
// @Failure 400 {object} er.Error
// @Failure 500 {object} er.Error
// @Router /v1/tasks [post]
func (i *impl) Create(c *gin.Context) {
	ctx, ok := c.MustGet(string(contextx.KeyCtx)).(contextx.Contextx)
	if !ok {
		_ = c.Error(errorx.ErrContextx)
		return
	}

	title := strings.ReplaceAll(c.PostForm(_formTitle), " ", "")
	if len(title) == 0 {
		ctx.Error(errorx.ErrInvalidTitle.Error(), zap.String(_formTitle, title))
		_ = c.Error(errorx.ErrInvalidTitle)
		return
	}

	ret, err := i.biz.Create(ctx, title)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(ret))
}