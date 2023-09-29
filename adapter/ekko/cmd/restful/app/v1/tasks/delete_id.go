package tasks

import (
	"net/http"

	"github.com/blackhorseya/ekko/internal/pkg/errorx"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/blackhorseya/ekko/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type deleteByIDRequest struct {
	ID string `uri:"id"`
}

// Delete
// @Summary Delete a issue by id
// @Description Delete a issue by id
// @Tags Tasks
// @Accept application/json
// @Produce application/json
// @Param id path string true "ID of issue"
// @Success 200 {object} response.Response{data=string}
// @Failure 400 {object} er.Error
// @Failure 404 {object} er.Error
// @Failure 500 {object} er.Error
// @Router /v1/tasks/{id} [delete]
func (i *impl) Delete(c *gin.Context) {
	ctx, ok := c.MustGet(string(contextx.KeyCtx)).(contextx.Contextx)
	if !ok {
		_ = c.Error(errorx.ErrContextx)
		return
	}

	var req deleteByIDRequest
	err := c.ShouldBindUri(&req)
	if err != nil {
		ctx.Error(errorx.ErrInvalidID.Error(), zap.Error(err))
		_ = c.Error(errorx.ErrInvalidID)
		return
	}

	err = i.task.DeleteTicket(ctx, req.ID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(req.ID))
}
