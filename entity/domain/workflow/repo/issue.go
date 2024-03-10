//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package repo

import (
	"github.com/blackhorseya/ekko/entity/domain/workflow/agg"
	"github.com/blackhorseya/ekko/pkg/contextx"
)

// IIssueRepo is the interface that represents the issue repository.
type IIssueRepo interface {
	// GetByID returns the issue by the given ID.
	GetByID(ctx contextx.Contextx, id string) (item *agg.Issue, err error)

	// Create creates a new issue.
	Create(ctx contextx.Contextx, item *agg.Issue) (err error)

	// Update updates the issue.
	Update(ctx contextx.Contextx, item *agg.Issue) (err error)

	// Delete deletes the issue by the given ID.
	Delete(ctx contextx.Contextx, id string) (err error)
}
