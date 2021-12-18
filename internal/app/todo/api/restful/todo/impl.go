package todo

import (
	"net/http"
	"strconv"

	"github.com/blackhorseya/todo-app/internal/app/todo/biz/todo"
	"github.com/blackhorseya/todo-app/internal/pkg/base/contextx"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/er"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/response"
	"github.com/blackhorseya/todo-app/internal/pkg/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// GetByID
// @Summary Get a task by id
// @Description Get a task by id
// @Tags Tasks
// @Accept application/json
// @Produce application/json
// @Param id path string true "ID of task"
// @Success 200 {object} response.Response{data=models.TaskResponse}
// @Failure 400 {object} er.APPError
// @Failure 404 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/tasks/{id} [get]
func (i *impl) GetByID(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var req reqID
	err := c.ShouldBindUri(&req)
	if err != nil {
		i.logger.Error(er.ErrBindID.Error(), zap.Error(err))
		_ = c.Error(er.ErrBindID)
		return
	}

	id, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		i.logger.Error(er.ErrInvalidID.Error(), zap.Error(err), zap.String("id", req.ID))
		_ = c.Error(er.ErrInvalidID)
		return
	}

	ret, err := i.biz.GetByID(ctx, id)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(models.NewTaskResponse(ret)))
}

// List
// @Summary List all tasks
// @Description List all tasks
// @Tags Tasks
// @Accept application/json
// @Produce application/json
// @Param start query integer false "start" default(0)
// @Param end query integer false "end" default(10)
// @Success 200 {object} response.Response{data=[]models.TaskResponse}
// @Failure 400 {object} er.APPError
// @Failure 404 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/tasks [get]
func (i *impl) List(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	start, err := strconv.Atoi(c.DefaultQuery("start", "0"))
	if err != nil {
		i.logger.Error(er.ErrInvalidStart.Error(), zap.Error(err), zap.String("start", c.Query("start")))
		_ = c.Error(er.ErrInvalidStart)
		return
	}

	end, err := strconv.Atoi(c.DefaultQuery("end", "10"))
	if err != nil {
		i.logger.Error(er.ErrInvalidEnd.Error(), zap.Error(err), zap.String("end", c.Query("end")))
		_ = c.Error(er.ErrInvalidEnd)
		return
	}

	tasks, total, err := i.biz.List(ctx, start, end)
	if err != nil {
		_ = c.Error(err)
		return
	}

	var ret []*models.TaskResponse
	for _, task := range tasks {
		ret = append(ret, models.NewTaskResponse(task))
	}

	c.Header("X-Total-Count", strconv.Itoa(total))
	c.JSON(http.StatusOK, response.OK.WithData(ret))
}

// Create
// @Summary Create a task
// @Description Create a task
// @Tags Tasks
// @Accept application/json
// @Produce application/json
// @Param created body reqTitle true "created task"
// @Success 201 {object} response.Response{data=models.TaskResponse}
// @Failure 400 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/tasks [post]
func (i *impl) Create(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var data *reqTitle
	if err := c.ShouldBindJSON(&data); err != nil {
		i.logger.Error(er.ErrBindTitle.Error(), zap.Error(err))
		_ = c.Error(er.ErrBindTitle)
		return
	}

	ret, err := i.biz.Create(ctx, data.Title)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, response.OK.WithData(models.NewTaskResponse(ret)))
}

// UpdateStatus
// @Summary Update task's status by id
// @Description Update task's status by id
// @Tags Tasks
// @Accept application/json
// @Produce application/json
// @Param id path string true "ID of task"
// @Param updated body reqStatus true "updated task"
// @Success 200 {object} response.Response{data=models.TaskResponse}
// @Failure 400 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/tasks/{id}/status [patch]
func (i *impl) UpdateStatus(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var req reqID
	err := c.ShouldBindUri(&req)
	if err != nil {
		i.logger.Error(er.ErrBindID.Error(), zap.Error(err))
		_ = c.Error(er.ErrBindID)
		return
	}

	id, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		i.logger.Error(er.ErrInvalidID.Error(), zap.Error(err), zap.String("id", req.ID))
		_ = c.Error(er.ErrInvalidID)
		return
	}

	var data *reqStatus
	err = c.ShouldBindJSON(&data)
	if err != nil {
		i.logger.Error(er.ErrBindStatus.Error(), zap.Error(err))
		_ = c.Error(er.ErrBindStatus)
		return
	}

	ret, err := i.biz.UpdateStatus(ctx, id, data.Status)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(models.NewTaskResponse(ret)))
}

// ChangeTitle
// @Summary Change task's title by id
// @Description Change task's title by id
// @Tags Tasks
// @Accept application/json
// @Produce application/json
// @Param id path string true "ID of task"
// @Param updated body reqTitle true "updated task"
// @Success 200 {object} response.Response{data=models.TaskResponse}
// @Failure 400 {object} er.APPError
// @Failure 500 {object} er.APPError
// @Router /v1/tasks/{id}/title [patch]
func (i *impl) ChangeTitle(c *gin.Context) {
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var req reqID
	err := c.ShouldBindUri(&req)
	if err != nil {
		i.logger.Error(er.ErrBindID.Error(), zap.Error(err))
		_ = c.Error(er.ErrBindID)
		return
	}

	id, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		i.logger.Error(er.ErrInvalidID.Error(), zap.Error(err), zap.String("id", req.ID))
		_ = c.Error(er.ErrInvalidID)
		return
	}

	var data *reqTitle
	err = c.ShouldBindJSON(&data)
	if err != nil {
		i.logger.Error(er.ErrBindTitle.Error(), zap.Error(err))
		_ = c.Error(er.ErrBindTitle)
		return
	}

	ret, err := i.biz.ChangeTitle(ctx, id, data.Title)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(models.NewTaskResponse(ret)))
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
	ctx := c.MustGet("ctx").(contextx.Contextx)

	var req reqID
	err := c.ShouldBindUri(&req)
	if err != nil {
		i.logger.Error(er.ErrBindID.Error(), zap.Error(err))
		_ = c.Error(er.ErrBindID)
		return
	}

	id, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		i.logger.Error(er.ErrInvalidID.Error(), zap.Error(err), zap.String("id", req.ID))
		_ = c.Error(er.ErrInvalidID)
		return
	}

	err = i.biz.Delete(ctx, id)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.Status(http.StatusNoContent)
}
