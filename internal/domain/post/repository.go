package post

import (
	"context"

	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/infrastructure/models"
)

type Repository interface {
	Create(ctx context.Context, post *models.Post) error
	GetPosts(ctx context.Context) ([]models.Post, error)
	GetPost(ctx context.Context, id uint) (models.Post, error)
	Update(ctx context.Context, post *models.Post) error
	Delete(ctx context.Context, post *models.Post) error
}
