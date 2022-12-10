package tasks

import (
	"net/http"

	"github.com/blackhorseya/todo-app/internal/pkg/errorx"
	"github.com/blackhorseya/todo-app/pkg/contextx"
	"github.com/blackhorseya/todo-app/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type getByIDRequest struct {
	ID int64 `uri:"id"`
}

// GetByID
// @Summary Get a task by id
// @Description Get a task by id
// @Tags Tasks
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID of task"
// @Success 200 {object} response.Response{data=model.Task}
// @Failure 400 {object} er.Error
// @Failure 404 {object} er.Error
// @Failure 500 {object} er.Error
// @Router /v1/tasks/{id} [get]
func (i *impl) GetByID(c *gin.Context) {
	ctx, ok := c.MustGet(string(contextx.KeyCtx)).(contextx.Contextx)
	if !ok {
		_ = c.Error(errorx.ErrContextx)
		return
	}

	var req getByIDRequest
	err := c.ShouldBindUri(&req)
	if err != nil {
		ctx.Error(errorx.ErrBindID.Error(), zap.Error(err))
		_ = c.Error(errorx.ErrBindID)
		return
	}

	ret, err := i.biz.GetByID(ctx, req.ID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(ret))
}
