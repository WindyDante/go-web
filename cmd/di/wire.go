//go:build wireinject
// +build wireinject

package di

import (
	"oa/internal/database"
	userHandler "oa/internal/handler/user"
	userRepository "oa/internal/repository/user"
	userService "oa/internal/service/user"

	"github.com/google/wire"
)

func BuildHandlers() (*Handlers, error) {
	wire.Build(
		DataBaseSet,
		UserSet,
		NewHandlers,
	)
	return &Handlers{}, nil
}

var DataBaseSet = wire.NewSet(
	database.NewDatabase,
	database.NewRedis,
	database.NewData,
)

var UserSet = wire.NewSet(
	userRepository.NewUserRepository,
	userService.NewUserService,
	userHandler.NewUserHandler,
)
