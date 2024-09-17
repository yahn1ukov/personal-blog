package service

import (
	"context"
	"time"

	"github.com/yahn1ukov/personal-blog/internal/dto"
	"github.com/yahn1ukov/personal-blog/internal/model"
	"github.com/yahn1ukov/personal-blog/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Service interface {
	Create(context.Context, *dto.CreateInput) error
	GetAll(context.Context) ([]*dto.GetOutput, error)
	GetByID(context.Context, bson.ObjectID) (*dto.GetOutput, error)
	Update(context.Context, bson.ObjectID, *dto.UpdateInput) error
	Delete(context.Context, bson.ObjectID) error
}

type service struct {
	repository repository.Repository
}

var _ Service = (*service)(nil)

func New(repository repository.Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) Create(ctx context.Context, input *dto.CreateInput) error {
	if input.Title == "" {
		return ErrTitleRequired
	}

	if input.Content == "" {
		return ErrContentRequired
	}

	blog := &model.Blog{
		ID:          bson.ObjectID(primitive.NewObjectID()),
		Title:       input.Title,
		Content:     input.Content,
		Tags:        input.Tags,
		PublishedAt: time.Now(),
		UpdatedAt:   time.Now(),
	}

	return s.repository.Create(ctx, blog)
}

func (s *service) GetAll(ctx context.Context) ([]*dto.GetOutput, error) {
	blogs, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var output []*dto.GetOutput
	for _, blog := range blogs {
		output = append(
			output,
			&dto.GetOutput{
				ID:          blog.ID.Hex(),
				Title:       blog.Title,
				Content:     blog.Content,
				Tags:        blog.Tags,
				PublishedAt: blog.PublishedAt,
				UpdatedAt:   blog.UpdatedAt,
			},
		)
	}

	return output, nil
}

func (s *service) GetByID(ctx context.Context, objectID bson.ObjectID) (*dto.GetOutput, error) {
	blog, err := s.repository.GetByID(ctx, objectID)
	if err != nil {
		return nil, err
	}

	output := &dto.GetOutput{
		ID:          blog.ID.Hex(),
		Title:       blog.Title,
		Content:     blog.Content,
		Tags:        blog.Tags,
		PublishedAt: blog.PublishedAt,
		UpdatedAt:   blog.UpdatedAt,
	}

	return output, nil
}

func (s *service) Update(ctx context.Context, objectID bson.ObjectID, input *dto.UpdateInput) error {
	updatedFields := make(map[string]interface{})

	if input.Title != nil && *input.Title != "" {
		updatedFields["title"] = *input.Title
	}

	if input.Content != nil && *input.Content != "" {
		updatedFields["content"] = *input.Content
	}

	if len(input.Tags) != 0 {
		updatedFields["tags"] = input.Tags
	}

	if len(updatedFields) == 0 {
		return ErrNoFieldsUpdate
	}

	updatedFields["updatedAt"] = time.Now()

	return s.repository.Update(ctx, objectID, updatedFields)
}

func (s *service) Delete(ctx context.Context, objectID bson.ObjectID) error {
	return s.repository.Delete(ctx, objectID)
}
