package app

import (
	"context"

	"github.com/Srgkharkov/qrsvc/internal/httpapi"
)

type App struct {
	httpServer *httpapi.Server
}

func New(ctx context.Context, cfg *Config) (*App, error) {
	router := httpapi.NewRouter()
	srv := httpapi.NewServer(cfg.HTTP.Addr, router)

	return &App{
		httpServer: srv,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	return a.httpServer.Run(ctx)
}

func (a *App) Shutdown(ctx context.Context) error {
	return a.httpServer.Shutdown(ctx)
}
