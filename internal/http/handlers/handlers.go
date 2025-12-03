package handlers

import (
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/config"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/infrastructure/repositories"
	appServices "github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/services"
	"gorm.io/gorm"
)

type Handlers struct {
	PostHandler     *PostHandlers
	AuthHandler     *AuthHandler
	OAuthHandler    *OAuthHandler
	RegisterHandler *RegisterHandler
}

func InitHandlers(cfg *config.Config, gormDB *gorm.DB) *Handlers {

	repos := repositories.InitRepositories(gormDB)
	services, _ := appServices.InitServices(repos, cfg)

	return &Handlers{
		PostHandler:     NewPostHandlers(services.PostService),
		AuthHandler:     NewAuthHandler(services.AuthService),
		OAuthHandler:    NewOAuthHandler(services.OAuthService),
		RegisterHandler: NewRegisterHandler(services.UserService),
	}
}
