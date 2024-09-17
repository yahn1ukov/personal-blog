package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Blog struct {
	ID          bson.ObjectID `bson:"_id"`
	Title       string        `bson:"title"`
	Content     string        `bson:"content"`
	Tags        []string      `bson:"tags,omitempty"`
	PublishedAt time.Time     `bson:"publishedAt"`
	UpdatedAt   time.Time     `bson:"updatedAt"`
}
