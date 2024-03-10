package mongodb

import (
	"time"

	"github.com/blackhorseya/ekko/entity/domain/workflow/agg"
	"github.com/blackhorseya/ekko/entity/domain/workflow/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type issue struct {
	ID        primitive.ObjectID `bson:"_id"`
	Title     string             `bson:"title"`
	Completed bool               `bson:"completed"`
	OwnerID   string             `bson:"owner_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

// ToAgg is to convert issue to aggregate issue.
func (x *issue) ToAgg() *agg.Issue {
	return &agg.Issue{
		Ticket: &model.Ticket{
			ID:        x.ID.Hex(),
			Title:     x.Title,
			Completed: x.Completed,
			OwnerID:   x.OwnerID,
			CreatedAt: x.CreatedAt,
			UpdatedAt: x.UpdatedAt,
		},
	}
}
