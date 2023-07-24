//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

import (
	issueM "github.com/blackhorseya/ekko/entity/domain/issue/model"
	userM "github.com/blackhorseya/ekko/entity/domain/user/model"
	"github.com/blackhorseya/ekko/pkg/contextx"
)

type ListTasksCondition struct {
	Page uint
	Size uint
}

// IBiz declare issue domain interface
type IBiz interface {
	// GetByID serve caller to given issue's id to get a issue
	GetByID(ctx contextx.Contextx, id int64) (info *issueM.Ticket, err error)

	// List serve caller to list all tasks
	List(ctx contextx.Contextx, page, size int) (info []*issueM.Ticket, total int, err error)

	// Create serve caller to create a issue
	Create(ctx contextx.Contextx, title string) (info *issueM.Ticket, err error)

	// UpdateStatus serve caller to update the issue's status by id
	UpdateStatus(ctx contextx.Contextx, id int64, status issueM.TicketStatus) (info *issueM.Ticket, err error)

	// Delete serve caller to given issue's id to delete the issue
	Delete(ctx contextx.Contextx, id int64) error

	// ListTasks serve caller to given user's profile to list all tasks
	ListTasks(ctx contextx.Contextx, who *userM.Profile, condition ListTasksCondition) (tickets []*issueM.Ticket, total int, err error)
}
