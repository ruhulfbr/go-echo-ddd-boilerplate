package post

import (
	"context"
	"fmt"

	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/domain/post"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/http/requests"
)

type postRepository interface {
	Create(ctx context.Context, post *post.Post) error
	GetPosts(ctx context.Context) ([]post.Post, error)
	GetPost(ctx context.Context, id uint) (post.Post, error)
	Update(ctx context.Context, post *post.Post) error
	Delete(ctx context.Context, post *post.Post) error
}

type Service struct {
	postRepository postRepository
}

func NewService(postRepository postRepository) *Service {
	return &Service{postRepository: postRepository}
}

func (s *Service) Create(ctx context.Context, post *post.Post) error {
	if err := s.postRepository.Create(ctx, post); err != nil {
		return fmt.Errorf("create post in repository: %w", err)
	}

	return nil
}

func (s *Service) GetPosts(ctx context.Context) ([]post.Post, error) {
	posts, err := s.postRepository.GetPosts(ctx)
	if err != nil {
		return nil, fmt.Errorf("get posts from repository: %w", err)
	}

	return posts, nil
}

func (s *Service) GetPost(ctx context.Context, id uint) (post.Post, error) {
	post, err := s.postRepository.GetPost(ctx, id)
	if err != nil {
		return post.Post{}, fmt.Errorf("get post from repository: %w", err)
	}

	return post, nil
}

func (s *Service) Update(ctx context.Context, post *post.Post, updatePostRequest requests.UpdatePostRequest) error {
	post.Content = updatePostRequest.Content
	post.Title = updatePostRequest.Title

	if err := s.postRepository.Update(ctx, post); err != nil {
		return fmt.Errorf("update post in repository: %w", err)
	}

	return nil
}

func (s *Service) Delete(ctx context.Context, post *post.Post) error {
	if err := s.postRepository.Delete(ctx, post); err != nil {
		return fmt.Errorf("delete post in repository: %w", err)
	}

	return nil
}
