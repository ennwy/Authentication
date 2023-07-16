package main

import (
	"context"
	"os/signal"
	"syscall"
	"time"

	"github.com/ennwy/auth/internal/app"
	"github.com/ennwy/auth/internal/logger"
	"github.com/ennwy/auth/internal/server"
)

var l app.Logger

func main() {
	config := NewConfig()

	l = logger.New(config.Logger.Level, config.Logger.OutputPath)
	l.Debug("CONFIG:", config)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	var s app.Server
	var err error
	if s, err = server.New(ctx, l, config.HTTP); err != nil {
		l.Error("New server:", err)
		return
	}

	go func() {
		<-ctx.Done()
		l.Info("Stop: ctx canceled")

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		if err := s.Stop(ctx); err != nil {
			l.Error("Failed to stop server:", err)
		}

		l.Debug("Server has stopped")
	}()

	l.Info("Server started on PORT =", config.HTTP.Port)
	if err := s.Start(); err != nil {
		l.Error("Server error", err)
		cancel()
	}
}
