package router

import (
	"oa/cmd/di"
	"oa/internal/middleware"

	"github.com/gin-gonic/gin"
)

type AppRouterGroup struct {
	PublicGroup *gin.RouterGroup
	AuthGroup   *gin.RouterGroup
}

func Route(r *gin.Engine, h *di.Handlers) {
	// 初始化路由

	// 中间件
	setupMiddleware(r)

	// 路由组
	appRouterGroup := setupRouterGroup(r)

	// 用户路由
	setupUserGroup(appRouterGroup, h)

}

func setupRouterGroup(r *gin.Engine) *AppRouterGroup {
	public := r.Group("/api")
	auth := r.Group("/api")
	auth.Use(middleware.AuthJWT())
	return &AppRouterGroup{
		PublicGroup: public,
		AuthGroup:   auth,
	}
}
