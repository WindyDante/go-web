package router

import (
	"oa/internal/config"
	"oa/internal/middleware"

	"github.com/gin-gonic/gin"
)

func setupMiddleware(r *gin.Engine) {
	// Cors
	r.Use(middleware.Cors())
	if config.AppConfig.Server.Mode == "debug" {
		r.Use(middleware.ZapLogger())
	}
}
