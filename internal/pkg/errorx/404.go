package errorx

import (
	"net/http"

	"github.com/blackhorseya/todo-app/pkg/er"
)

var (
	// ErrTaskNotExists means task is not exists
	ErrTaskNotExists = er.New(http.StatusNotFound, 40401, "task is not exists", "task is not exists")
)
