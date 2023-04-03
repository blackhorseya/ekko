package testdata

import (
	im "github.com/blackhorseya/ekko/pkg/entity/domain/issue/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	// Ticket1 issue 1
	Ticket1 = &im.Ticket{
		Id:        1,
		Title:     "issue 1",
		Status:    im.TicketStatus_TICKET_STATUS_TODO,
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
	}
)
