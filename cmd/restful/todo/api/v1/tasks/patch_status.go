package tasks

import (
	"net/http"
	"strconv"

	"github.com/blackhorseya/todo-app/internal/pkg/errorx"
	"github.com/blackhorseya/todo-app/pkg/contextx"
	"github.com/blackhorseya/todo-app/pkg/entity/domain/task/model"
	"github.com/blackhorseya/todo-app/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	_formStatus = "status"
)

type patchStatusIDIDRequest struct {
	ID int64 `uri:"id"`
}

// UpdateStatus
// @Summary Update task's status by id
// @Description Update task's status by id
// @Tags Tasks
// @Accept application/x-www-form-urlencoded
// @Produce application/json
// @Param id path int true "ID of task"
// @Param status formData integer true "status"
// @Success 200 {object} response.Response{data=model.Task}
// @Failure 400 {object} er.Error
// @Failure 500 {object} er.Error
// @Router /v1/tasks/{id}/status [patch]
func (i *impl) UpdateStatus(c *gin.Context) {
	ctx, ok := c.MustGet(string(contextx.KeyCtx)).(contextx.Contextx)
	if !ok {
		_ = c.Error(errorx.ErrContextx)
		return
	}

	var req patchStatusIDIDRequest
	err := c.ShouldBindUri(&req)
	if err != nil {
		ctx.Error(errorx.ErrInvalidID.Error(), zap.Error(err))
		_ = c.Error(errorx.ErrInvalidID)
		return
	}

	statusVal, err := strconv.Atoi(c.PostForm(_formStatus))
	if err != nil {
		ctx.Error(errorx.ErrInvalidStatus.Error(), zap.String(_formStatus, c.PostForm(_formStatus)))
		_ = c.Error(errorx.ErrInvalidStatus)
		return
	}

	_, ok = model.TaskStatus_name[int32(statusVal)]
	if !ok {
		ctx.Error(errorx.ErrInvalidStatus.Error(), zap.String(_formStatus, c.PostForm(_formStatus)))
		_ = c.Error(errorx.ErrInvalidStatus)
		return
	}

	ret, err := i.biz.UpdateStatus(ctx, req.ID, model.TaskStatus(statusVal))
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(ret))
}
