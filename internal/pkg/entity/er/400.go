package er

import "net/http"

var (
	// ErrInvalidStart means given start is invalid
	ErrInvalidStart = newAPPError(http.StatusBadRequest, 40000, "start is invalid")

	// ErrInvalidEnd means given end is invalid
	ErrInvalidEnd = newAPPError(http.StatusBadRequest, 40001, "end is invalid")

	// ErrMissingID means given id is empty
	ErrMissingID = newAPPError(http.StatusBadRequest, 40002, "id is empty")

	// ErrMissingEmail means given email is empty
	ErrMissingEmail = newAPPError(http.StatusBadRequest, 40003, "email is empty")

	// ErrMissingToken means given token is empty
	ErrMissingToken = newAPPError(http.StatusBadRequest, 40004, "token is empty")

	// ErrMissingPassword means given password is empty
	ErrMissingPassword = newAPPError(http.StatusBadRequest, 40005, "password is empty")
)

var (
	// ErrBindTitle means request body bind title is invalid
	ErrBindTitle = newAPPError(http.StatusBadRequest, 40010, "request body bind title is invalid")

	// ErrBindStatus means request body bind status is invalid
	ErrBindStatus = newAPPError(http.StatusBadRequest, 40011, "request body bind status is invalid")

	// ErrBindID means bind id is invalid
	ErrBindID = newAPPError(http.StatusBadRequest, 40012, "bind id is invalid")

	// ErrInvalidID means id is invalid
	ErrInvalidID = newAPPError(http.StatusBadRequest, 40013, "id is invalid")

	// ErrEmptyID means id must be NOT empty
	ErrEmptyID = newAPPError(http.StatusBadRequest, 40014, "id must be NOT empty")

	// ErrEmptyTitle means title must be NOT empty
	ErrEmptyTitle = newAPPError(http.StatusBadRequest, 40015, "title must be NOT empty")
)
