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

type patchTitleIDRequest struct {
	ID int64 `uri:"id"`
}

// ChangeTitle
// @Summary Change task's title by id
// @Description Change task's title by id
// @Tags Tasks
// @Accept application/x-www-form-urlencoded
// @Produce application/json
// @Param id path int true "ID of task"
// @Param title formData string true "title"
// @Success 200 {object} response.Response{data=model.Task}
// @Failure 400 {object} er.Error
// @Failure 500 {object} er.Error
// @Router /v1/tasks/{id}/title [patch]
func (i *impl) ChangeTitle(c *gin.Context) {
	ctx, ok := c.MustGet(string(contextx.KeyCtx)).(contextx.Contextx)
	if !ok {
		_ = c.Error(errorx.ErrContextx)
		return
	}

	var req patchTitleIDRequest
	err := c.ShouldBindUri(&req)
	if err != nil {
		ctx.Error(errorx.ErrInvalidID.Error(), zap.Error(err))
		_ = c.Error(errorx.ErrInvalidID)
		return
	}

	title := strings.ReplaceAll(c.PostForm(_formTitle), " ", "")
	if len(title) == 0 {
		ctx.Error(errorx.ErrInvalidTitle.Error(), zap.String(_formTitle, title))
		_ = c.Error(errorx.ErrInvalidTitle)
		return
	}

	ret, err := i.biz.ChangeTitle(ctx, req.ID, title)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(ret))
}
