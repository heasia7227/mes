package routers

import (
	"common/utils/conf"
	"fmt"

	"github.com/beego/beego/v2/server/web"
)

func RegisterRouter(rootpath string, c web.ControllerInterface, mappingMethods ...string) {
	prefix := conf.GetStringConf("api_prefix", "/api")
	rootpath = fmt.Sprintf("%s%s", prefix, rootpath)
	web.Router(rootpath, c, mappingMethods...)
}

// func init() {
// 	// Auth
// 	RegisterRouter("/auth/login", &system.AuthController{}, "post:Login")                 // 登录
// 	RegisterRouter("/auth/captcha", &system.AuthController{}, "get:Captcha")              // 获取验证码
// 	RegisterRouter("/auth/logout", &system.AuthController{}, "post:Logout")               // 退出登录
// 	RegisterRouter("/auth/validate-token", &system.AuthController{}, "get:ValidateToken") // 校验Token

// 	// System - users
// 	RegisterRouter("/system/user/list", &system.UserController{}, "get:GetUserList")
// 	RegisterRouter("/system/user", &system.UserController{}, "post:CreateUser")
// 	RegisterRouter("/auth/change-password", &system.UserController{}, "post:ChangePassword")
// 	RegisterRouter("/system/user/profile", &system.UserController{}, "put:UpdateUserProfile")
// 	RegisterRouter("/system/user/:id", &system.UserController{}, "get:GetUserDetail;put:UpdateUser;delete:DeleteUser")
// 	RegisterRouter("/system/user/:id/reset-password", &system.UserController{}, "put:ResetPassword")
// 	RegisterRouter("/system/user/:id/status", &system.UserController{}, "put:ChangeUserStatus")
// 	RegisterRouter("/system/user/:id/avatar", &system.UserController{}, "put:UpdateUserAvatar")
// 	RegisterRouter("/system/user/:id/force-logout", &system.UserController{}, "post:ForceLogout")
// 	RegisterRouter("/system/user/UpdateUserAvatar", &system.UserController{}, "post:UpdateUserAvatarBySelf")
// 	RegisterRouter("/system/user/:id/enable", &system.UserController{}, "post:EnableUser")

// 	// System - roles
// 	RegisterRouter("/system/role/list", &system.RoleController{}, "get:GetRoleList")
// 	RegisterRouter("/system/role", &system.RoleController{}, "post:CreateRole")
// 	RegisterRouter("/system/role/:id", &system.RoleController{}, "get:GetRoleDetail;put:UpdateRole;delete:DeleteRole")
// 	RegisterRouter("/system/role/stats", &system.RoleController{}, "get:GetRoleStats")
// 	RegisterRouter("/system/role/:id/enable", &system.RoleController{}, "post:EnableRole")
// 	RegisterRouter("/system/role/:id/disable", &system.RoleController{}, "post:DisableRole")

// 	// System - config
// 	RegisterRouter("/system/config/list", &system.ConfigController{}, "get:GetSystemConfig")
// 	RegisterRouter("/system/config/value", &system.ConfigController{}, "get:GetConfigValue")
// 	RegisterRouter("/system/config/:id", &system.ConfigController{}, "get:GetConfigDetail;delete:DeleteConfig")
// 	RegisterRouter("/system/config", &system.ConfigController{}, "post:AddConfig;put:UpdateSystemConfig")
// 	RegisterRouter("/system/config/refreshCache", &system.ConfigController{}, "delete:RefreshCache")
// 	RegisterRouter("/system/config/logo", &system.ConfigController{}, "post:UploadLogo")

// 	// System - logs
// 	RegisterRouter("/system/log/list", &system.LogController{}, "get:GetOperationLogList")
// 	RegisterRouter("/system/log/export", &system.LogController{}, "get:ExportOperationLogs")
// 	RegisterRouter("/system/log/stats", &system.LogController{}, "get:GetLogStats")

// 	// System - info
// 	RegisterRouter("/system/info/list", &system.SystemInfoController{}, "get:GetSystemInfoList")
// 	RegisterRouter("/system/info", &system.SystemInfoController{}, "post:AddSystemInfo;put:UpdateSystemInfo")
// 	RegisterRouter("/system/info/:id", &system.SystemInfoController{}, "get:GetSystemInfo;delete:DeleteSystemInfo")
// 	RegisterRouter("/system/info/current", &system.SystemInfoController{}, "get:GetCurrentSystemInfo")
// 	RegisterRouter("/system/info/:id/enable", &system.SystemInfoController{}, "post:EnableSystemInfo")

