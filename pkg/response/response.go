package response

import (
	"net/http"
)

// Response declare unite adapters response format
type Response struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"msg,omitempty"`
	Data    any    `json:"data,omitempty"`
}

// WithMessage append message into response
func (resp *Response) WithMessage(message string) *Response {
	return &Response{
		Code:    resp.Code,
		Message: message,
		Data:    resp.Data,
	}
}

// WithData append data into response
func (resp *Response) WithData(data interface{}) *Response {
	return &Response{
		Code:    resp.Code,
		Message: resp.Message,
		Data:    data,
	}
}

var (
	// OK request is success
	OK = &Response{Code: http.StatusOK, Message: "ok"}

	// Err request is failed
	Err = &Response{Code: http.StatusInternalServerError, Message: "internal server error"}
)
