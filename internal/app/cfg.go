package app

import (
	"github.com/caarlos0/env/v11"
	"koutube-tg-reply/internal/http"
	koutube_conv "koutube-tg-reply/internal/koutube-conv"
	"koutube-tg-reply/internal/tg"
)

type Config struct {
	Token   tg.Token             `env:"TG_TOKEN"`
	HTTPort http.HTTPPort        `env:"HTTP_PORT" envDefault:"8080"`
	BaseURL koutube_conv.BaseURL `env:"BASE_URL" envDefault:"https://koutu.be/shorts"`
}

func NewConfig() Config {
	return env.Must(env.ParseAs[Config]())
}
