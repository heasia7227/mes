package middleware

import (
	"api-service/services/system"
	"common/utils/strings"

	"github.com/beego/beego/v2/server/web/context"
)

var logService = system.LogService{}

func RequestLogger(ctx *context.Context) {
	// 记录开始时间
	// startTime := time.Now()

	// Skip preflight
	method := ctx.Input.Method()
	if method == "OPTIONS" {
		return
	}

	// 过滤掉 GET 请求的日志, 但保留其他涉及增删改的请求
	if method == "GET" {
		return
	}

	path := ctx.Request.URL.Path

	// Skip 日志列表和查看
	if strings.HasPrefix(path, "/api/system/log") {
		return
	}

	// Skip 登录请求
	if strings.HasPrefix(path, "/api/auth/login") {
		return
	}

	// Skip 上传工艺文件
	if strings.HasPrefix(path, "/api/file/upload/chunk") {
		return
	}

	// login_name := ctx.Input.Header("login_name")

	// // 获取请求参数
	// params := ""
	// if len(ctx.Input.RequestBody) > 0 {
	// 	params = string(ctx.Input.RequestBody)
	// } else if ctx.Request.Body != nil {
	// 	bodyBytes, err := io.ReadAll(ctx.Request.Body)
	// 	if err != nil {
	// 		params = "[无法读取请求体]"
	// 	} else {
	// 		params = string(bodyBytes)
	// 		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	// 	}
	// }

	// // 从URL路径或页面路径中提取模块信息
	// module := ""

	// // 从请求头中获取页面路径
	// pagePath := ctx.Input.Header("X-Page-Path")

	// // 优先从页面路径中提取模块信息
	// if pagePath != "" {
	// 	// 去掉页面路径末尾的斜杠
	// 	pagePath = strings.TrimSuffix(pagePath, "/")
	// 	// 映射页面路径到系统模块
	// 	pagePathMap := map[string]string{
	// 		"/printer/printer-model":            "打印设备型号",
	// 		"/printer/printer":                  "打印设备",
	// 		"/printer/printer-software-version": "打印设备软件版本",
	// 		"/cleaner-line/cleaner-line":        "清粉线",
	// 		"/cleaner-line/cleaner-line-device": "清粉线设备",
	// 		"/cleaner-line/alarm-config":        "清粉线报警配置",
	// 		"/powder-line/powder-line":          "粉末循环线",
	// 		"/powder-line/powder-line-device":   "粉末循环线设备",
	// 		"/powder-line/alarm-config":         "粉末循环线报警配置",
	// 		"/powder-line/powder-line-task":     "粉末循环线任务",
	// 		"/agv/agv":                          "AGV 小车",
	// 		"/agv/agv-task":                     "AGV 小车任务",
	// 		"/resource/powder-type":             "粉末类型",
	// 		"/resource/cylinder":                "缸体管理",
	// 		"/resource/substrate":               "基材管理",
	// 		"/resource/storage-location":        "存储位置",
	// 		"/resource/storage-area":            "区域管理",
	// 		"/production/production-task":       "生产任务",
	// 		"/production/quality-traceability":  "质量追溯",
	// 		"/inout/cylinder-in":                "缸体入库",
	// 		"/inout/substrate-in":               "基材入库",
	// 		"/inout/part-out":                   "零件出库",
	// 		"/inout/cylinder-out":               "缸体出库",
	// 		"/inout/powder-in":                  "缸体出库",
	// 		"/system/user":                      "用户管理",
	// 		"/system/role":                      "角色权限",
	// 		"/system/menu":                      "菜单管理",
	// 		"/system/dictionary":                "字典管理",
	// 		"/system/config":                    "系统内容配置",
	// 		"/system/log":                       "操作日志",
	// 		"/system/login-log":                 "登录日志",
	// 		"/system/rank":                      "职级管理",
	// 		"/system/post":                      "岗位管理",
	// 		"/system/dept":                      "部门管理",
	// 		"/system/info":                      "系统信息",
	// 	}

	// 	// 检查页面路径是否在映射中, 如果不在, 则直接跳过
	// 	mappedModule, ok := pagePathMap[pagePath]
	// 	if !ok {
	// 		return
	// 	} else {
	// 		module = mappedModule
	// 	}
	// } else { // 直接访问接口
	// 	apiPaths := map[string]string{
	// 		"/api/cylinder/inbound":  "缸体入库",
	// 		"/api/cylinder/outbound": "缸体出库",
	// 		"/api/substrate/inbound": "基材入库",
	// 		"/api/part/outbound":     "零件出库",
	// 		"/api/powder/inbound":    "粉体入库",
	// 	}
	// 	// 检查API路径是否在映射中, 如果不在, 则直接跳过
	// 	for apiPath, mappedModule := range apiPaths {
	// 		if strings.HasPrefix(path, apiPath) {
	// 			module = mappedModule
	// 			break
	// 		}
	// 	}
	// }

	// // 检查模块是否为空, 如果是, 则直接跳过
	// if module == "" {
	// 	return
	// }

	// // 生成操作类型
	// action := ""
	// switch method {
	// case "POST":
	// 	action = "新增"
	// case "PUT":
	// 	action = "修改"
	// case "DELETE":
	// 	action = "删除"
	// default:
	// 	action = "其他"
	// }

	// // 生成操作描述
	// description := ""
	// switch method {
	// case "POST":
	// 	description = "新增数据"
	// case "PUT":
	// 	description = "修改数据"
	// case "DELETE":
	// 	description = "删除数据"
	// default:
	// 	description = "其他操作"
	// }

	// // 定义状态变量
	// status := "成功"
	// isException := "0"

	// // 使用 defer 确保在请求处理完成后记录日志
	// defer func() {
	// 	// 计算消耗时间
	// 	duration := int(time.Since(startTime).Milliseconds())

	// 	// 处理IP地址, 将IPv6的本地回环地址转换为IPv4格式
	// 	ip := ctx.Input.IP()
	// 	if ip == "::1" {
	// 		ip = "127.0.0.1"
	// 	}

	// 	logService.CreateLog(&systemDto.LogDto{
	// 		ID:            strings.GenerateID(),
	// 		Ip:            ip,
	// 		LogContent:    fmt.Sprintf("路径: %s, 方法: %s, 参数: %s", path, method, params),
	// 		IsException:   isException,
	// 		OperateType:   action,
	// 		OperateModule: module,
	// 		OperateBy:     login_name,
	// 		OperateDate:   time.Now(),
	// 		Status:        status,
	// 		Duration:      duration,
	// 		Description:   description,
	// 	})
	// }()
}
