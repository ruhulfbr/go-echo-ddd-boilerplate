package user

import (
	"context"

	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/infrastructure/models"
)

type Repository interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id uint) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	CreateUserAndOAuthProvider(ctx context.Context, user *User, oauthProvider *models.OAuthProviders) error
}
