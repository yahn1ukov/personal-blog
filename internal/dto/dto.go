package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateInput struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags,omitempty"`
}

type GetOutput struct {
	ID          primitive.ObjectID `json:"id"`
	Title       string             `json:"title"`
	Content     string             `json:"content"`
	Tags        []string           `json:"tags,omitempty"`
	PublishedAt time.Time          `json:"published_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

type UpdateInput struct {
	Title   *string  `json:"title,omitempty"`
	Content *string  `json:"content,omitempty"`
	Tags    []string `json:"tags,omitempty"`
}
