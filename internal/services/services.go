package services

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/config"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/infrastructure/repositories"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/services/auth"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/services/oauth"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/services/post"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/services/token"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/services/user"
)

type Services struct {
	UserService  *user.Service
	PostService  *post.Service
	AuthService  *auth.Service
	OAuthService *oauth.Service
}

func InitServices(repos *repositories.Repositories, cfg *config.Config) (*Services, error) {
	tokenService := initTokenService(cfg)
	_, verifier, _ := initOIDC(cfg)

	sv := &Services{
		UserService:  user.NewService(repos.UserRepository),
		PostService:  post.NewService(repos.PostRepository),
		AuthService:  auth.NewService(user.NewService(repos.UserRepository), tokenService),
		OAuthService: oauth.NewService(verifier, tokenService, user.NewService(repos.UserRepository)),
	}

	return sv, nil
}

func initTokenService(cfg *config.Config) *token.Service {
	return token.NewService(
		time.Now,
		cfg.Auth.AccessTokenDuration,
		cfg.Auth.RefreshTokenDuration,
		[]byte(cfg.Auth.AccessSecret),
		[]byte(cfg.Auth.RefreshSecret),
	)
}

func initOIDC(cfg *config.Config) (*oidc.Provider, *oidc.IDTokenVerifier, error) {
	provider, err := oidc.NewProvider(context.Background(), "https://accounts.google.com")
	if err != nil {
		return nil, nil, fmt.Errorf("oidc.NewProvider: %w", err)
	}

	verifier := provider.Verifier(&oidc.Config{
		ClientID: cfg.OAuth.ClientID,
	})

	return provider, verifier, nil
}
