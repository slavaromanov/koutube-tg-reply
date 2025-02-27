package main

import (
	"context"
	"log/slog"

	"go.uber.org/zap"
	"koutube-tg-reply/internal/app"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	a, err := app.New()
	if err != nil {
		slog.Default().Log(ctx, slog.LevelError, "app.New", err)
		return
	}
	if err := a.Run(ctx); err != nil {
		zap.L().Fatal("app.Run", zap.Error(err))
	}
}
