package app

import "github.com/rightDesire/facultativeproject/internal/configs"

type App struct {
	Options *configs.Configs
}

func NewApp(cfg *configs.Configs) *App {
	return &App{
		Options: cfg,
	}
}
