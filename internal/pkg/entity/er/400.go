package er

import "net/http"

var (
	// ErrBindTitle means request body bind title is invalid
	ErrBindTitle = newAPPError(http.StatusBadRequest, 40010, "request body bind title is invalid")

	// ErrBindStatus means request body bind status is invalid
	ErrBindStatus = newAPPError(http.StatusBadRequest, 40011, "request body bind status is invalid")
)
