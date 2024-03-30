package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"github.com/VanLavr/auth/internal/auth/delivery"
	"github.com/VanLavr/auth/internal/auth/repository"
	"github.com/VanLavr/auth/internal/auth/usecase"
	"github.com/VanLavr/auth/internal/pkg/config"
	"github.com/VanLavr/auth/internal/pkg/logging"
)

func main() {
	ctx, close := signal.NotifyContext(context.Background(), os.Interrupt)
	defer close()

	cfg := config.New()
	logger := logging.New()
	logger.SetAsDefault()

	repo := repository.New(cfg)
	repo.Connect(ctx, cfg)

	usecase := usecase.New(repo, cfg)
	srv := delivery.New(usecase, cfg)
	srv.BindRoutes()

	go func() {
		slog.Info("running")
		if err := srv.Run(); err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
	}()

	<-ctx.Done()

	if err := srv.ShutDown(ctx); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	if err := repo.CloseConnetion(context.TODO()); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
