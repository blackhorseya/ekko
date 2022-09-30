package restful

import (
	// import swagger docs
	_ "github.com/blackhorseya/todo-app/api/docs"
	"github.com/blackhorseya/todo-app/internal/app/todo/api/restful/health"
	"github.com/blackhorseya/todo-app/internal/app/todo/api/restful/todo"
	"github.com/blackhorseya/todo-app/internal/pkg/infra/transports/http"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// CreateInitHandlerFn serve caller to create init handler
func CreateInitHandlerFn(healthH health.IHandler, todoH todo.IHandler) http.InitHandlers {
	return func(r *gin.Engine) {
		api := r.Group("api")
		{
			// open any environments can access swagger
			api.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}
	}
}

// ProviderSet is an apis provider set
var ProviderSet = wire.NewSet(
	health.ProviderSet,
	todo.ProviderSet,
	CreateInitHandlerFn,
)
