package errorx

// Error is a custom error type.
type Error struct {
	StatusCode int   `json:"status_code"`
	Code       int   `json:"code"`
	Err        error `json:"error"`
}

func (e *Error) Error() string {
	return e.Err.Error()
}

// Wrap wraps an error with a code.
func Wrap(status, code int, err error) error {
	return &Error{
		StatusCode: status,
		Code:       code,
		Err:        err,
	}
}
