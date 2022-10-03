package er

import (
	"net/http"

	"github.com/blackhorseya/gocommon/pkg/er"
)

var (
	// ErrInvalidPage means given start is invalid
	ErrInvalidPage = er.NewAPPError(http.StatusBadRequest, 40000, "start is invalid")

	// ErrInvalidSize means given end is invalid
	ErrInvalidSize = er.NewAPPError(http.StatusBadRequest, 40001, "end is invalid")

	// ErrMissingID means given id is empty
	ErrMissingID = er.NewAPPError(http.StatusBadRequest, 40002, "id is empty")

	// ErrMissingEmail means given email is empty
	ErrMissingEmail = er.NewAPPError(http.StatusBadRequest, 40003, "email is empty")

	// ErrMissingToken means given token is empty
	ErrMissingToken = er.NewAPPError(http.StatusBadRequest, 40004, "token is empty")

	// ErrMissingPassword means given password is empty
	ErrMissingPassword = er.NewAPPError(http.StatusBadRequest, 40005, "password is empty")
)

var (
	// ErrBindTitle means request body bind title is invalid
	ErrBindTitle = er.NewAPPError(http.StatusBadRequest, 40010, "request body bind title is invalid")

	// ErrBindStatus means request body bind status is invalid
	ErrBindStatus = er.NewAPPError(http.StatusBadRequest, 40011, "request body bind status is invalid")

	// ErrBindID means bind id is invalid
	ErrBindID = er.NewAPPError(http.StatusBadRequest, 40012, "bind id is invalid")

	// ErrInvalidID means id is invalid
	ErrInvalidID = er.NewAPPError(http.StatusBadRequest, 40013, "id is invalid")

	// ErrEmptyID means id must be NOT empty
	ErrEmptyID = er.NewAPPError(http.StatusBadRequest, 40014, "id must be NOT empty")

	// ErrEmptyTitle means title must be NOT empty
	ErrEmptyTitle = er.NewAPPError(http.StatusBadRequest, 40015, "title must be NOT empty")

	// ErrInvalidStatus means status value is invalid
	ErrInvalidStatus = er.NewAPPError(http.StatusBadRequest, 40016, "status value is invalid")
)
