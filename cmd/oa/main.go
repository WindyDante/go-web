package main

import (
	"oa/internal/config"
	"oa/internal/model/common"
	"oa/internal/server"
	errUtil "oa/internal/util/err"
	"os"
	"os/signal"
	"syscall"
)

var s *server.Server // s 是全局的服务器实例

func main() {

	// 加载配置
	if err := config.LoadConfig(); err != nil {
		errUtil.HandlePanicError(&common.ServerErr{
			Msg: common.CONFIG_LOAD_ERR,
		})
	}

	s = server.New()

	err := s.Init()
	if err != nil {
		errUtil.HandlePanicError(&common.ServerErr{
			Msg: common.DATABASE_CONNECT_ERR,
		})
	}

	s.Start()

	// 阻塞主线程，等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit // 阻塞，直到收到信号
}
