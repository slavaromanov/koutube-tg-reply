//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	koutube_conv "koutube-tg-reply/internal/koutube-conv"
	"koutube-tg-reply/internal/proxy"
	"koutube-tg-reply/internal/tg"
)

//go:generate go run github.com/google/wire/cmd/wire@v0.6.0

func New() (*App, error) {
	wire.Build(wire.NewSet(
		koutube_conv.NewConverter,
		tg.New,
		NewConfig,
		newApp,
		newLogger,
		proxy.NewServer,
		wire.FieldsOf(new(Config), "Token", "HTTPort"),
		wire.Bind(new(tg.Converter), new(*koutube_conv.Converter)),
		wire.Bind(new(proxy.VideoIDExtractor), new(*koutube_conv.Converter)),
	))
	return &App{}, nil
}
