package restful

import (
	"time"

	"github.com/blackhorseya/todo-app/internal/adapter/task/restful/api"
	"github.com/blackhorseya/todo-app/pkg/adapters"
	"github.com/blackhorseya/todo-app/pkg/contextx"
	"github.com/blackhorseya/todo-app/pkg/cors"
	tb "github.com/blackhorseya/todo-app/pkg/entity/domain/task/biz"
	"github.com/blackhorseya/todo-app/pkg/er"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go.uber.org/zap"
)

type restful struct {
	router *gin.Engine
	biz    tb.IBiz
}

func NewRestful(logger *zap.Logger, router *gin.Engine, biz tb.IBiz) adapters.Restful {
	router.Use(cors.AddAllowAll())
	router.Use(ginzap.RecoveryWithZap(logger, true))
	router.Use(ginzap.GinzapWithConfig(logger, &ginzap.Config{
		TimeFormat: time.RFC3339,
		UTC:        true,
		SkipPaths:  []string{"/api/readiness", "/api/liveness"},
	}))
	router.Use(contextx.AddContextxWitLoggerMiddleware(logger))
	router.Use(er.AddErrorHandlingMiddleware())

	return &restful{
		router: router,
		biz:    biz,
	}
}

func (i *restful) InitRouting() error {
	api.Handle(i.router.Group("/api"), i.biz)

	return nil
}

var TaskSet = wire.NewSet(NewRestful)