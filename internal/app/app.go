package app

import (
	"context"

	"go.uber.org/zap"
	"koutube-tg-reply/internal/http"
	"koutube-tg-reply/internal/tg"
)

type App struct {
	bot    *tg.Bot
	server *http.Server
}

func newLogger() (*zap.Logger, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}
	zap.ReplaceGlobals(logger)
	return logger, nil
}

func newApp(bot *tg.Bot, proxyServer *http.Server) (*App, error) {
	return &App{
		bot:    bot,
		server: proxyServer,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	go a.server.Start()
	return a.bot.Run(ctx)
}
