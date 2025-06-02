//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"koutube-tg-reply/internal/http"
	koutube_conv "koutube-tg-reply/internal/koutube-conv"
	"koutube-tg-reply/internal/og"
	"koutube-tg-reply/internal/page"
	"koutube-tg-reply/internal/tg"
	"koutube-tg-reply/internal/ytdl"
)

//go:generate go run github.com/google/wire/cmd/wire@v0.6.0

func New() (*App, error) {
	wire.Build(wire.NewSet(
		koutube_conv.NewConverter,
		tg.New,
		ytdl.NewYoutubeDL,
		http.NewServer,
		og.NewBuilder,
		page.NewService,
		NewConfig,
		newApp,
		newLogger,
		wire.FieldsOf(new(Config), "Token", "HTTPort", "BaseURL"),
		wire.Bind(new(tg.Converter), new(*koutube_conv.Converter)),
		wire.Bind(new(http.PageBuilder), new(*page.Service)),
		wire.Bind(new(page.VideoStreamURLExtractor), new(*ytdl.YoutubeDL)),
		wire.Bind(new(page.Builder), new(*og.Builder)),
	))
	return &App{}, nil
}
