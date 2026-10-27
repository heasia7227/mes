package main

import (
	"log"

	_ "common/db" // 初始化数据库连接
	"common/utils/conf"

	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func main() {

	// 启动服务前, 给 Beego 加全局异常处理
	web.BConfig.RecoverPanic = true // 开启 Beego 自带 recover (基础版)
	// Basic config
	web.BConfig.AppName = conf.GetStringConf("appname", "api-service")
	web.BConfig.RunMode = conf.GetStringConf("runmode", "dev")
	port := conf.GetIntConf("httpport", 8080)
	web.BConfig.Listen.HTTPPort = port
	web.BConfig.CopyRequestBody = true

	// // Global filters: logger then auth
	// web.InsertFilter("/*", web.BeforeRouter, func(ctx *context.Context) {
	// 	middleware.RequestLogger(ctx)
	// })
	// web.InsertFilter("/*", web.BeforeExec, func(ctx *context.Context) {
	// 	middleware.AuthGuard(ctx)
	// })

	// Allow CORS for dev convenience
	web.InsertFilter("/*", web.AfterExec, func(ctx *context.Context) {
		ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Token")
		ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
		if ctx.Input.Method() == "OPTIONS" {
			ctx.ResponseWriter.WriteHeader(200)
		}
	})

	web.SetStaticPath("/upload", conf.GetStringConf("upload_file_path", "./upload"))

	// 记录服务器启动信息
	log.Printf("HTTP Server 已启动, 监听端口: %d", web.BConfig.Listen.HTTPPort)
	// Serve
	web.Run()
}
