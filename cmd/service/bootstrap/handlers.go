package bootstrap

import (
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/http/handlers"
)

type Handlers struct {
	Post          *handlers.PostHandlers
	Auth          *handlers.AuthHandler
	OAuth         *handlers.OAuthHandler
	Register      *handlers.RegisterHandler
	UrlDownloader *handlers.UrlDownloaderHandler
}

func InitHandlers(services *Services) *Handlers {
	return &Handlers{
		Post:          handlers.NewPostHandlers(services.PostService),
		Auth:          handlers.NewAuthHandler(services.AuthService),
		OAuth:         handlers.NewOAuthHandler(services.OAuthService),
		Register:      handlers.NewRegisterHandler(services.UserService),
		UrlDownloader: handlers.NewUrlDownloaderHandler(),
	}
}
