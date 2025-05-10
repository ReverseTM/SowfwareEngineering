package app

import (
	"context"
	"software-engineering-2/internal/infrastructure/http"
	"time"
)

type App struct {
	server *http.Server
}

func NewApp() *App {
	dependencyProvider := newProvider()

	return &App{
		server: dependencyProvider.Server(),
	}
}

func (a *App) Run() error {
	return a.server.Start()
}

func (a *App) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return a.server.Stop(ctx)
}
