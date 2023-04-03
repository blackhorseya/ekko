package testdata

import (
	tm "github.com/blackhorseya/ekko/pkg/entity/domain/task/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	// Ticket1 task 1
	Ticket1 = &tm.Ticket{
		Id:        1,
		Title:     "task 1",
		Status:    tm.TicketStatus_TICKET_STATUS_TODO,
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
	}
)
