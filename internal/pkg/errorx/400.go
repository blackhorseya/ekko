package errorx

import (
	"net/http"

	"github.com/blackhorseya/todo-app/pkg/er"
)

var (
	// ErrInvalidPage means given start is invalid
	ErrInvalidPage = er.New(http.StatusBadRequest, 40000, "start is invalid", "start is invalid")

	// ErrInvalidSize means given end is invalid
	ErrInvalidSize = er.New(http.StatusBadRequest, 40001, "end is invalid", "end is invalid")
)

var (
	// ErrBindID means bind id is invalid
	ErrBindID = er.New(http.StatusBadRequest, 40012, "bind id is invalid", "bind id is invalid")

	// ErrEmptyTitle means title must be NOT empty
	ErrEmptyTitle = er.New(http.StatusBadRequest, 40015, "title must be NOT empty", "title must be NOT empty")

	// ErrInvalidStatus means status value is invalid
	ErrInvalidStatus = er.New(http.StatusBadRequest, 40016, "status value is invalid", "status value is invalid")
)
