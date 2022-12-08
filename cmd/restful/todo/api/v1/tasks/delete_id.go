package tasks

import (
	"github.com/gin-gonic/gin"
)

// Delete
// @Summary Delete a task by id
// @Description Delete a task by id
// @Tags Tasks
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID of task"
// @Success 200 {object} response.Response{data=string}
// @Failure 400 {object} er.Error
// @Failure 404 {object} er.Error
// @Failure 500 {object} er.Error
// @Router /v1/tasks/{id} [delete]
func (i *impl) Delete(c *gin.Context) {
}
