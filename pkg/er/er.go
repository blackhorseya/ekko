package er

// Error is a struct for error
type Error struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"msg,omitempty"`
}

func (e *Error) Error() string {
	return e.Message
}

// New is a constructor for Error
func New(code int, msg string) *Error {
	return &Error{
		Code:    code,
		Message: msg,
	}
}
