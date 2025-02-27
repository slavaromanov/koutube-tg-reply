package app

import (
	"github.com/caarlos0/env/v11"
	"koutube-tg-reply/internal/tg"
)

type Config struct {
	Token tg.Token `env:"TG_TOKEN"`
}

func NewConfig() Config {
	return env.Must(env.ParseAs[Config]())
}