// 	// System - menu
// 	RegisterRouter("/system/menu/list", &system.MenuController{}, "get:GetMenuList")
// 	RegisterRouter("/system/menu", &system.MenuController{}, "post:CreateMenu")
// 	RegisterRouter("/system/menu/:id", &system.MenuController{}, "get:GetMenuDetail;put:UpdateMenu;delete:DeleteMenu")
// 	RegisterRouter("/system/menu/user/list", &system.MenuController{}, "get:GetUserMenuList")

// 	// System - rank
// 	RegisterRouter("/system/rank/list", &system.RankController{}, "get:GetRankList")
// 	RegisterRouter("/system/rank", &system.RankController{}, "post:CreateRank")
// 	RegisterRouter("/system/rank/:id", &system.RankController{}, "get:GetRankDetail;put:UpdateRank;delete:DeleteRank")
// 	RegisterRouter("/system/rank/:id/status", &system.RankController{}, "put:ChangeRankStatus")

// 	// System - post
// 	RegisterRouter("/system/post/list", &system.PostController{}, "get:GetPostList")
// 	RegisterRouter("/system/post", &system.PostController{}, "post:CreatePost")
// 	RegisterRouter("/system/post/:id", &system.PostController{}, "get:GetPostDetail;put:UpdatePost;delete:DeletePost")
// 	RegisterRouter("/system/post/:id/status", &system.PostController{}, "put:ChangePostStatus")

// 	// System - dept
// 	RegisterRouter("/system/dept/list", &system.DeptController{}, "get:GetDeptList")
// 	RegisterRouter("/system/dept", &system.DeptController{}, "post:CreateDept")
// 	RegisterRouter("/system/dept/:id", &system.DeptController{}, "get:GetDeptDetail;put:UpdateDept;delete:DeleteDept")
// 	RegisterRouter("/system/dept/:id/status", &system.DeptController{}, "put:ChangeDeptStatus")

// 	// System - login logs
// 	RegisterRouter("/system/login-log/list", &system.LoginLogController{}, "get:GetLoginLogList")
// 	RegisterRouter("/system/login-log/:id", &system.LoginLogController{}, "delete:DeleteLoginLog")
// 	RegisterRouter("/system/login-log/clear", &system.LoginLogController{}, "post:ClearLoginLogs")

// 	// System - role-menu
// 	RegisterRouter("/system/role-menu/list", &system.RoleRelationController{}, "get:GetRoleMenuList")
// 	RegisterRouter("/system/role-menu", &system.RoleRelationController{}, "post:SaveRoleMenu")
// 	RegisterRouter("/system/role-menu/:id", &system.RoleRelationController{}, "delete:DeleteRoleMenu")

// 	// System - role-user
// 	RegisterRouter("/system/role-user/list", &system.RoleRelationController{}, "get:GetRoleUserList")
// 	RegisterRouter("/system/role-user", &system.RoleRelationController{}, "post:SaveRoleUser")
// 	RegisterRouter("/system/role-user/:id", &system.RoleRelationController{}, "delete:DeleteRoleUser")

// 	// System - dictionary catalog
// 	RegisterRouter("/system/dictionary-catalog/list", &system.DictionaryController{}, "get:GetDictionaryCatalogList")
// 	RegisterRouter("/system/dictionary-catalog", &system.DictionaryController{}, "post:CreateDictionaryCatalog")
// 	RegisterRouter("/system/dictionary-catalog/:id", &system.DictionaryController{}, "put:UpdateDictionaryCatalog;delete:DeleteDictionaryCatalog")

// 	// System - dictionary
// 	RegisterRouter("/system/dictionary/list", &system.DictionaryController{}, "get:GetDictionaryList")
// 	RegisterRouter("/system/dictionary/list-by-code", &system.DictionaryController{}, "get:GetDictionaryListByCode")
// 	RegisterRouter("/system/dictionary/list-by-code-with-association", &system.DictionaryController{}, "get:GetDictionaryListByCodeWithAssociation")
// 	RegisterRouter("/system/dictionary/by-catalog-value", &system.DictionaryController{}, "get:GetDictionaryByCatalogAndValue")
// 	RegisterRouter("/system/dictionary/list-by-catalog-ids", &system.DictionaryController{}, "get:GetDictionaryListByCatalogIDs")
// 	RegisterRouter("/system/dictionary", &system.DictionaryController{}, "post:CreateDictionary")
// 	RegisterRouter("/system/dictionary/:id", &system.DictionaryController{}, "put:UpdateDictionary;delete:DeleteDictionary")
// }
