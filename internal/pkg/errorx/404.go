package errorx

import (
	"net/http"

	"github.com/blackhorseya/ekko/pkg/er"
)

var (
	// ErrTaskNotExists means issue is not exists
	ErrTaskNotExists = er.New(http.StatusNotFound, "issue is not exists")

	// ErrUserNotFound means user is not exists
	ErrUserNotFound = er.New(http.StatusNotFound, "user is not exists")
)
