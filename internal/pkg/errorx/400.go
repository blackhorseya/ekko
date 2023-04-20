package errorx

import (
	"net/http"

	"github.com/blackhorseya/ekko/pkg/er"
)

var (
	// ErrInvalidPage means given start is invalid
	ErrInvalidPage = er.New(http.StatusBadRequest, 40000, "start is invalid", "start is invalid")

	// ErrInvalidSize means given end is invalid
	ErrInvalidSize = er.New(http.StatusBadRequest, 40001, "end is invalid", "end is invalid")
)

var (
	// ErrInvalidID means bind id is invalid
	ErrInvalidID = er.New(http.StatusBadRequest, 40012, "id is invalid", "id is invalid")

	// ErrInvalidTitle means title must be NOT empty
	ErrInvalidTitle = er.New(http.StatusBadRequest, 40015, "title is invalid", "title is invalid")

	// ErrInvalidStatus means status value is invalid
	ErrInvalidStatus = er.New(http.StatusBadRequest, 40016, "status value is invalid", "status value is invalid")
)

var (
	// ErrInvalidUsername means username is invalid
	ErrInvalidUsername = er.New(http.StatusBadRequest, 40020, "username is invalid", "username is invalid")

	// ErrInvalidPassword means password is invalid
	ErrInvalidPassword = er.New(http.StatusBadRequest, 40021, "password is invalid", "password is invalid")

	// ErrInvalidProfile means profile is invalid
	ErrInvalidProfile = er.New(http.StatusBadRequest, 40022, "profile is invalid", "profile is invalid")

	// ErrInvalidToken means token is invalid
	ErrInvalidToken = er.New(http.StatusBadRequest, 40023, "token is invalid", "token is invalid")
)
