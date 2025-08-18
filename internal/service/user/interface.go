package user

import "oa/internal/model/user"

type UserServiceInterface interface {
	Login(user.UserLoginDto) (user.UserLoginVo, error)
}
