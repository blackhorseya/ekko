//go:generate mockgen -destination=./mock_${GOFILE} -package=repo -source=${GOFILE}

package repo

import (
	"github.com/blackhorseya/ekko/pkg/contextx"
	im "github.com/blackhorseya/ekko/pkg/entity/domain/issue/model"
	"github.com/google/wire"
)

// QueryTicketsCondition declare list tasks condition
type QueryTicketsCondition struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

// IRepo declare issue repo interface
type IRepo interface {
	// GetByID serve caller to get a issue by id
	GetByID(ctx contextx.Contextx, id int64) (info *im.Ticket, err error)

	// List serve caller to list all tasks
	List(ctx contextx.Contextx, condition QueryTicketsCondition) (info []*im.Ticket, err error)

	// Create serve caller to create a issue with title
	Create(ctx contextx.Contextx, created *im.Ticket) (info *im.Ticket, err error)

	// Count serve caller to count all tasks
	Count(ctx contextx.Contextx, condition QueryTicketsCondition) (total int, err error)

	// Update serve caller to update a issue
	Update(ctx contextx.Contextx, updated *im.Ticket) error

	// DeleteByID serve caller to remove a issue by id
	DeleteByID(ctx contextx.Contextx, id int64) error
}

// ProvideMariadb is a provider set for mariadb implementation
var ProvideMariadb = wire.NewSet(NewMariadb)

// HTTPClientSet is a provider set for http client implementation
var HTTPClientSet = wire.NewSet(NewHTTPClient)
