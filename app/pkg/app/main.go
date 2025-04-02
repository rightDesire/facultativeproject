package app

import (
	"github.com/rightDesire/facultativeproject/app/pkg/configs"
	pg "github.com/rightDesire/facultativeproject/common/postgres"
	"github.com/rightDesire/facultativeproject/dbservice/pkg/users"
)

type App struct {
	Options *configs.Configs
	Gdb     *pg.GDB

	UserRepository users.IUserRepository
	UserService    *users.UserServiceServer
}

func NewApp(cfg *configs.Configs, gdb *pg.GDB, userRepo users.IUserRepository, userService *users.UserServiceServer) *App {
	return &App{
		Options:        cfg,
		Gdb:            gdb,
		UserRepository: userRepo,
		UserService:    userService,
	}
}
