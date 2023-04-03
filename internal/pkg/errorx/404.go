package errorx

import (
	"net/http"

	"github.com/blackhorseya/ekko/pkg/er"
)

var (
	// ErrTaskNotExists means issue is not exists
	ErrTaskNotExists = er.New(http.StatusNotFound, 40401, "issue is not exists", "issue is not exists")
)
