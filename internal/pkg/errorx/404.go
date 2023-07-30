package errorx

import (
	"net/http"

	"github.com/blackhorseya/ekko/pkg/er"
)

var (
	// ErrTicketNotExists means ticket is not exists
	ErrTicketNotExists = er.New(http.StatusNotFound, "ticket is not exists")

	// ErrUserNotFound means user is not exists
	ErrUserNotFound = er.New(http.StatusNotFound, "user is not exists")
)
