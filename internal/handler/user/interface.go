package user

import "github.com/gin-gonic/gin"

type UserHandlerInterface interface {
	Login() gin.HandlerFunc
}
