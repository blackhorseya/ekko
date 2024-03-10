package agg

import (
	"github.com/blackhorseya/ekko/entity/domain/workflow/model"
)

// Issue is an aggregate root that represents an issue.
type Issue struct {
	*model.Ticket
}
