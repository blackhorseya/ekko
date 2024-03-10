package mongodb

import (
	"time"

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
