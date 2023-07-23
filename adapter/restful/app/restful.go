package app

import (
	"github.com/blackhorseya/ekko/pkg/adapters"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type restful struct {
	logger *zap.Logger
	router *gin.Engine
}

// NewRestful will create a restful adapter
func NewRestful(logger *zap.Logger, router *gin.Engine) adapters.Restful {
	return &restful{
		logger: logger,
		router: router,
	}
}

func (r *restful) InitRouting() {
	// todo: 2023/7/24|sean|impl me
	panic("implement me")
}
