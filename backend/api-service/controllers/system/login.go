package system

import (
	"api-service/controllers"
	systemDto "api-service/dtos/system"
	"common/models/system"
	"common/utils/captcha"
	"common/utils/http"
	"common/utils/redis"
	"common/utils/strings"
	"common/utils/token"
	"fmt"
	"image/png"
	"math/rand"
	"time"
)

type LoginController struct {
	controllers.BaseController
}

// 定时清理过期的验证码
func init() {
	go func() {
		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()
		for range ticker.C {
			captcha.CleanExpiredCaptchas()
		}
	}()

	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())
}

// Login 登录
func (c *LoginController) Login() {
	// 获取IP地址并将IPv6本地回环地址转换为IPv4本地回环地址
	loginIP := c.Ctx.Input.IP()
	if loginIP == "::1" {
		loginIP = "127.0.0.1"
	}
	// 解析User-Agent获取浏览器和操作系统
	browser, os := http.ParseUserAgent(c.Ctx.Input.UserAgent())

	var req systemDto.LoginRequest
	if err := c.BindJSON(&req); err != nil {
		loginLogService.CreateLoginLog(systemDto.LoginLogDto{
			LoginLog: system.LoginLog{
				ID:        strings.GenerateID(),
				LoginName: "unknown",
				LoginIP:   loginIP,
				Browser:   browser,
				OS:        os,
				Status:    "failed",
				Message:   "无效的请求数据格式",
			},
		})
		c.JSONErr(400, err.Error())
		return
	}

	if !captcha.ValidateCaptcha(req.CaptchaKey, req.Captcha) {
		loginLogService.CreateLoginLog(systemDto.LoginLogDto{
			LoginLog: system.LoginLog{
				ID:        strings.GenerateID(),
				LoginName: req.LoginName,
				LoginIP:   loginIP,
				Browser:   browser,
				OS:        os,
				Status:    "failed",
				Message:   "验证码错误",
			},
		})
		c.JSONErr(400, "验证码错误")
		return
	}

	user, err := userService.LoginUser(req.LoginName, req.Password)
	if err != nil {
		message := fmt.Sprintf("登录失败: %s", err.Error())
		loginLogService.CreateLoginLog(systemDto.LoginLogDto{
			LoginLog: system.LoginLog{
				ID:        strings.GenerateID(),
				LoginName: req.LoginName,
				LoginIP:   loginIP,
				Browser:   browser,
				OS:        os,
				Status:    "failed",
				Message:   message,
				LoginTime: time.Now(),
			},
		})
		c.JSONErr(401, message)
		return
	}

	tokenStr, err := token.GenerateToken(user.ID, user.UserName, user.LoginName)
	if err != nil {
		c.JSONErr(500, "生成Token失败: "+err.Error())
		return
	}
	// 存储token到redis
	redis.Set("token:"+user.ID, tokenStr, 0)

	// 记录登录日志
	loginLogService.CreateLoginLog(systemDto.LoginLogDto{
		LoginLog: system.LoginLog{
			ID:        strings.GenerateID(),
			LoginName: user.LoginName,
			LoginIP:   loginIP,
			Browser:   browser,
			OS:        os,
			Status:    "success",
			Message:   "登录成功",
			LoginTime: time.Now(),
		},
	})

	c.JSONObject(map[string]interface{}{
		"token": tokenStr,
		"user": map[string]interface{}{
			"id":         user.ID,
			"login_name": user.LoginName,
			"user_name":  user.UserName,
			"avatar":     user.AvatarUrl,
		}})
}

// 生成验证码图片接口
func (c *LoginController) Captcha() {
	// 从请求参数中获取验证码键
	captchaKey := c.GetString("captcha_key")
	if captchaKey == "" {
		c.JSONErr(400, "验证码键不能为空")
		return
	}

	// 生成验证码图片
	img := captcha.GenerateImage(captchaKey)

	// 设置响应头
	c.Ctx.Output.Header("Content-Type", "image/png")
	c.Ctx.Output.Header("Cache-Control", "no-store, no-cache, must-revalidate")
	c.Ctx.Output.Header("Pragma", "no-cache")

	// 将图片编码为PNG并写入响应
	if err := png.Encode(c.Ctx.ResponseWriter, img); err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte("生成验证码失败"))
	}
}

// Logout 退出登录
func (c *LoginController) Logout() {
	tokenStr := c.Ctx.Input.Header("X-Token")
	// If X-Token is empty, try Authorization header
	if tokenStr == "" {
		tokenStr = c.Ctx.Input.Header("Authorization")
	}
	// If Authorization header is empty, return error
	if tokenStr == "" {
		c.JSONErr(400, "Authorization header is empty")
		return
	}

	user := userService.GetUserByToken(tokenStr)
	if user != nil {
		// 获取IP地址并将IPv6本地回环地址转换为IPv4本地回环地址
		loginIP := c.Ctx.Input.IP()
		if loginIP == "::1" {
			loginIP = "127.0.0.1"
		}

		// 解析User-Agent获取浏览器和操作系统
		browser, os := http.ParseUserAgent(c.Ctx.Input.UserAgent())

		// 记录退出登录日志
		loginLogService.CreateLoginLog(systemDto.LoginLogDto{
			LoginLog: system.LoginLog{
				ID:        strings.GenerateID(),
				LoginName: user.LoginName,
				LoginIP:   loginIP,
				Browser:   browser,
				OS:        os,
				Status:    "success",
				Message:   "退出登录",
				LoginTime: time.Now(),
			},
		})
	}
	redis.Del("token:" + user.ID)
	c.JSONObject(map[string]interface{}{"code": 200, "msg": "退出成功"})
}
