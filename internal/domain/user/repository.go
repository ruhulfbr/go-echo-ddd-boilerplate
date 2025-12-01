package user

import (
	"context"

	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/infrastructure/models"
)

type Repository interface {
	Create(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, id uint) (models.User, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	CreateUserAndOAuthProvider(ctx context.Context, user *models.User, oauthProvider *models.OAuthProviders) error
}
