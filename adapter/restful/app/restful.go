package app

import (
	"net/http"
	"time"

	_ "github.com/blackhorseya/ekko/adapter/restful/api/docs" // swagger docs
	v1 "github.com/blackhorseya/ekko/adapter/restful/app/v1"
	issueB "github.com/blackhorseya/ekko/entity/domain/issue/biz"
	"github.com/blackhorseya/ekko/pkg/adapters"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/blackhorseya/ekko/pkg/er"
	"github.com/blackhorseya/ekko/pkg/response"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

type restful struct {
	logger *zap.Logger
	router *gin.Engine
	issue  issueB.IBiz
}

// NewRestful will create a restful adapter
func NewRestful(logger *zap.Logger, router *gin.Engine, issue issueB.IBiz) adapters.Restful {
	router.Use(ginzap.GinzapWithConfig(logger, &ginzap.Config{
		TimeFormat: time.RFC3339,
		UTC:        true,
		SkipPaths:  []string{"/api/healthz"},
		TraceID:    false,
		Context:    nil,
	}))
	router.Use(ginzap.RecoveryWithZap(logger, true))
	router.Use(contextx.WithContextx(logger))
	router.Use(er.HandleError())

	return &restful{
		logger: logger,
		router: router,
		issue:  issue,
	}
}

func (r *restful) InitRouting() {
	api := r.router.Group("/api")
	{
		api.GET("/healthz", func(c *gin.Context) {
			c.JSON(http.StatusOK, response.OK)
		})

		api.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		v1.Handle(api.Group("/v1"), r.issue)
	}
}