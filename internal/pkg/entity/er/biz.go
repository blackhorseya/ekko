package er

import "net/http"

var (
	// ErrGetUser means get user is failure
	ErrGetUser = newAPPError(http.StatusInternalServerError, 50000, "get user is failure")

	// ErrUserNotExists means user is not exists
	ErrUserNotExists = newAPPError(http.StatusNotFound, 40400, "user is not exists")

	// ErrUserAlreadyExists means user is already exists
	ErrUserAlreadyExists = newAPPError(http.StatusInternalServerError, 40400, "user is already exists")

	// ErrSignup means user signup is failure
	ErrSignup = newAPPError(http.StatusInternalServerError, 50001, "user signup is failure")

	// ErrWrongPassword means given password is wrong
	ErrWrongPassword = newAPPError(http.StatusInternalServerError, 50001, "wrong password")

	// ErrChangePassword means change password is failure
	ErrChangePassword = newAPPError(http.StatusInternalServerError, 50002, "change password is failure")

	// ErrSamePassword means given password is same origin password
	ErrSamePassword = newAPPError(http.StatusBadRequest, 50003, "password is same origin password")

	// ErrLogout means logout is failure
	ErrLogout = newAPPError(http.StatusInternalServerError, 50004, "logout is failure")

	// ErrLogin means login is failure
	ErrLogin = newAPPError(http.StatusInternalServerError, 50005, "login is failure")
)
