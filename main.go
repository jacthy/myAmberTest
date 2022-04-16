package main

import (
	_ "github.com/liaojuntao/infrastruct/repo"
	"github.com/liaojuntao/server"
)

func main() {
	addr := loadConfig()
	s := server.NewServer(addr)
	s.Init()
	println("启动服务，监听地址：", addr)
	if err := s.Run(); err != nil {
		println("启动失败：", err.Error())
	}
}

// loadConfig 这里应该读配置文件或命令行启动时键入，从而获取丰富的配置项，这里demo简化，直接硬编码
func loadConfig() string {
	return "127.0.0.1:8001"
}
