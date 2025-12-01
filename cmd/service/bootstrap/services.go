package bootstrap

import (
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/cmd/service/wiring"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/config"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/services/auth"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/services/oauth"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/services/post"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/services/user"
	"gorm.io/gorm"
)

type Services struct {
	UserService  *user.Service
	PostService  *post.Service
	AuthService  *auth.Service
	OAuthService *oauth.Service
}

func InitServices(cfg *config.Config, db *gorm.DB) (*Services, error) {
	repos := wiring.InitRepositories(db)
	tokenService := InitTokenService(cfg)
	_, verifier, _ := InitOIDC(cfg)

	sv := &Services{
		UserService:  user.NewService(repos.UserRepository),
		PostService:  post.NewService(repos.PostRepository),
		AuthService:  auth.NewService(user.NewService(repos.UserRepository), tokenService),
		OAuthService: oauth.NewService(verifier, tokenService, user.NewService(repos.UserRepository)),
	}

	return sv, nil
}
