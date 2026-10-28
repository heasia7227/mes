package system

import (
	"api-service/controllers"
	systemDto "api-service/dtos/system"
	"api-service/services/system"
	"common/utils/strings"
)

type UserController struct {
	controllers.BaseController
}

var (
	userService = &system.UserService{}
)

// GetUserList 获取用户列表
func (c *UserController) GetUserList() {
	params := c.GetGetParams()

	users, total := userService.GetUserList(params)

	c.JSONList(controllers.ListData{
		List: users,
		Pagination: controllers.PaginationData{
			Total:    total,
			PageNum:  params["page_num"],
			PageSize: params["page_size"],
		},
	})
}

// CreateUser 创建用户
func (c *UserController) CreateUser() {
	var user systemDto.UserDto
	if err := c.BindJSON(&user); err != nil {
		c.JSONErr(400, err.Error())
		return
	}

	user.ID = strings.GenerateID()
	user.Status = "enable"

	// 设置创建人和更新人
	login_name := c.Ctx.Input.Header("login_name")
	user.CreateBy = login_name
	user.UpdateBy = login_name

	// 保存到数据库
	err := userService.CreateUser(user)
	if err != nil {
		c.JSONErr(500, "创建用户失败: "+err.Error())
		return
	}

	c.JSONObject(map[string]interface{}{
		"success": true,
		"id":      user.ID,
	})
}

// UpdateUser 更新用户
func (c *UserController) UpdateUser() {
	id := c.Ctx.Input.Param(":id")

	// 检查用户是否存在
	existingUser, err := userService.GetUserDetail(id)
	if err != nil {
		c.JSONErr(404, "用户不存在")
		return
	}

	var user systemDto.UserDto
	if err := c.BindJSON(&user); err != nil {
		c.JSONErr(400, err.Error())
		return
	}

	// 复制必要的字段, 保留原有ID和创建信息
	existingUser.UserName = user.UserName
	existingUser.Sex = user.Sex
	existingUser.PostCode = user.PostCode
	existingUser.Phone = user.Phone
	existingUser.Email = user.Email
	existingUser.AvatarUrl = user.AvatarUrl
	existingUser.Introduction = user.Introduction

	// 设置更新人和更新时间
	login_name := c.Ctx.Input.Header("login_name")
	existingUser.UpdateBy = login_name

	// 更新到数据库
	err = userService.UpdateUser(existingUser)
	if err != nil {
		c.JSONErr(500, "更新用户失败: "+err.Error())
		return
	}

	c.JSONObject(map[string]interface{}{
		"success": true,
	})
}

// DeleteUser 删除用户
func (c *UserController) DeleteUser() {
	id := c.Ctx.Input.Param(":id")

	// 检查用户是否存在
	existingUser, err := userService.GetUserDetail(id)
	if err != nil {
		c.JSONErr(404, "用户不存在")
		return
	}
	if existingUser.Status == "enable" {
		c.JSONErr(400, "已启用的用户不能删除")
		return
	}

	login_name := c.Ctx.Input.Header("login_name")
	existingUser.UpdateBy = login_name
	existingUser.DeleteFlag = 1

	// 执行软删除
	err = userService.UpdateUser(existingUser)
	if err != nil {
		c.JSONErr(500, "删除用户失败: "+err.Error())
		return
	}

	c.JSONObject(map[string]interface{}{
		"success": true,
	})
}

// GetUserDetail 获取用户详情
func (c *UserController) GetUserDetail() {
	id := c.Ctx.Input.Param(":id")

	user, err := userService.GetUserDetail(id)
	if err != nil {
		c.JSONErr(404, "用户不存在")
		return
	}

	c.JSONObject(user)
}

// UpdateUserStatus 更新用户状态
func (c *UserController) UpdateUserStatus() {
	id := c.Ctx.Input.Param(":id")

	// 检查用户是否存在
	existingUser, err := userService.GetUserDetail(id)
	if err != nil {
		c.JSONErr(404, "用户不存在")
		return
	}

	// 状态
	status := c.Ctx.Input.Param(":status")
	if status == "" {
		c.JSONErr(400, "状态不能为空")
		return
	} else if status != "enable" && status != "disable" {
		c.JSONErr(400, "状态必须为 enable 或 disable")
		return
	}

	login_name := c.Ctx.Input.Header("login_name")
	existingUser.UpdateBy = login_name
	existingUser.Status = status

	// 更新状态
	err = userService.UpdateUser(existingUser)
	if err != nil {
		c.JSONErr(500, "更新用户状态失败: "+err.Error())
		return
	}

	c.JSONObject(map[string]interface{}{
		"success": true,
	})
}

// ChangePassword 更新用户密码
func (c *UserController) ChangePassword() {
	id := c.Ctx.Input.Param(":id")

	// 检查用户是否存在
	existingUser, err := userService.GetUserDetail(id)
	if err != nil {
		c.JSONErr(404, "用户不存在")
		return
	}

	var changePassword systemDto.UserChangePasswordDto
	if err := c.BindJSON(&changePassword); err != nil {
		c.JSONErr(400, err.Error())
		return
	}

	// 校验旧密码
	if existingUser.Password != strings.MD5Hash(changePassword.OldPassword) {
		c.JSONErr(400, "旧密码错误")
		return
	}

	// 校验新密码
	if changePassword.NewPassword == "" {
		c.JSONErr(400, "新密码不能为空")
		return
	}

	// 更新密码
	existingUser.Password = changePassword.NewPassword

	login_name := c.Ctx.Input.Header("login_name")
	existingUser.UpdateBy = login_name

	// 更新到数据库
	err = userService.UpdateUser(existingUser)
	if err != nil {
		c.JSONErr(500, "更新用户密码失败: "+err.Error())
		return
	}

	c.JSONObject(map[string]interface{}{
		"success": true,
	})
}

// ResetPassword 重置用户密码
func (c *UserController) ResetPassword() {
	id := c.Ctx.Input.Param(":id")

	// 检查用户是否存在
	existingUser, err := userService.GetUserDetail(id)
	if err != nil {
		c.JSONErr(404, "用户不存在")
		return
	}

	login_name := c.Ctx.Input.Header("login_name")
	existingUser.UpdateBy = login_name
	existingUser.Password = "123456"

	// 更新到数据库
	err = userService.UpdateUser(existingUser)
	if err != nil {
		c.JSONErr(500, "重置用户密码失败: "+err.Error())
		return
	}

	c.JSONObject(map[string]interface{}{
		"success": true,
	})
}
