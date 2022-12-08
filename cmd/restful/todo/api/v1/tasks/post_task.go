package tasks

import (
	"github.com/gin-gonic/gin"
)

// Create
// @Summary Create a task
// @Description Create a task
// @Tags Tasks
// @Accept application/x-www-form-urlencoded
// @Produce application/json
// @Param title formData string true "title"
// @Success 201 {object} response.Response{data=model.Task}
// @Failure 400 {object} er.Error
// @Failure 500 {object} er.Error
// @Router /v1/tasks [post]
func (i *impl) Create(c *gin.Context) {
	// todo: 2022/12/8|sean|impl me
}
