package routes

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/config"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/http/handlers"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/http/middleware"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/infrastructure/logger/slogx"
	"gorm.io/gorm"
)

func ConfigureRoutes(engine *echo.Echo, cfg *config.Config, gormDB *gorm.DB) error {
	handlers := handlers.InitHandlers(cfg, gormDB)

	tracer := slogx.NewTraceStarter(uuid.NewV7)

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
