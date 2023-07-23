package errorx

import (
	"net/http"

	"github.com/blackhorseya/ekko/pkg/er"
)

var (
	// ErrInvalidPage means given page is invalid
	ErrInvalidPage = er.New(http.StatusBadRequest, "page is invalid")

	// ErrInvalidSize means given size is invalid
	ErrInvalidSize = er.New(http.StatusBadRequest, "size is invalid")
)

var (
	// ErrInvalidID means bind id is invalid
	ErrInvalidID = er.New(http.StatusBadRequest, "id is invalid")

	// ErrInvalidTitle means title must be NOT empty
	ErrInvalidTitle = er.New(http.StatusBadRequest, "title is invalid")

	// ErrInvalidStatus means status value is invalid
	ErrInvalidStatus = er.New(http.StatusBadRequest, "status value is invalid")
)

var (
	// ErrInvalidUsername means username is invalid
	ErrInvalidUsername = er.New(http.StatusBadRequest, "username is invalid")

	// ErrInvalidPassword means password is invalid
	ErrInvalidPassword = er.New(http.StatusBadRequest, "password is invalid")

	// ErrInvalidProfile means profile is invalid
	ErrInvalidProfile = er.New(http.StatusBadRequest, "profile is invalid")

	// ErrInvalidToken means token is invalid
	ErrInvalidToken = er.New(http.StatusBadRequest, "token is invalid")
)
