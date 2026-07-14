package main

import (
	"context"
	"go-timestampbc/internal/api/handler"
	"go-timestampbc/internal/config"
	"go-timestampbc/internal/logger"
	"go-timestampbc/internal/store/sqlite"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.MustLoad()

	logger := logger.New(cfg.LogLevel)
	slog.SetDefault(logger)

	logger.Info("starting application",
		"env", cfg.AppEnv,
		"log_level", cfg.LogLevel,
		"http_port", cfg.HTTPPort,
	)

	startupCtx, startupCancel := context.WithTimeout(
		context.Background(),
		time.Duration(cfg.StartupTimeout)*time.Second,
	)
	defer startupCancel()

	sqliteClient, err := sqlite.NewClient(startupCtx, cfg.SQLitePath)
	if err != nil {
		logger.Error("failed to init sqlite", "error", err)
		os.Exit(1)
	}
	defer sqliteClient.Close()
	logger.Info("sqlite initialized", "path", cfg.SQLitePath)

	sqliteStore := sqlite.NewStore(sqliteClient.DB())
	r := handler.NewRouter(sqliteStore, logger)

	addr := net.JoinHostPort(cfg.HTTPHost, cfg.HTTPPort)
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("server error", "error", err)
			os.Exit(1)
		}
	}()

	<-ctx.Done()
	logger.Info("shutting down server")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Error("server shutdown error", "error", err)
	}
}
