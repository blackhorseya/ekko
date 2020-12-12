package apis

import (
	"net/http"

	"github.com/blackhorseya/todo-app/internal/app/biz/task"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// TaskSet is a Task provider set
var TaskSet = wire.NewSet(wire.Struct(new(Task), "*"))

// Task define task apis
type Task struct {
	TaskBiz task.Biz
}

// List all tasks
// @Summary List
// @Description list all tasks
// @Tags Task
// @Accept application/json
// @Produce application/json
// @Param page query integer false "page" default(1)
// @Param size query integer false "size of page" default(3)
// @Success 200 {string} string "success"
// @Router /v1/tasks [get]
func (t *Task) List(c *gin.Context) {
	// todo: 2020-12-12|21:58|doggy|implement me
	page := c.DefaultQuery("page", "1")
	size := c.DefaultQuery("size", "3")

	c.String(http.StatusOK, "page: %s, size: %s", page, size)
}

// Create a task
// @Summary Create
// @Description create a task
// @Tags Task
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "success"
// @Router /v1/tasks [post]
func (t *Task) Create(c *gin.Context) {
	// todo: 2020-12-12|21:58|doggy|implement me
	c.String(http.StatusCreated, "ok")
}
