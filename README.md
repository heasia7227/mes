MES 制造执行系统 (Manufacturing Exection System)

# 项目介绍

frontend: 前端项目, 基于 Next.js + Ant Design 实现
backend: 后端项目, 基于 Beego 实现

- api-service: API 服务, 提供 RESTful API 接口
  - controllers: 控制器, 处理 HTTP 请求和响应
  - routers: 路由, 配置 HTTP 路由
  - services: 服务, 处理业务逻辑
- worker-service: 任务服务, 处理业务逻辑
  - tasks: 任务, 处理业务逻辑
  - workers: 工作线程, 处理任务
- common: 公共模块, 提供通用的工具函数和常量
  - db: 数据库项目, 基于 MySQL 实现
  - models: 数据库模型, 定义数据库表结构
  - utils: 工具函数, 提供通用的工具函数和常量
    - conf: 读取应用配置
    - goroutine: Goroutine 工具函数
    - http: HTTP 工具函数
    - mqtt: MQTT 工具函数
    - opcua: OPC UA 工具函数
    - redis: Redis 工具函数
    - strings: 字符串工具函数
    - token: 令牌工具函数
    - wlogs: 日志工具函数
- conf: 配置模块, 提供应用配置
  - app.conf: 应用配置文件
  - app.dev.conf: 开发环境配置文件
  - app.prod.conf: 生产环境配置文件
