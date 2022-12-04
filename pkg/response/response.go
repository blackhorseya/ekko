package response

import (
	"net/http"
)

// Response declare unite adapters response format
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// WithMessage append message into response
func (resp *Response) WithMessage(message string) *Response {
	return &Response{
		Code: resp.Code,
		Msg:  message,
		Data: resp.Data,
	}
}

// WithData append data into response
func (resp *Response) WithData(data interface{}) *Response {
	return &Response{
		Code: resp.Code,
		Msg:  resp.Msg,
		Data: data,
	}
}

var (
	// OK request is success
	OK = &Response{Code: http.StatusOK, Msg: "ok"}
)
