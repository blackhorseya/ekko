package er

// Error declare custom error
type Error struct {
	Status         int         `json:"-"`
	Code           int         `json:"code"`
	DisplayMessage string      `json:"msg"`
	LogMessage     string      `json:"-"`
	Data           interface{} `json:"data,omitempty"`
}

func (e *Error) Error() string {
	return e.DisplayMessage
}

// WithData append data into response
func (e *Error) WithData(data interface{}) *Error {
	return &Error{
		Status:         e.Status,
		Code:           e.Code,
		DisplayMessage: e.DisplayMessage,
		Data:           data,
	}
}

// New a error
func New(status int, code int, msg string, log string) *Error {
	return &Error{Status: status, Code: code, DisplayMessage: msg, LogMessage: log}
}
