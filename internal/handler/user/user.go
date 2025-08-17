package user

import (
	res "oa/internal/response"
	"oa/internal/service/user"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService user.UserServiceInterface
}

func NewUserHandler(userService user.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (userHandler *UserHandler) Login() gin.HandlerFunc {
	return res.Execute(
		func(c *gin.Context) res.Response {
			return res.Response{
				Msg:  "登录成功",
				Data: gin.H{"token": "example_token"},
			}
		})
}
