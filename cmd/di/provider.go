package di

import (
	"oa/internal/database"
	handler "oa/internal/handler/user"
)

type Handlers struct {
	UserHandler *handler.UserHandler
	Database    *database.Data
}

func NewHandlers(userHandler *handler.UserHandler, database *database.Data) *Handlers {
	return &Handlers{
		UserHandler: userHandler,
		Database:    database,
	}
}
