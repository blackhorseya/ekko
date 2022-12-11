package tasks

import (
	"github.com/gin-gonic/gin"
)

// ChangeTitle
// @Summary Change task's title by id
// @Description Change task's title by id
// @Tags Tasks
// @Accept application/x-www-form-urlencoded
// @Produce application/json
// @Param id path int true "ID of task"
// @Param title formData string true "title"
// @Success 200 {object} response.Response{data=model.Task}
// @Failure 400 {object} er.Error
// @Failure 500 {object} er.Error
// @Router /v1/tasks/{id}/title [patch]
func (i *impl) ChangeTitle(c *gin.Context) {
	// todo: 2022/12/8|sean|impl me
}
