package restful

import (
	"time"

	"github.com/blackhorseya/ekko/internal/adapter/issue/restful/api"
	"github.com/blackhorseya/ekko/pkg/adapters"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/blackhorseya/ekko/pkg/cors"
	ib "github.com/blackhorseya/ekko/pkg/entity/domain/issue/biz"
	"github.com/blackhorseya/ekko/pkg/er"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type restful struct {
	router *gin.Engine
	biz    ib.IBiz
}

func NewRestful(logger *zap.Logger, router *gin.Engine, biz ib.IBiz) adapters.Restful {
	router.Use(cors.AddAllowAll())
	router.Use(ginzap.RecoveryWithZap(logger, true))
	router.Use(ginzap.GinzapWithConfig(logger, &ginzap.Config{
		TimeFormat: time.RFC3339,
		UTC:        true,
		SkipPaths:  []string{"/api/readiness", "/api/liveness"},
	}))
	router.Use(contextx.WithContextx(logger))
	router.Use(er.HandleError())

	panic("implement me")
}

func (i *restful) InitRouting() error {
	api.Handle(i.router.Group("/api"), i.biz)

	return nil
}
