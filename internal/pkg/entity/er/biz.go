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

var (
	// ErrGetTask means get a task by id is failure
	ErrGetTask = newAPPError(http.StatusInternalServerError, 50006, "get a task by id is failure")

	// ErrTaskNotExists means task is not exists
	ErrTaskNotExists = newAPPError(http.StatusNotFound, 40401, "task is not exists")

	// ErrListTasks means list all tasks is failure
	ErrListTasks = newAPPError(http.StatusInternalServerError, 50007, "list all tasks is failure")

	// ErrCountTask means count task is failure
	ErrCountTask = newAPPError(http.StatusInternalServerError, 50008, "count task is failure")

	// ErrCreateTask means create a task is failure
	ErrCreateTask = newAPPError(http.StatusInternalServerError, 50009, "create a task is failure")

	// ErrDeleteTask means delete a task by id is failure
	ErrDeleteTask = newAPPError(http.StatusInternalServerError, 50010, "delete a task by id is failure")

	// ErrUpdateStatusTask means update the task's status is failure
	ErrUpdateStatusTask = newAPPError(http.StatusInternalServerError, 50011, "update the task's status is failure")
)
