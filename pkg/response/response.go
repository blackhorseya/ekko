package response

import (
	"net/http"
)

var (
	// OK request is successful.
	OK = &Response{
		Code:    http.StatusOK,
		Message: "ok",
		Data:    nil,
	}

	// Err request is failed.
	Err = &Response{
		Code:    http.StatusInternalServerError,
		Message: "unknown error",
		Data:    nil,
	}
)

// Response defines the response struct.
type Response struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"msg,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// WithMessage set response message.
func (resp *Response) WithMessage(message string) *Response {
	return &Response{
		Code:    resp.Code,
		Message: message,
		Data:    resp.Data,
	}
}

// WithData set response data.
func (resp *Response) WithData(data any) *Response {
	return &Response{
		Code:    resp.Code,
		Message: resp.Message,
		Data:    data,
	}
}
