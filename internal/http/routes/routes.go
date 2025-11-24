package routes

import (
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/http/handlers"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/http/middleware"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/slogx"

	"github.com/labstack/echo/v4"
)

type Handlers struct {
	PostHandler     *handlers.PostHandlers
	AuthHandler     *handlers.AuthHandler
	OAuthHandler    *handlers.OAuthHandler
	RegisterHandler *handlers.RegisterHandler

	EchoJWTMiddleware echo.MiddlewareFunc
}

func ConfigureRoutes(tracer *slogx.TraceStarter, engine *echo.Echo, handlers Handlers) error {
	engine.Use(middleware.NewRequestLogger(tracer))

	engine.POST("/login", handlers.AuthHandler.Login)
	engine.POST("/register", handlers.RegisterHandler.Register)
	engine.POST("/google-oauth", handlers.OAuthHandler.GoogleOAuth)
	engine.POST("/refresh", handlers.AuthHandler.RefreshToken)

	r := engine.Group("", middleware.NewRequestDebugger())

	r.Use(handlers.EchoJWTMiddleware)

	r.GET("/posts", handlers.PostHandler.GetPosts)
	r.POST("/posts", handlers.PostHandler.CreatePost)
	r.DELETE("/posts/:id", handlers.PostHandler.DeletePost)
	r.PUT("/posts/:id", handlers.PostHandler.UpdatePost)

	return nil
}
