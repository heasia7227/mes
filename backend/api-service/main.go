package main

import (
	"log"
	"net/http"

	_ "api-service/routers" // 初始化路由
	_ "common/db"           // 初始化数据库连接
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

	// Allow CORS
	// 注意: 必须注册在 BeforeRouter 阶段。OPTIONS 预检请求不会匹配到具体 Controller,
	// 挂在 AfterExec 时预检请求根本拿不到响应头, 浏览器会以 CORS 错误拦截后续的真实请求。
	web.InsertFilter("/*", web.BeforeRouter, func(ctx *context.Context) {
		origin := ctx.Input.Header("Origin")
		if origin == "" {
			origin = "*"
		}

		header := ctx.ResponseWriter.Header()
		header.Set("Access-Control-Allow-Origin", origin)
		header.Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		header.Set("Access-Control-Allow-Headers", "Content-Type,Authorization,X-Token")
		header.Set("Access-Control-Allow-Credentials", "true")
		header.Set("Access-Control-Max-Age", "86400")
		header.Set("Vary", "Origin")

		// 预检请求直接响应 200 并结束: 写入响应头后 Started 被置位, 后续路由不再执行
		if ctx.Input.Method() == "OPTIONS" {
			ctx.ResponseWriter.WriteHeader(http.StatusOK)
		}
	})

	web.SetStaticPath("/upload", conf.GetStringConf("upload_file_path", "./upload"))

	// 记录服务器启动信息
	log.Printf("HTTP Server 已启动, 监听端口: %d", web.BConfig.Listen.HTTPPort)
	// Serve
	web.Run()
}
