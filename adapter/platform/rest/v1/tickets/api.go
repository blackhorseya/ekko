package tickets

import (
	"errors"

	"github.com/blackhorseya/ekko/adapter/platform/wirex"
	"github.com/blackhorseya/ekko/entity/domain/task/model"
	"github.com/gin-gonic/gin"
)

type impl struct {
	injector *wirex.Injector
}

// Handle is used to handle the tickets restful api.
func Handle(g *gin.RouterGroup, injector *wirex.Injector) {
	instance := &impl{
		injector: injector,
	}

	tickets := g.Group("/tickets")
	{
		tickets.POST("", instance.CreateTicket)
		tickets.GET("", instance.ListTicket)
		tickets.GET("/:id", instance.GetTicket)
		tickets.PUT("/:id", instance.UpdateTicket)
		tickets.DELETE("/:id", instance.DeleteTicket)
	}
}

// CreateTicketPayload defines the create payload.
type CreateTicketPayload struct {
	Title string `json:"title" binding:"required" example:"task title"`
}

// CreateTicket is used to create a ticket.
// @Summary Create a ticket
// @Description create a ticket
// @Tags tickets
// @Accept json
// @Produce json
// @Param payload body CreateTicketPayload true "create ticket payload"
// @Security Bearer
// @Success 200 {object} responsex.Response{data=model.Ticket}
// @Failure 400 {object} responsex.Response
// @Failure 500 {object} responsex.Response
// @Router /v1/tickets [post]
func (i *impl) CreateTicket(c *gin.Context) {
	// todo: 2024/6/6|sean|add some logic here
	_ = c.Error(errors.New("not implemented"))
}

// ListTicketQuery defines the list query.
type ListTicketQuery struct {
	Page int `form:"page" default:"1" minimum:"1"`
	Size int `form:"size" default:"10" minimum:"1" maximum:"100"`
}

// ListTicket is used to list tickets.
// @Summary List tickets
// @Description list tickets
// @Tags tickets
// @Accept json
// @Produce json
// @Param query query ListTicketQuery false "query string"
// @Security Bearer
// @Success 200 {object} responsex.Response{data=[]model.Ticket}
// @Failure 400 {object} responsex.Response
// @Failure 500 {object} responsex.Response
// @Header 200 {int} X-Total-Count "total count"
// @Header 200 {int} X-Page "page"
// @Header 200 {int} X-Page-Size "page size"
// @Router /v1/tickets [get]
func (i *impl) ListTicket(c *gin.Context) {
	// todo: 2024/6/6|sean|add some logic here
	_ = c.Error(errors.New("not implemented"))
}

// GetTicket is used to get a ticket.
// @Summary Get a ticket
// @Description get a ticket
// @Tags tickets
// @Accept json
// @Produce json
// @Param id path string true "ticket id"
// @Security Bearer
// @Success 200 {object} responsex.Response{data=model.Ticket}
// @Failure 400 {object} responsex.Response
// @Failure 404 {object} responsex.Response
// @Failure 500 {object} responsex.Response
// @Router /v1/tickets/{id} [get]
func (i *impl) GetTicket(c *gin.Context) {
	// todo: 2024/6/6|sean|add some logic here
	_ = c.Error(errors.New("not implemented"))
}

// UpdateTicketPayload defines the update payload.
type UpdateTicketPayload struct {
	*model.Ticket `json:",inline"`
}

// UpdateTicket is used to update a ticket.
// @Summary Update a ticket
// @Description update a ticket
// @Tags tickets
// @Accept json
// @Produce json
// @Param id path string true "ticket id"
// @Param payload body UpdateTicketPayload true "update ticket payload"
// @Security Bearer
// @Success 200 {object} responsex.Response{data=model.Ticket}
// @Failure 400 {object} responsex.Response
// @Failure 404 {object} responsex.Response
// @Failure 500 {object} responsex.Response
// @Router /v1/tickets/{id} [put]
func (i *impl) UpdateTicket(c *gin.Context) {
	// todo: 2024/6/6|sean|add some logic here
	_ = c.Error(errors.New("not implemented"))
}

// DeleteTicket is used to delete a ticket.
// @Summary Delete a ticket
// @Description delete a ticket
// @Tags tickets
// @Accept json
// @Produce json
// @Param id path string true "ticket id"
// @Security Bearer
// @Success 200 {object} responsex.Response
// @Failure 400 {object} responsex.Response
// @Failure 404 {object} responsex.Response
// @Failure 500 {object} responsex.Response
// @Router /v1/tickets/{id} [delete]
func (i *impl) DeleteTicket(c *gin.Context) {
	// todo: 2024/6/6|sean|add some logic here
	_ = c.Error(errors.New("not implemented"))
}
