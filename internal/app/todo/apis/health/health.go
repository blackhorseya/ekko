package health

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// IHandler declare health api handler
type IHandler interface {
	Readiness(c *gin.Context)
	Liveness(c *gin.Context)
}

// ProviderSet is a health of api provider set
var ProviderSet = wire.NewSet(NewImpl)
