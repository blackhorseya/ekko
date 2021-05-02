package todo

import (
	"net/http"
	"strconv"

	"github.com/blackhorseya/todo-app/internal/app/todo/biz/todo"
	"github.com/blackhorseya/todo-app/internal/pkg/base/contextx"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/er"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type impl struct {
	logger *zap.Logger
	biz    todo.IBiz
}

// NewImpl serve caller to create an IHandler
func NewImpl(logger *zap.Logger, biz todo.IBiz) IHandler {
	return &impl{
		logger: logger.With(zap.String("type", "TodoHandler")),
		biz:    biz,
	}
}

type reqID struct {
	ID string `uri:"id" binding:"required,uuid"`
}

// GetByID
// @Summary Get a task by id
// @Description Get a task by id
// @Tags Tasks
// @Accept application/json
// @Produce application/json
// @Param id path string true "ID of task"
// @Success 200 {object} response.Response{data=todo.Task}
// @Failure 400 {object} er.APPError
// @Failure 404 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/tasks/{id} [get]
func (i *impl) GetByID(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var req reqID
	if err := c.ShouldBindUri(&req); err != nil {
		i.logger.Error(er.ErrInvalidID.Error(), zap.Error(err))
		c.Error(er.ErrInvalidID)
		return
	}

	ret, err := i.biz.GetByID(ctx, req.ID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(ret))
}

// List
// @Summary List all tasks
// @Description List all tasks
// @Tags Tasks
// @Accept application/json
// @Produce application/json
// @Param start query integer false "start" default(0)
// @Param end query integer false "end" default(10)
// @Success 200 {object} response.Response{data=[]todo.Task}
// @Failure 400 {object} er.APPError
// @Failure 404 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/tasks [get]
func (i *impl) List(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	start, err := strconv.Atoi(c.DefaultQuery("start", "0"))
	if err != nil {
		i.logger.Error(er.ErrInvalidStart.Error(), zap.Error(err), zap.String("start", c.Query("start")))
		c.Error(er.ErrInvalidStart)
		return
	}

	end, err := strconv.Atoi(c.DefaultQuery("end", "10"))
	if err != nil {
		i.logger.Error(er.ErrInvalidEnd.Error(), zap.Error(err), zap.String("end", c.Query("end")))
		c.Error(er.ErrInvalidEnd)
		return
	}

	ret, total, err := i.biz.List(ctx, start, end)
	if err != nil {
		c.Error(err)
		return
	}

	c.Header("x-total", strconv.Itoa(total))
	c.JSON(http.StatusOK, response.OK.WithData(ret))
}

// Create
// @Summary Create a task
// @Description Create a task
// @Tags Tasks
// @Accept application/json
// @Produce application/json
// @Param created body todo.Task true "created task"
// @Success 201 {object} response.Response{data=todo.Task}
// @Failure 400 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/tasks [post]
func (i *impl) Create(c *gin.Context) {
	// todo: 2021-05-02|19:47|doggy|implement me
	panic("implement me")
}

// UpdateStatus
// @Summary Update task's status by id
// @Description Update task's status by id
// @Tags Tasks
// @Accept application/json
// @Produce application/json
// @Param id path string true "ID of task"
// @Param updated body todo.Task true "updated task"
// @Success 200 {object} response.Response{data=todo.Task}
// @Failure 400 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/tasks/{id}/status [patch]
func (i *impl) UpdateStatus(c *gin.Context) {
	// todo: 2021-05-02|19:47|doggy|implement me
	panic("implement me")
}

// ChangeTitle
// @Summary Change task's title by id
// @Description Change task's title by id
// @Tags Tasks
// @Accept application/json
// @Produce application/json
// @Param id path string true "ID of task"
// @Param updated body todo.Task true "updated task"
// @Success 200 {object} response.Response{data=todo.Task}
// @Failure 400 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/tasks/{id}/title [patch]
func (i *impl) ChangeTitle(c *gin.Context) {
	// todo: 2021-05-02|19:47|doggy|implement me
	panic("implement me")
}

// Delete
// @Summary Delete a task by id
// @Description Delete a task by id
// @Tags Tasks
// @Accept application/json
// @Produce application/json
// @Param id path string true "ID of task"
// @Success 204 {object} response.Response
// @Failure 400 {object} er.APPError
// @Failure 404 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/tasks/{id} [delete]
func (i *impl) Delete(c *gin.Context) {
	// todo: 2021-05-02|19:47|doggy|implement me
	panic("implement me")
}
