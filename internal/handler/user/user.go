package user

import (
	res "oa/internal/response"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
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
