package v1

import (
	"github.com/blackhorseya/ekko/adapter/platform/rest/v1/tickets"
	"github.com/blackhorseya/ekko/adapter/platform/rest/v1/todos"
	"github.com/blackhorseya/ekko/adapter/platform/wirex"
	"github.com/gin-gonic/gin"
)

// Handle is used to handle the v1 restful api.
func Handle(g *gin.RouterGroup, injector *wirex.Injector) {
	tickets.Handle(g, injector)
	todos.Handle(g.Group("/todos"), injector)
}
