package user

import "oa/internal/dal/model"

type UserRepositoryInterface interface {
	GetUserByUsername(username string) (model.User, error)
}
