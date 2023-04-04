package testdata

import (
	"time"

	im "github.com/blackhorseya/ekko/pkg/entity/domain/issue/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	// Ticket1 issue 1
	Ticket1 = &im.Ticket{
		Id:        1,
		Title:     "issue 1",
		Status:    im.TicketStatus_TICKET_STATUS_TODO,
		CreatedAt: timestamppb.New(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
		UpdatedAt: timestamppb.New(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
	}
)
