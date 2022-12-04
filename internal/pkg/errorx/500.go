package errorx

import (
	"net/http"

	"github.com/blackhorseya/todo-app/pkg/er"
)

const (
	errDatabase = "Failed to connect to database"
)

var (
	// ErrContextx means Missing contextx
	ErrContextx = er.New(http.StatusInternalServerError, 50001, "Internal server error", "Missing contextx")

	// ErrPing means db ping is failure
	ErrPing = er.New(http.StatusInternalServerError, 50002, errDatabase, "db ping is failure")
)

var (
	// ErrGetTask means get a task by id is failure
	ErrGetTask = er.New(http.StatusInternalServerError, 50010, errDatabase, "get a task by id is failure")

	// ErrListTasks means list all tasks is failure
	ErrListTasks = er.New(http.StatusInternalServerError, 50011, errDatabase, "list all tasks is failure")

	// ErrCountTask means count task is failure
	ErrCountTask = er.New(http.StatusInternalServerError, 50012, errDatabase, "count task is failure")

	// ErrCreateTask means create a task is failure
	ErrCreateTask = er.New(http.StatusInternalServerError, 50013, errDatabase, "create a task is failure")

	// ErrDeleteTask means delete a task by id is failure
	ErrDeleteTask = er.New(http.StatusInternalServerError, 50014, errDatabase, "delete a task by id is failure")

	// ErrUpdateStatusTask means update the task's status is failure
	ErrUpdateStatusTask = er.New(http.StatusInternalServerError, 50015, errDatabase, "update the task's status is failure")

	// ErrChangeTitleTask means change the task's title is failure
	ErrChangeTitleTask = er.New(http.StatusInternalServerError, 50016, errDatabase, "change the task's title is failure")
)
