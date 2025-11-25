package post

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, post *Post) error
	GetPosts(ctx context.Context) ([]Post, error)
	GetPost(ctx context.Context, id uint) (Post, error)
	Update(ctx context.Context, post *Post) error
	Delete(ctx context.Context, post *Post) error
}
