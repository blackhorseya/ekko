package response

import (
	"encoding/json"
	"net/http"
)

// Response declare unite api response format
type Response struct {
	Code int         `json:"code,omitempty"`
	Msg  string      `json:"msg,omitempty"`
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

func (resp *Response) String() string {
	ret, _ := json.Marshal(resp)
	return string(ret)
}

func newResponse(code int, msg string) *Response {
	return &Response{Code: code, Msg: msg}
}

var (
	// OK request is success
	OK = newResponse(http.StatusOK, "ok")
)
