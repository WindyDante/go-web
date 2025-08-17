package server

import (
	"errors"
	"fmt"
	"net/http"
	"oa/cmd/di"
	"oa/internal/config"
	"oa/internal/router"

	"github.com/gin-gonic/gin"
)

// Server 服务器结构体,包含Gin引擎
type Server struct {
	GinEngine  *gin.Engine
	httpServer *http.Server
}

// New 创建一个新的服务器实例
func New() *Server {
	return &Server{}
}

// Init 初始化服务器
func (s *Server) Init() error {
	// 设置Gin Mode
	if config.AppConfig.Server.Mode == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	s.GinEngine = gin.New()

	handlers, err := di.BuildHandlers()
	if err != nil {
		return err
	}

	router.Route(s.GinEngine, handlers)

	return nil
}

func (s *Server) Start() {
	port := config.AppConfig.Server.Port
	s.httpServer = &http.Server{
		Addr:    ":" + port,
		Handler: s.GinEngine,
	}

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic("Failed to start server: " + err.Error())
		}
	}()

	fmt.Println("Server is running on port " + port)
}
