package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/config"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/http/handlers"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/http/middleware"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/infrastructure/logger/slogx"
)

type Handlers struct {
	PostHandler     *handlers.PostHandlers
	AuthHandler     *handlers.AuthHandler
	OAuthHandler    *handlers.OAuthHandler
	RegisterHandler *handlers.RegisterHandler
}

func ConfigureRoutes(tracer *slogx.TraceStarter, engine *echo.Echo, cfg *config.Config, handlers Handlers) error {
	engine.HTTPErrorHandler = middleware.EchoHTTPErrorHandler
	engine.Use(middleware.NewRequestLogger(tracer))

	engine.POST("/login", handlers.AuthHandler.Login)
	engine.POST("/register", handlers.RegisterHandler.Register)
	engine.POST("/google-oauth", handlers.OAuthHandler.GoogleOAuth)
	engine.POST("/refresh", handlers.AuthHandler.RefreshToken)

	r := engine.Group("", middleware.NewRequestDebugger())

	r.Use(middleware.EchoJWTMiddleware(cfg.Auth.AccessSecret))

	r.GET("/posts", handlers.PostHandler.GetPosts)
	r.POST("/posts", handlers.PostHandler.CreatePost)
	r.DELETE("/posts/:id", handlers.PostHandler.DeletePost)
	r.PUT("/posts/:id", handlers.PostHandler.UpdatePost)

	return nil
}
