package user

import (
	userModel "oa/internal/model/user"
	"oa/internal/repository/user"
)

type UserService struct {
	userRepository user.UserRepositoryInterface
}

func NewUserService(userRepository user.UserRepositoryInterface) UserServiceInterface {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) Login(dto userModel.UserLoginDto) (userModel.UserLoginVo, error) {
	// Implement login logic here
	// For now, returning a dummy response
	return userModel.UserLoginVo{}, nil
}
