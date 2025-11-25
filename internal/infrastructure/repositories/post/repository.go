package post

import (
	"context"
	"errors"
	"fmt"

	errors2 "github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/common/errors"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/domain/post"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) Create(ctx context.Context, post *post.Post) error {
	if err := r.db.WithContext(ctx).Create(post).Error; err != nil {
		return fmt.Errorf("execute insert post query: %w", err)
	}

	return nil
}

func (r *PostRepository) GetPosts(ctx context.Context) ([]post.Post, error) {
	var posts []post.Post
	if err := r.db.WithContext(ctx).Find(&posts).Error; err != nil {
		return nil, fmt.Errorf("execute select posts query: %w", err)
	}

	return posts, nil
}

func (r *PostRepository) GetPost(ctx context.Context, id uint) (post.Post, error) {
	var post post.Post
	err := r.db.WithContext(ctx).Where("id = ?", id).Take(&post).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return post.Post{}, errors.Join(errors2.ErrPostNotFound, err)
	} else if err != nil {
		return post.Post{}, fmt.Errorf("execute select post by id query: %w", err)
	}

	return post, nil
}

func (r *PostRepository) Update(ctx context.Context, post *post.Post) error {
	if err := r.db.WithContext(ctx).Save(post).Error; err != nil {
		return fmt.Errorf("execute update post query: %w", err)
	}

	return nil
}

func (r *PostRepository) Delete(ctx context.Context, post *post.Post) error {
	if err := r.db.WithContext(ctx).Delete(post).Error; err != nil {
		return fmt.Errorf("execute delete post query: %w", err)
	}

	return nil
}
