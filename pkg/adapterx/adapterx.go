package adapterx

import (
	"github.com/gin-gonic/gin"
)

// Servicer is the interface that wraps the basic Serve method.
type Servicer interface {
	// Start a service asynchronously.
	Start() error

	// AwaitSignal waits for a signal to shut down the service.
	AwaitSignal() error
}

// Restful is the interface that wraps the restful api method.
type Restful interface {
	Servicer

	// InitRouting init the routing of restful api.
	InitRouting() error

	// GetRouter returns the router of restful api.
	GetRouter() *gin.Engine
}
