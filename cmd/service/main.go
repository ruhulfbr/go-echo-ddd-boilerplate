package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/common/config"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/http"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/http/handlers"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/http/routes"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/infrastructure/database"
	slogx2 "github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/infrastructure/logger/slogx"
	post2 "github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/infrastructure/repositories/post"
	repositories2 "github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/infrastructure/repositories/user"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/services/auth"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/services/oauth"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/services/post"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/services/token"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/services/user"
)

const shutdownTimeout = 20 * time.Second

func main() {
	if err := run(); err != nil {
		slog.Error("Service run error", "err", err.Error())
		os.Exit(1)
	}
}

func run() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("load env file: %w", err)
	}

	var cfg config.Config
	if err := env.Parse(&cfg); err != nil {
		return fmt.Errorf("parse env: %w", err)
	}

	if err := slogx2.Init(cfg.Logger); err != nil {
		return fmt.Errorf("init logger: %w", err)
	}

	traceStarter := slogx2.NewTraceStarter(uuid.NewV7)

	gormDB, err := database.NewGormDB(cfg.DB)
	if err != nil {
		return fmt.Errorf("new conn connection: %w", err)
	}

	userRepository := repositories2.NewUserRepository(gormDB)
	userService := user.NewService(userRepository)

	postRepository := post2.NewPostRepository(gormDB)
	postService := post.NewService(postRepository)

	provider, err := oidc.NewProvider(context.Background(), "https://accounts.google.com")
	if err != nil {
		return fmt.Errorf("oidc.NewProvider: %w", err)
	}

	verifier := provider.Verifier(&oidc.Config{ClientID: cfg.OAuth.ClientID})

	tokenService := token.NewService(
		time.Now,
		cfg.Auth.AccessTokenDuration,
		cfg.Auth.RefreshTokenDuration,
		[]byte(cfg.Auth.AccessSecret),
		[]byte(cfg.Auth.RefreshSecret),
	)

	authService := auth.NewService(userService, tokenService)
	oAuthService := oauth.NewService(verifier, tokenService, userService)

	postHandler := handlers.NewPostHandlers(postService)
	authHandler := handlers.NewAuthHandler(authService)
	oAuthHandler := handlers.NewOAuthHandler(oAuthService)
	registerHandler := handlers.NewRegisterHandler(userService)

	// Configure middleware with the custom claims type
	echoJWTConfig := echojwt.Config{
		NewClaimsFunc: func(echo.Context) jwt.Claims {
			return new(token.JwtCustomClaims)
		},
		SigningKey: []byte(cfg.Auth.AccessSecret),
	}

	echoJWTMiddleware := echojwt.WithConfig(echoJWTConfig)

	engine := echo.New()
	err = routes.ConfigureRoutes(traceStarter, engine, routes.Handlers{
		PostHandler:       postHandler,
		AuthHandler:       authHandler,
		OAuthHandler:      oAuthHandler,
		RegisterHandler:   registerHandler,
		EchoJWTMiddleware: echoJWTMiddleware,
	})
	if err != nil {
		return fmt.Errorf("configure routes: %w", err)
	}

	app := http.NewServer(engine)
	go func() {
		if err = app.Start(cfg.HTTP.Port); err != nil {
			slog.Error("Server error", "err", err.Error())
		}
	}()

	shutdownChannel := make(chan os.Signal, 1)
	signal.Notify(shutdownChannel, os.Interrupt, syscall.SIGHUP, syscall.SIGTERM)
	<-shutdownChannel

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := app.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("http http shutdown: %w", err)
	}

	dbConnection, err := gormDB.DB()
	if err != nil {
		return fmt.Errorf("get conn connection: %w", err)
	}

	if err := dbConnection.Close(); err != nil {
		return fmt.Errorf("close conn connection: %w", err)
	}

	return nil
}
