//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/rightDesire/facultativeproject/app/pkg/app"
	cfg "github.com/rightDesire/facultativeproject/app/pkg/configs"
	pg "github.com/rightDesire/facultativeproject/common/postgres"
	"github.com/rightDesire/facultativeproject/dbservice/pkg/users"
)

func InitializeApp() (*app.App, error) {
	wire.Build(
		cfg.NewConfigsFromEnv,
		pg.ProviderSet,

		users.RepositoryProviderSet,
		users.ServiceProviderSet,

		app.NewApp,
	)
	return &app.App{}, nil
}
