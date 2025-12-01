package user

import (
	"context"
	"fmt"

	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/domain/user"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/http/requests"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/infrastructure/models"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	userRepository user.Repository
}

func NewService(userRepository user.Repository) *Service {
	return &Service{userRepository: userRepository}
}

func (s *Service) Register(ctx context.Context, request *requests.RegisterRequest) error {
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return fmt.Errorf("encrypt password: %w", err)
	}

	userInfo := &models.User{
		Email:    request.Email,
		Name:     request.Name,
		Password: string(encryptedPassword),
	}

	if err := s.userRepository.Create(ctx, userInfo); err != nil {

		return fmt.Errorf("create user in repository: %w", err)
	}

	return nil
}

func (s *Service) GetByID(ctx context.Context, id uint) (models.User, error) {
	userInfo, err := s.userRepository.GetByID(ctx, id)
	if err != nil {
		return models.User{}, fmt.Errorf("get user by id from repository: %w", err)
	}

	return userInfo, nil
}

func (s *Service) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	userInfo, err := s.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return models.User{}, fmt.Errorf("get user by email from repository: %w", err)
	}

	return userInfo, nil
}

func (s *Service) CreateUserAndOAuthProvider(ctx context.Context, user *models.User, oauthProvider *models.OAuthProviders) error {
	err := s.userRepository.CreateUserAndOAuthProvider(ctx, user, oauthProvider)
	if err != nil {
		return fmt.Errorf("create user and oauth provider from repository: %w", err)
	}

	return nil
}
