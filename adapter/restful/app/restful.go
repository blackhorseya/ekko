package app

import (
	"net/http"

	_ "github.com/blackhorseya/ekko/adapter/restful/api/docs" // swagger docs
	"github.com/blackhorseya/ekko/pkg/adapters"
	"github.com/blackhorseya/ekko/pkg/response"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	api := r.router.Group("/api")
	{
		api.GET("/healthz", func(c *gin.Context) {
			c.JSON(http.StatusOK, response.OK)
		})

		api.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
