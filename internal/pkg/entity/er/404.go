package er

import (
	"net/http"

	"github.com/blackhorseya/gocommon/pkg/er"
)

var (
	// ErrTaskNotExists means task is not exists
	ErrTaskNotExists = er.NewAPPError(http.StatusNotFound, 40401, "task is not exists")
)
