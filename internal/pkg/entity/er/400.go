package er

import "net/http"

var (
	// ErrBindTitle means request body bind title is invalid
	ErrBindTitle = newAPPError(http.StatusBadRequest, 40010, "request body bind title is invalid")

	// ErrBindStatus means request body bind status is invalid
	ErrBindStatus = newAPPError(http.StatusBadRequest, 40011, "request body bind status is invalid")

	// ErrInvalidID means id is invalid
	ErrInvalidID = newAPPError(http.StatusBadRequest, 40012, "id is invalid")

	// ErrEmptyID means id must be NOT empty
	ErrEmptyID = newAPPError(http.StatusBadRequest, 40013, "id must be NOT empty")

	// ErrEmptyTitle means title must be NOT empty
	ErrEmptyTitle = newAPPError(http.StatusBadRequest, 40014, "title must be NOT empty")
)
