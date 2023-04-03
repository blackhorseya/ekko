package tasks

import (
	"net/http"
	"strconv"

	"github.com/blackhorseya/ekko/internal/pkg/errorx"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/blackhorseya/ekko/pkg/entity/domain/task/model"
	"github.com/blackhorseya/ekko/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	_queryPage   = "page"
	_defaultPage = "1"

	_querySize   = "size"
	_defaultSize = "10"
)

type listResponse struct {
	Total int             `json:"total"`
	List  []*model.Ticket `json:"list"`
}

// List
// @Summary List all tasks
// @Description List all tasks
// @Tags Tasks
// @Accept application/json
// @Produce application/json
// @Param page query integer false "page" default(1)
// @Param size query integer false "size" default(10)
// @Success 200 {object} response.Response{data=listResponse}
// @Failure 400 {object} er.Error
// @Failure 404 {object} er.Error
// @Failure 500 {object} er.Error
// @Router /v1/tasks [get]
func (i *impl) List(c *gin.Context) {
	ctx, ok := c.MustGet(string(contextx.KeyCtx)).(contextx.Contextx)
	if !ok {
		_ = c.Error(errorx.ErrContextx)
		return
	}

	page, err := strconv.Atoi(c.DefaultQuery(_queryPage, _defaultPage))
	if err != nil {
		ctx.Error(errorx.ErrInvalidPage.Error(), zap.Error(err), zap.String(_queryPage, c.Query(_queryPage)))
		_ = c.Error(errorx.ErrInvalidPage)
		return
	}

	size, err := strconv.Atoi(c.DefaultQuery(_querySize, _defaultSize))
	if err != nil {
		ctx.Error(errorx.ErrInvalidSize.Error(), zap.Error(err), zap.String(_querySize, c.Query(_querySize)))
		_ = c.Error(errorx.ErrInvalidSize)
		return
	}

	ret, total, err := i.biz.List(ctx, page, size)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(&listResponse{
		Total: total,
		List:  ret,
	}))
}
