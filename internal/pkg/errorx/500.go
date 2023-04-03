package errorx

import (
	"net/http"

	"github.com/blackhorseya/ekko/pkg/er"
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
	// ErrGetTask means get a issue by id is failure
	ErrGetTask = er.New(http.StatusInternalServerError, 50010, errDatabase, "get a issue by id is failure")

	// ErrListTasks means list all tasks is failure
	ErrListTasks = er.New(http.StatusInternalServerError, 50011, errDatabase, "list all tasks is failure")

	// ErrCountTask means count issue is failure
	ErrCountTask = er.New(http.StatusInternalServerError, 50012, errDatabase, "count issue is failure")

	// ErrCreateTask means create a issue is failure
	ErrCreateTask = er.New(http.StatusInternalServerError, 50013, errDatabase, "create a issue is failure")

	// ErrDeleteTask means delete a issue by id is failure
	ErrDeleteTask = er.New(http.StatusInternalServerError, 50014, errDatabase, "delete a issue by id is failure")

	// ErrUpdateStatusTask means update the issue's status is failure
	ErrUpdateStatusTask = er.New(http.StatusInternalServerError, 50015, errDatabase, "update the issue's status is failure")

	// ErrChangeTitleTask means change the issue's title is failure
	ErrChangeTitleTask = er.New(http.StatusInternalServerError, 50016, errDatabase, "change the issue's title is failure")
)
