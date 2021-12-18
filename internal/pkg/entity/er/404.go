package er

import "net/http"

var (
	// ErrTaskNotExists means task is not exists
	ErrTaskNotExists = newAPPError(http.StatusNotFound, 40401, "task is not exists")
)
