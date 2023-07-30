package app

import (
	"fmt"
	"net/http"
	"time"

	_ "github.com/blackhorseya/ekko/adapter/restful/api/docs" // swagger docs
	v1 "github.com/blackhorseya/ekko/adapter/restful/app/v1"
	issueB "github.com/blackhorseya/ekko/entity/domain/issue/biz"
	taskB "github.com/blackhorseya/ekko/entity/domain/task/biz"
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
	taskB  taskB.IBiz
}

// NewRestful will create a restful adapter
func NewRestful(logger *zap.Logger, router *gin.Engine, issue issueB.IBiz, taskB taskB.IBiz) adapters.Restful {
	router.Use(ginzap.GinzapWithConfig(logger, &ginzap.Config{
		TimeFormat: time.RFC3339,
		UTC:        true,
		SkipPaths:  []string{"/api/healthz"},
		TraceID:    false,
		Context:    nil,
	}))
	router.Use(ginzap.CustomRecoveryWithZap(logger, true, func(c *gin.Context, err any) {
		msg := fmt.Sprintf("%v", err)
		resp := er.New(http.StatusInternalServerError, msg)
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp)
	}))
	router.Use(contextx.WithContextx(logger))
	router.Use(er.HandleError())

	return &restful{
		logger: logger,
		router: router,
		issue:  issue,
		taskB:  taskB,
	}
}

func (r *restful) InitRouting() {
	api := r.router.Group("/api")
	{
		api.GET("/healthz", func(c *gin.Context) {
			c.JSON(http.StatusOK, response.OK)
		})

		api.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		v1.Handle(api.Group("/v1"), r.issue, r.taskB)
	}
}
