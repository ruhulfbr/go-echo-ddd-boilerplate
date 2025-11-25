package user

import (
	"context"

	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/http/requests"
)

type Service interface {
	GetUserByEmail(ctx context.Context, email string) (User, error)
	Register(ctx context.Context, request *requests.RegisterRequest) error
}
