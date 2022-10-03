package todo

import (
	"net/http"
	"strconv"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/gocommon/pkg/response"
	"github.com/blackhorseya/todo-app/internal/app/todo/biz/todo"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/er"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/ticket"
	"github.com/blackhorseya/todo-app/pb"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	// import entity
	_ "github.com/blackhorseya/gocommon/pkg/er"
)

type impl struct {
	biz todo.ITodoBiz
}

// NewImpl serve caller to create an IHandler
func NewImpl(e *gin.Engine, biz todo.ITodoBiz) IHandler {
	ret := &impl{biz: biz}

	api := e.Group("api")
	{
		v1 := api.Group("v1")
		{
			tasks := v1.Group("tasks")
			{
				tasks.GET("", ret.List)
				tasks.GET(":id", ret.GetByID)
				tasks.POST("", ret.Create)
				tasks.PATCH(":id/status", ret.UpdateStatus)
				tasks.PATCH(":id/title", ret.ChangeTitle)
				tasks.DELETE(":id", ret.Delete)
			}
		}
	}

	return ret
}

// GetByID
// @Summary Get a task by id
// @Description Get a task by id
// @Tags Tasks
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID of task"
// @Success 200 {object} response.Response{data=pb.Task}
// @Failure 400 {object} er.APPError
// @Failure 404 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/tasks/{id} [get]
func (i *impl) GetByID(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var req reqID
	err := c.ShouldBindUri(&req)
	if err != nil {
		ctx.Error(er.ErrBindID.Error(), zap.Error(err))
		_ = c.Error(er.ErrBindID)
		return
	}

	ret, err := i.biz.GetByID(ctx, req.ID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(ticket.NewTaskResponse(ret)))
}

// List
// @Summary List all tasks
// @Description List all tasks
// @Tags Tasks
// @Accept application/json
// @Produce application/json
// @Param page query integer false "page" default(1)
// @Param size query integer false "size" default(10)
// @Success 200 {object} response.Response{data=[]pb.Task}
// @Failure 400 {object} er.APPError
// @Failure 404 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/tasks [get]
func (i *impl) List(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		ctx.Error(er.ErrInvalidPage.Error(), zap.Error(err), zap.String("page", c.Query("page")))
		_ = c.Error(er.ErrInvalidPage)
		return
	}

	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	if err != nil {
		ctx.Error(er.ErrInvalidSize.Error(), zap.Error(err), zap.String("size", c.Query("size")))
		_ = c.Error(er.ErrInvalidSize)
		return
	}

	tasks, total, err := i.biz.List(ctx, page, size)
	if err != nil {
		_ = c.Error(err)
		return
	}

	ret := make([]*pb.Task, len(tasks))
	for idx, task := range tasks {
		ret[idx] = ticket.NewTaskResponse(task)
	}

	c.Header("X-Total-Count", strconv.Itoa(total))
	c.JSON(http.StatusOK, response.OK.WithData(ret))
}

// Create
// @Summary Create a task
// @Description Create a task
// @Tags Tasks
// @Accept application/x-www-form-urlencoded
// @Produce application/json
// @Param title formData string true "title"
// @Success 201 {object} response.Response{data=pb.Task}
// @Failure 400 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/tasks [post]
func (i *impl) Create(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	title := c.PostForm("title")

	ret, err := i.biz.Create(ctx, title)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, response.OK.WithData(ticket.NewTaskResponse(ret)))
}

// UpdateStatus
// @Summary Update task's status by id
// @Description Update task's status by id
// @Tags Tasks
// @Accept application/x-www-form-urlencoded
// @Produce application/json
// @Param id path int true "ID of task"
// @Param status formData integer true "status"
// @Success 200 {object} response.Response{data=pb.Task}
// @Failure 400 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/tasks/{id}/status [patch]
func (i *impl) UpdateStatus(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var req reqID
	err := c.ShouldBindUri(&req)
	if err != nil {
		ctx.Error(er.ErrBindID.Error(), zap.Error(err))
		_ = c.Error(er.ErrBindID)
		return
	}

	statusStr := c.PostForm("status")
	statusVal, err := strconv.Atoi(statusStr)
	if err != nil {
		_ = c.Error(er.ErrInvalidStatus)
		return
	}

	ret, err := i.biz.UpdateStatus(ctx, req.ID, pb.TaskStatus(int32(statusVal)))
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(ticket.NewTaskResponse(ret)))
}

// ChangeTitle
// @Summary Change task's title by id
// @Description Change task's title by id
// @Tags Tasks
// @Accept application/x-www-form-urlencoded
// @Produce application/json
// @Param id path int true "ID of task"
// @Param title formData string true "title"
// @Success 200 {object} response.Response{data=pb.Task}
// @Failure 400 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/tasks/{id}/title [patch]
func (i *impl) ChangeTitle(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var req reqID
	err := c.ShouldBindUri(&req)
	if err != nil {
		ctx.Error(er.ErrBindID.Error(), zap.Error(err))
		_ = c.Error(er.ErrBindID)
		return
	}

	title := c.PostForm("title")

	ret, err := i.biz.ChangeTitle(ctx, req.ID, title)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(ticket.NewTaskResponse(ret)))
}

// Delete
// @Summary Delete a task by id
// @Description Delete a task by id
// @Tags Tasks
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID of task"
// @Success 200 {object} response.Response{data=string}
// @Failure 400 {object} er.APPError
// @Failure 404 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/tasks/{id} [delete]
func (i *impl) Delete(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var req reqID
	err := c.ShouldBindUri(&req)
	if err != nil {
		ctx.Error(er.ErrBindID.Error(), zap.Error(err))
		_ = c.Error(er.ErrBindID)
		return
	}

	err = i.biz.Delete(ctx, req.ID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(req.ID))
}
