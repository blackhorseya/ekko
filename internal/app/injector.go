package app

import (
	"github.com/blackhorseya/todo-app/internal/app/config"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// InjectorSet inject Injector
var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

// Injector define inject something
type Injector struct {
	Engine *gin.Engine
	C      *config.Config
}
