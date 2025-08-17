package router

import (
	"oa/cmd/di"
)

func setupUserGroup(group *AppRouterGroup, h *di.Handlers) {
	// 设置公共路由
	group.PublicGroup.GET("/user/login", h.UserHandler.Login())
}
