package er

import "net/http"

var (
	// ErrPing means db ping is failure
	ErrPing = newAPPError(http.StatusInternalServerError, 50001, "db ping is failure")
)

var (
	// ErrGetTask means get a task by id is failure
	ErrGetTask = newAPPError(http.StatusInternalServerError, 50010, "get a task by id is failure")

	// ErrListTasks means list all tasks is failure
	ErrListTasks = newAPPError(http.StatusInternalServerError, 50011, "list all tasks is failure")

	// ErrCountTask means count task is failure
	ErrCountTask = newAPPError(http.StatusInternalServerError, 50012, "count task is failure")

	// ErrCreateTask means create a task is failure
	ErrCreateTask = newAPPError(http.StatusInternalServerError, 50013, "create a task is failure")

	// ErrDeleteTask means delete a task by id is failure
	ErrDeleteTask = newAPPError(http.StatusInternalServerError, 50014, "delete a task by id is failure")

	// ErrUpdateStatusTask means update the task's status is failure
	ErrUpdateStatusTask = newAPPError(http.StatusInternalServerError, 50015, "update the task's status is failure")

	// ErrChangeTitleTask means change the task's title is failure
	ErrChangeTitleTask = newAPPError(http.StatusInternalServerError, 50016, "change the task's title is failure")
)
