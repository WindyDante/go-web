package user

import "oa/internal/repository/user"

type UserService struct {
	userRepository user.UserRepositoryInterface
}

func NewUserService(userRepository user.UserRepositoryInterface) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}
