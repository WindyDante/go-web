//go:build wireinject
// +build wireinject

package di

import (
	"oa/internal/database"
	"oa/internal/handler/user"

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
	// user.NewUserRepository,
	// user.NewUserService,
	user.NewUserHandler,
)
