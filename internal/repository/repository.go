package repository

import (
	"context"
	"errors"

	"github.com/yahn1ukov/personal-blog/internal/database"
	"github.com/yahn1ukov/personal-blog/internal/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Repository interface {
	Create(context.Context, *model.Blog) error
	GetAll(context.Context) ([]*model.Blog, error)
	GetByID(context.Context, bson.ObjectID) (*model.Blog, error)
	Update(context.Context, bson.ObjectID, map[string]interface{}) error
	Delete(context.Context, bson.ObjectID) error
}

type repository struct {
	collection *mongo.Collection
}

var _ Repository = (*repository)(nil)

func New(database *database.Database) *repository {
	collection := database.Database("personal-blog").Collection("blogs")

	return &repository{
		collection: collection,
	}
}

func (r *repository) Create(ctx context.Context, blog *model.Blog) error {
	if _, err := r.collection.InsertOne(ctx, blog); err != nil {
		return err
	}

	return nil
}

func (r *repository) GetAll(ctx context.Context) ([]*model.Blog, error) {
	options := options.Find().SetSort(bson.D{{Key: "publishedAt", Value: -1}})

	cursor, err := r.collection.Find(ctx, bson.M{}, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var blogs []*model.Blog
	if err = cursor.All(ctx, &blogs); err != nil {
		return nil, err
	}

	return blogs, nil
}

func (r *repository) GetByID(ctx context.Context, objectID bson.ObjectID) (*model.Blog, error) {
	var blog model.Blog
	if err := r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&blog); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return &blog, nil
}

func (r *repository) Update(ctx context.Context, objectID bson.ObjectID, updatedFields map[string]interface{}) error {
	options := options.Update().SetUpsert(false)

	result, err := r.collection.UpdateByID(ctx, objectID, bson.M{"$set": updatedFields}, options)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return ErrNotFound
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, objectID bson.ObjectID) error {
	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return ErrNotFound
	}

	return nil
}
