package routers

import (
	"api-service/controllers/system"
	"common/utils/conf"
	"fmt"

	"github.com/beego/beego/v2/server/web"
)

func RegisterRouter(rootpath string, c web.ControllerInterface, mappingMethods ...string) {
	prefix := conf.GetStringConf("api_prefix", "/api")
	rootpath = fmt.Sprintf("%s%s", prefix, rootpath)
	web.Router(rootpath, c, mappingMethods...)
}

func init() {
	// Auth
	RegisterRouter("/auth/login", &system.LoginController{}, "post:Login")    // 登录
	RegisterRouter("/auth/captcha", &system.LoginController{}, "get:Captcha") // 获取验证码
	RegisterRouter("/auth/logout", &system.LoginController{}, "post:Logout")  // 退出登录

	// System - users
	RegisterRouter("/system/user/list", &system.UserController{}, "get:GetUserList")
	RegisterRouter("/system/user", &system.UserController{}, "post:CreateUser")
	RegisterRouter("/system/user/:id", &system.UserController{}, "get:GetUserDetail;put:UpdateUser;delete:DeleteUser")
	RegisterRouter("/system/user/:id/change-password", &system.UserController{}, "post:ChangePassword")
	RegisterRouter("/system/user/:id/reset-password", &system.UserController{}, "put:ResetPassword")
	RegisterRouter("/system/user/:id/status", &system.UserController{}, "put:UpdateUserStatus")

	// System - roles
	RegisterRouter("/system/role/list", &system.RoleController{}, "get:GetRoleList")
	RegisterRouter("/system/role", &system.RoleController{}, "post:CreateRole")
	RegisterRouter("/system/role/:id", &system.RoleController{}, "get:GetRoleDetail;put:UpdateRole;delete:DeleteRole")
	RegisterRouter("/system/role/:id/status", &system.RoleController{}, "put:UpdateRoleStatus")

	// System - config
	RegisterRouter("/system/config/list", &system.ConfigController{}, "get:GetSystemConfig")
	RegisterRouter("/system/config", &system.ConfigController{}, "post:CreateConfig")
	RegisterRouter("/system/config/:id", &system.ConfigController{}, "get:GetConfigDetail;put:UpdateConfig;delete:DeleteConfig")
	RegisterRouter("/system/config/refreshCache", &system.ConfigController{}, "put:RefreshCache")

	// System - logs
	RegisterRouter("/system/log/list", &system.LogController{}, "get:GetOperationLogList")

	// System - login logs
	RegisterRouter("/system/login-log/list", &system.LoginLogController{}, "get:GetLoginLogList")

	// System - info
	RegisterRouter("/system/info", &system.InfoController{}, "get:GetSystemInfo")
	RegisterRouter("/system/info/:id", &system.InfoController{}, "put:UpdateSystemInfo")

	// System - menu
	RegisterRouter("/system/menu/list", &system.MenuController{}, "get:GetMenuList")
	RegisterRouter("/system/menu", &system.MenuController{}, "post:CreateMenu")
	RegisterRouter("/system/menu/:id", &system.MenuController{}, "get:GetMenuDetail;put:UpdateMenu;delete:DeleteMenu")
	RegisterRouter("/system/menu/:id/status", &system.MenuController{}, "put:UpdateMenuStatus")

	// System - rank
	RegisterRouter("/system/rank/list", &system.RankController{}, "get:GetRankList")
	RegisterRouter("/system/rank", &system.RankController{}, "post:CreateRank")
	RegisterRouter("/system/rank/:id", &system.RankController{}, "get:GetRankDetail;put:UpdateRank;delete:DeleteRank")
	RegisterRouter("/system/rank/:id/status", &system.RankController{}, "put:UpdateRankStatus")

	// System - post
	RegisterRouter("/system/post/list", &system.PostController{}, "get:GetPostList")
	RegisterRouter("/system/post", &system.PostController{}, "post:CreatePost")
	RegisterRouter("/system/post/:id", &system.PostController{}, "get:GetPostDetail;put:UpdatePost;delete:DeletePost")
	RegisterRouter("/system/post/:id/status", &system.PostController{}, "put:UpdatePostStatus")

	// System - dept
	RegisterRouter("/system/dept/list", &system.DeptController{}, "get:GetDeptList")
	RegisterRouter("/system/dept", &system.DeptController{}, "post:CreateDept")
	RegisterRouter("/system/dept/:id", &system.DeptController{}, "get:GetDeptDetail;put:UpdateDept;delete:DeleteDept")
	RegisterRouter("/system/dept/:id/status", &system.DeptController{}, "put:UpdateDeptStatus")

	// 	// System - role-menu
	// 	RegisterRouter("/system/role-menu/list", &system.RoleRelationController{}, "get:GetRoleMenuList")
	// 	RegisterRouter("/system/role-menu", &system.RoleRelationController{}, "post:SaveRoleMenu")
	// 	RegisterRouter("/system/role-menu/:id", &system.RoleRelationController{}, "delete:DeleteRoleMenu")

	// 	// System - role-user
	// 	RegisterRouter("/system/role-user/list", &system.RoleRelationController{}, "get:GetRoleUserList")
	// 	RegisterRouter("/system/role-user", &system.RoleRelationController{}, "post:SaveRoleUser")
	// 	RegisterRouter("/system/role-user/:id", &system.RoleRelationController{}, "delete:DeleteRoleUser")

	// System - dictionary catalog
	RegisterRouter("/system/dictionary-catalog/list", &system.DictionaryCatalogController{}, "get:GetDictionaryCatalogList")
	RegisterRouter("/system/dictionary-catalog", &system.DictionaryCatalogController{}, "post:CreateDictionaryCatalog")
	RegisterRouter("/system/dictionary-catalog/:id", &system.DictionaryCatalogController{}, "get:GetDictionaryCatalogDetail;put:UpdateDictionaryCatalog;delete:DeleteDictionaryCatalog")

	// System - dictionary
	RegisterRouter("/system/dictionary/list", &system.DictionaryController{}, "get:GetDictionaryList")
	RegisterRouter("/system/dictionary", &system.DictionaryController{}, "post:CreateDictionary")
	RegisterRouter("/system/dictionary/:id", &system.DictionaryController{}, "get:GetDictionaryDetail;put:UpdateDictionary;delete:DeleteDictionary")
}
