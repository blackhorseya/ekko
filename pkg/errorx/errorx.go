package errorx

import (
	"errors"
)

// Error is a custom error type.
type Error struct {
	StatusCode int   `json:"status_code"`
	Code       int   `json:"code"`
	Err        error `json:"error"`
}

func (e *Error) Error() string {
	return e.Err.Error()
}

// New creates a new error with a code.
func New(status, code int, message string) error {
	return Wrap(status, code, errors.New(message))
}

// Wrap wraps an error with a code.
func Wrap(status, code int, err error) error {
	return &Error{
		StatusCode: status,
		Code:       code,
		Err:        err,
	}
}
