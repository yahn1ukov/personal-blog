package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Blog struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Content     string             `bson:"content"`
	Tags        []string           `bson:"tags,omitempty"`
	PublishedAt time.Time          `bson:"publishedAt"`
	UpdatedAt   time.Time          `bson:"updatedAt"`
}
