package er

// APPError declare custom error
type APPError struct {
	Status int    `json:"-"`
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
}

func (e *APPError) Error() string {
	return e.Msg
}

func newAPPError(status int, code int, msg string) *APPError {
	return &APPError{Status: status, Code: code, Msg: msg}
}
