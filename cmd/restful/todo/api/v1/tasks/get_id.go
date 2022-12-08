package tasks

import (
	"github.com/gin-gonic/gin"
)

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
	// todo: 2022/12/8|sean|impl me
}
