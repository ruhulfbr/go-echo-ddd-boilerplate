package bootstrap

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/config"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/http"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/http/routes"
	"github.com/ruhulfbr/go-echo-ddd-boilerplate/internal/infrastructure/logger/slogx"
	"gorm.io/gorm"
)

const shutdownTimeout = 20 * time.Second

func InitServer(cfg *config.Config, gormDB *gorm.DB) (*http.Server, error) {
	traceStarter := slogx.NewTraceStarter(uuid.NewV7)

	services, _ := InitServices(cfg, gormDB)
	handlers := InitHandlers(services)
	engine := echo.New()

	if err := routes.ConfigureRoutes(traceStarter, engine, cfg, routes.Handlers{
		PostHandler:     handlers.Post,
		AuthHandler:     handlers.Auth,
		OAuthHandler:    handlers.OAuth,
		RegisterHandler: handlers.Register,
	}); err != nil {
		return nil, err
	}

	return http.NewServer(engine), nil
}

func StartServer(app *http.Server, cfg *config.Config) error {
	go func() {
		_ = app.Start(cfg.HTTP.Port)
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	return app.Shutdown(ctx)
}
