package repository

import (
	"context"

	"github.com/sirArthurDayne/rest-ws/models"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, id string) (*models.User, error)
	InsertPost(ctx context.Context, post *models.Post) error
	GetPostById(ctx context.Context, id string) (*models.Post, error)
	UpdatePost(ctx context.Context, post *models.Post) error
	DeletePost(ctx context.Context, id, userId string) error
    ListPosts(ctx context.Context, page uint64) ([]*models.Post, error)
	Close() error
}

var implementation UserRepository

func SetRepository(repository UserRepository) {
	implementation = repository
}

// USER CRUD
func InsertUser(ctx context.Context, user *models.User) error {
	return implementation.InsertUser(ctx, user)
}

func GetUserById(ctx context.Context, id string) (*models.User, error) {
	return implementation.GetUserById(ctx, id)
}

func GetUserByEmail(ctx context.Context, id string) (*models.User, error) {
	return implementation.GetUserByEmail(ctx, id)
}

// POST CRUD

func InsertPost(ctx context.Context, post *models.Post) error {
	return implementation.InsertPost(ctx, post)
}

func GetPostById(ctx context.Context, id string) (*models.Post, error) {
	return implementation.GetPostById(ctx, id)
}

// TODO: make full CRUD functionality for both user and Post
func UpdatePost(ctx context.Context, post *models.Post) error {
	return implementation.UpdatePost(ctx, post)
}

func DeletePost(ctx context.Context, id, userId string) error {
	return implementation.DeletePost(ctx, id, userId)
}

func ListPosts(ctx context.Context, page uint64) ([]*models.Post, error) {
    return implementation.ListPosts(ctx, page)
}

func Close() error {
	return implementation.Close()
}
