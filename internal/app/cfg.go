package app

import (
	"github.com/caarlos0/env/v11"
	"koutube-tg-reply/internal/proxy"
	"koutube-tg-reply/internal/tg"
)

type Config struct {
	Token   tg.Token        `env:"TG_TOKEN"`
	HTTPort proxy.ProxyPort `env:"HTTP_PORT" envDefault:"8080"`
}

func NewConfig() Config {
	return env.Must(env.ParseAs[Config]())
}
