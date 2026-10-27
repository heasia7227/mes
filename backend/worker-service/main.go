package main

import (
	_ "common/db" // 初始化数据库连接
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	start()

	// 信号监听, 常驻阻塞
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// 阻塞等待退出信号
	<-sigChan

	stop()
}

func start() {
	// 启动服务
	log.Println("Worker Service 启动")
}

func stop() {
	// 停止服务
	log.Println("Worker Service 停止")
}
