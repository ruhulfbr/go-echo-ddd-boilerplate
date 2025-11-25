package post

import (
	"context"

	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/http/requests"
)

type Service interface {
	Create(ctx context.Context, post *Post) error
	GetPosts(ctx context.Context) ([]Post, error)
	GetPost(ctx context.Context, id uint) (Post, error)
	Update(ctx context.Context, post *Post, updatePostRequest requests.UpdatePostRequest) error
	Delete(ctx context.Context, post *Post) error
}
