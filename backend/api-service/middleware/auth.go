package middleware

import (
	"api-service/services/system"
	"strings"

	"github.com/beego/beego/v2/server/web/context"
)

var (
	userService = &system.UserService{}
)

func AuthGuard(ctx *context.Context) {
	path := ctx.Input.URL()
	// 认证接口、大屏接口直接放行
	if strings.HasPrefix(path, "/api/auth") ||
		strings.HasPrefix(path, "/api/bigscreen") {
		return
	}
	// Token check - support both X-Token and Authorization header
	token := ctx.Input.Header("X-Token")
	// If X-Token is empty, try Authorization header
	if token == "" {
		token = ctx.Input.Header("Authorization")
	}
	if token == "" {
		ctx.ResponseWriter.WriteHeader(401)
		ctx.Output.JSON(map[string]interface{}{"code": 401, "message": "Unauthorized"}, false, false)
	} else {
		// 校验Token的合法性
		user := userService.GetUserByToken(token)
		if user == nil {
			ctx.ResponseWriter.WriteHeader(401)
			ctx.Output.JSON(map[string]interface{}{"code": 401, "message": "Unauthorized"}, false, false)
		} else {
			// 头信息写入userId
			ctx.Request.Header.Set("user_id", user.ID)
			ctx.Request.Header.Set("user_name", user.UserName)
			ctx.Request.Header.Set("login_name", user.LoginName)
		}
	}
}
