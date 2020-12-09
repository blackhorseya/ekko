package app

import (
	"github.com/blackhorseya/todo-app/internal/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// InjectorSet inject Injector
var InjectorSet = wire.NewSet(NewInjector)

// Injector define inject something
type Injector struct {
	Engine *gin.Engine
	C      *config.Config
}

// NewInjector constructor of Injector
func NewInjector(engine *gin.Engine, c *config.Config) *Injector {
	return &Injector{Engine: engine, C: c}
}
