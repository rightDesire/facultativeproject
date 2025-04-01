//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/rightDesire/facultativeproject/api/internal/configs"
)

func InitApp() *App {
	wire.Build(
		configs.NewConfigsFromEnv,

		NewApp,
	)
	return &App{}
}
