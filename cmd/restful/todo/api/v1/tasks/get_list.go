package tasks

import (
	"github.com/gin-gonic/gin"
)

// List
// @Summary List all tasks
// @Description List all tasks
// @Tags Tasks
// @Accept application/json
// @Produce application/json
// @Param page query integer false "page" default(1)
// @Param size query integer false "size" default(10)
// @Success 200 {object} response.Response{data=[]model.Task}
// @Failure 400 {object} er.Error
// @Failure 404 {object} er.Error
// @Failure 500 {object} er.Error
// @Router /v1/tasks [get]
func (i *impl) List(c *gin.Context) {
	// todo: 2022/12/8|sean|impl me
}
