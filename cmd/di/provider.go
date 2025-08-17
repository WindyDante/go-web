package di

import (
	"oa/internal/database"
	"oa/internal/handler/user"
)

type Handlers struct {
	UserHandler *user.UserHandler
	Database    *database.Data
}

func NewHandlers(userHandler *user.UserHandler, database *database.Data) *Handlers {
	return &Handlers{
		UserHandler: userHandler,
		Database:    database,
	}
}
