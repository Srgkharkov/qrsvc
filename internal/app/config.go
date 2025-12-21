package app

import (
	"os"
)

type Config struct {
	HTTP HTTPConfig
}

type HTTPConfig struct {
	Addr string
}

func LoadConfig() (*Config, error) {
	addr := os.Getenv("HTTP_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	return &Config{
		HTTP: HTTPConfig{
			Addr: addr,
		},
	}, nil
}
