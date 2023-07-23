package errorx

import (
	"net/http"

	"github.com/blackhorseya/ekko/pkg/er"
)

var (
	// ErrContextx means Missing contextx
	ErrContextx = er.New(http.StatusInternalServerError, "missing contextx")
)

var (
	// ErrGetTask means get a issue by id is failure
	ErrGetTask = er.New(http.StatusInternalServerError, "get an issue by id is failure")

	// ErrListTasks means list all tasks is failure
	ErrListTasks = er.New(http.StatusInternalServerError, "list all tasks is failure")

	// ErrCountTask means count issue is failure
	ErrCountTask = er.New(http.StatusInternalServerError, "count issue is failure")

	// ErrCreateTask means create a issue is failure
	ErrCreateTask = er.New(http.StatusInternalServerError, "create a issue is failure")

	// ErrDeleteTask means delete a issue by id is failure
	ErrDeleteTask = er.New(http.StatusInternalServerError, "delete a issue by id is failure")

	// ErrUpdateStatusTask means update the issue's status is failure
	ErrUpdateStatusTask = er.New(http.StatusInternalServerError, "update the issue's status is failure")
)

var (
	// ErrRegisterProfile means register profile is failure
	ErrRegisterProfile = er.New(http.StatusInternalServerError, "register profile is failure")

	// ErrGetProfile means get profile is failure
	ErrGetProfile = er.New(http.StatusInternalServerError, "get profile is failure")

	// ErrNewToken means new token is failure
	ErrNewToken = er.New(http.StatusInternalServerError, "new token is failure")

	// ErrUpdateToken means update token is failure
	ErrUpdateToken = er.New(http.StatusInternalServerError, "update token is failure")
)
