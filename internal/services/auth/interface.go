package auth

import (
	"context"

	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/http/requests"
	apiResponses "github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/http/responses"
)

type AuthServiceInterface interface {
	GenerateToken(ctx context.Context, request *requests.LoginRequest) (*apiResponses.LoginResponse, error)
	RefreshToken(ctx context.Context, request *requests.RefreshRequest) (*apiResponses.LoginResponse, error)
}
