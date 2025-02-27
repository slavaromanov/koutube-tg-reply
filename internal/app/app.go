package app

import (
	"context"

	"go.uber.org/zap"
	"koutube-tg-reply/internal/tg"
)

type App struct {
	bot *tg.Bot
}

func newLogger() (*zap.Logger, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}
	zap.ReplaceGlobals(logger)
	return logger, nil
}

func newApp(bot *tg.Bot) (*App, error) {
	return &App{
		bot: bot,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	return a.bot.Run(ctx)
}
