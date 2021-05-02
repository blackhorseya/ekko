package apis

import (
	// import swagger docs
	_ "github.com/blackhorseya/todo-app/api/docs"
	"github.com/blackhorseya/todo-app/internal/app/todo/apis/health"
	"github.com/blackhorseya/todo-app/internal/app/todo/apis/todo"
	"github.com/blackhorseya/todo-app/internal/pkg/infra/transports/http"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// CreateInitHandlerFn serve caller to create init handler
func CreateInitHandlerFn(
	healthH health.IHandler,
	todoH todo.IHandler) http.InitHandlers {
	return func(r *gin.Engine) {
		api := r.Group("api")
		{
			api.GET("readiness", healthH.Readiness)
			api.GET("liveness", healthH.Liveness)

			// open any environments can access swagger
			api.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

			v1 := api.Group("v1")
			{
				taskG := v1.Group("tasks")
				{
					taskG.GET("", todoH.List)
					taskG.GET(":id", todoH.GetByID)
					taskG.POST("", todoH.Create)
					taskG.PATCH(":id/status", todoH.UpdateStatus)
					taskG.PATCH(":id/title", todoH.ChangeTitle)
					taskG.DELETE(":id", todoH.Delete)
				}
			}
		}
	}
}

// ProviderSet is an apis provider set
var ProviderSet = wire.NewSet(
	health.ProviderSet,
	todo.ProviderSet,
	CreateInitHandlerFn,
)
