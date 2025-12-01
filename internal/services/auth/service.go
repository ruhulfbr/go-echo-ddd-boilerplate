package auth

import (
	"context"
	"errors"
	"fmt"

	errors2 "github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/common/errors"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/domain/user"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/http/requests"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/http/responses"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/services/token"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	userService  user.Service
	tokenService token.TokenServiceInterface
}

func NewService(userService user.Service, tokenService token.TokenServiceInterface) *Service {
	return &Service{
		userService:  userService,
		tokenService: tokenService,
	}
}

func (s *Service) GenerateToken(ctx context.Context, request *requests.LoginRequest) (*responses.LoginResponse, error) {
	user, err := s.userService.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return nil, fmt.Errorf("get user by email: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return nil, errors.Join(fmt.Errorf("compare hash and passowrd: %w", err), errors2.ErrInvalidPassword)
	}

	accessToken, exp, err := s.tokenService.CreateAccessToken(ctx, &user)
	if err != nil {
		return nil, fmt.Errorf("create access token: %w", err)
	}

	refreshToken, err := s.tokenService.CreateRefreshToken(ctx, &user)
	if err != nil {
		return nil, fmt.Errorf("create refresh token: %w", err)
	}

	response := responses.NewLoginResponse(accessToken, refreshToken, exp)

	return response, nil
}

func (s *Service) RefreshToken(ctx context.Context, request *requests.RefreshRequest) (*responses.LoginResponse, error) {
	claims, err := s.tokenService.ParseRefreshToken(ctx, request.Token)
	if err != nil {
		return nil, errors.Join(fmt.Errorf("parse token: %w", err), errors2.ErrInvalidAuthToken)
	}

	user, err := s.userService.GetByID(ctx, claims.ID)
	if err != nil {
		return nil, fmt.Errorf("get user by email: %w", err)
	}

	accessToken, exp, err := s.tokenService.CreateAccessToken(ctx, &user)
	if err != nil {
		return nil, fmt.Errorf("create access token: %w", err)
	}

	refreshToken, err := s.tokenService.CreateRefreshToken(ctx, &user)
	if err != nil {
		return nil, fmt.Errorf("create refresh token: %w", err)
	}

	response := responses.NewLoginResponse(accessToken, refreshToken, exp)

	return response, nil
}
