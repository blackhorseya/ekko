package tasks

import (
	"net/http"
	"strconv"

	taskM "github.com/blackhorseya/ekko/entity/domain/task/model"
	"github.com/blackhorseya/ekko/internal/pkg/errorx"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/blackhorseya/ekko/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	_formStatus = "status"
)

type patchStatusIDRequest struct {
	ID string `uri:"id"`
}

// UpdateStatus
// @Summary Update issue's status by id
// @Description Update issue's status by id
// @Tags Tasks
// @Accept application/x-www-form-urlencoded
// @Produce application/json
// @Param id path string true "ID of issue"
// @Param status formData integer true "status"
// @Success 200 {object} response.Response{data=model.Ticket}
// @Failure 400 {object} er.Error
// @Failure 500 {object} er.Error
// @Router /v1/tasks/{id}/status [patch]
func (i *impl) UpdateStatus(c *gin.Context) {
	ctx, ok := c.MustGet(string(contextx.KeyCtx)).(contextx.Contextx)
	if !ok {
		_ = c.Error(errorx.ErrContextx)
		return
	}

	var req patchStatusIDRequest
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

	_, ok = taskM.TicketStatus_name[int32(statusVal)]
	if !ok {
		ctx.Error(errorx.ErrInvalidStatus.Error(), zap.String(_formStatus, c.PostForm(_formStatus)))
		_ = c.Error(errorx.ErrInvalidStatus)
		return
	}

	ret, err := i.task.UpdateTicketStatus(ctx, req.ID, taskM.TicketStatus(statusVal))
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(ret))
}
