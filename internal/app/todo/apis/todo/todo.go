package todo

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// IHandler declare handler service function
type IHandler interface {
	// GetByID serve caller to given id to get a task
	GetByID(c *gin.Context)

	// List serve caller to list all task by start and end
	List(c *gin.Context)

	// Create serve caller to create a task
	Create(c *gin.Context)

	// UpdateStatus serve caller to update the task's status by id
	UpdateStatus(c *gin.Context)

	// ChangeTitle serve caller to change the task's title by id
	ChangeTitle(c *gin.Context)

	// Delete serve caller to delete a task by id
	Delete(c *gin.Context)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl)
