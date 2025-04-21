package app

import (
	"context"

	"go.uber.org/zap"
	"koutube-tg-reply/internal/proxy"
	"koutube-tg-reply/internal/tg"
)

type App struct {
	bot         *tg.Bot
	proxyServer *proxy.Server
}

func newLogger() (*zap.Logger, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}
	zap.ReplaceGlobals(logger)
	return logger, nil
}

func newApp(bot *tg.Bot, proxyServer *proxy.Server) (*App, error) {
	return &App{
		bot:         bot,
		proxyServer: proxyServer,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	go a.proxyServer.Run()
	return a.bot.Run(ctx)
}
