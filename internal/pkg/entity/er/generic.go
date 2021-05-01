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
