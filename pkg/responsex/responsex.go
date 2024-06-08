package responsex

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response is a struct that represents the response of the API.
type Response struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

// OK is a function that returns a response with status code 200.
func OK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "ok",
		Data:    data,
	})
}

// Err is a function that returns a response with status code 500.
func Err(c *gin.Context, err error) {
	_ = c.Error(err)
}
