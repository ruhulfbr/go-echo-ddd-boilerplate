package user

import (
	"context"

	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/http/requests"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/infrastructure/models"
)

type Service interface {
	GetByID(ctx context.Context, id uint) (models.User, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	Register(ctx context.Context, request *requests.RegisterRequest) error
	CreateUserAndOAuthProvider(ctx context.Context, user *models.User, oAuthProvider *models.OAuthProviders) error
}
