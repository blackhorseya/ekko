package er

import "net/http"

var (
	// ErrDBConnect means db connect is failure
	ErrDBConnect = newAPPError(http.StatusInternalServerError, 50001, "db connect is failure")

	// ErrPing means db ping is failure
	ErrPing = newAPPError(http.StatusInternalServerError, 50002, "db ping is failure")
)
