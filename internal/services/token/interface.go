package token

import (
	"context"

	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/infrastructure/models"
)

type TokenServiceInterface interface {
	ParseRefreshToken(ctx context.Context, token string) (*JwtCustomRefreshClaims, error)
	CreateAccessToken(ctx context.Context, user *models.User) (string, int64, error)
	CreateRefreshToken(ctx context.Context, user *models.User) (string, error)
}
