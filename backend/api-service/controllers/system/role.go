package system

import (
	"api-service/controllers"
	systemDto "api-service/dtos/system"
	"api-service/services/system"
	"common/utils/strings"
)

type RoleController struct {
	controllers.BaseController
}

var (
	roleService = &system.RoleService{}
)

// GetRoleList 获取角色列表
func (c *RoleController) GetRoleList() {
	params := c.GetGetParams()

	roles, total := roleService.GetRoleList(params)

	c.JSONList(controllers.ListData{
		List: roles,
		Pagination: controllers.PaginationData{
			Total:    total,
			PageNum:  params["page_num"],
			PageSize: params["page_size"],
		},
	})
}

// CreateRole 创建角色
func (c *RoleController) CreateRole() {
	var role systemDto.RoleDto
	if err := c.BindJSON(&role); err != nil {
		c.JSONErr(400, err.Error())
		return
	}

	role.ID = strings.GenerateID()
	role.Status = "enable"

	// 设置创建人和更新人
	login_name := c.Ctx.Input.Header("login_name")
	role.CreateBy = login_name
	role.UpdateBy = login_name

	// 保存到数据库
	err := roleService.CreateRole(role)
	if err != nil {
		c.JSONErr(500, "创建角色失败: "+err.Error())
		return
	}

	c.JSONObject(map[string]interface{}{
		"success": true,
		"id":      role.ID,
	})
}

// UpdateRole 更新角色
func (c *RoleController) UpdateRole() {
	id := c.Ctx.Input.Param(":id")

	// 检查角色是否存在
	existingRole, err := roleService.GetRoleDetail(id)
	if err != nil {
		c.JSONErr(404, "角色不存在")
		return
	}

	var role systemDto.RoleDto
	if err := c.BindJSON(&role); err != nil {
		c.JSONErr(400, err.Error())
		return
	}

	// 复制必要的字段, 保留原有ID和创建信息
	existingRole.Title = role.Title
	existingRole.Code = role.Code
	existingRole.Description = role.Description

	// 设置更新人和更新时间
	login_name := c.Ctx.Input.Header("login_name")
	existingRole.UpdateBy = login_name

	// 更新到数据库
	err = roleService.UpdateRole(existingRole)
	if err != nil {
		c.JSONErr(500, "更新角色失败: "+err.Error())
		return
	}

	c.JSONObject(map[string]interface{}{
		"success": true,
	})
}

// DeleteRole 删除角色
func (c *RoleController) DeleteRole() {
	id := c.Ctx.Input.Param(":id")

	// 检查角色是否存在
	existingRole, err := roleService.GetRoleDetail(id)
	if err != nil {
		c.JSONErr(404, "角色不存在")
		return
	}
	if existingRole.Status == "enable" {
		c.JSONErr(400, "已启用的角色不能删除")
		return
	}

	login_name := c.Ctx.Input.Header("login_name")
	existingRole.UpdateBy = login_name
	existingRole.DeleteFlag = 1

	// 执行软删除
	err = roleService.UpdateRole(existingRole)
	if err != nil {
		c.JSONErr(500, "删除角色失败: "+err.Error())
		return
	}

	c.JSONObject(map[string]interface{}{
		"success": true,
	})
}

// GetRoleDetail 获取角色详情
func (c *RoleController) GetRoleDetail() {
	id := c.Ctx.Input.Param(":id")

	role, err := roleService.GetRoleDetail(id)
	if err != nil {
		c.JSONErr(404, "角色不存在")
		return
	}

	c.JSONObject(role)
}

// UpdateRoleStatus 更新角色状态
func (c *RoleController) UpdateRoleStatus() {
	id := c.Ctx.Input.Param(":id")

	// 检查角色是否存在
	existingRole, err := roleService.GetRoleDetail(id)
	if err != nil {
		c.JSONErr(404, "角色不存在")
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
	existingRole.UpdateBy = login_name
	existingRole.Status = status

	// 更新状态
	err = roleService.UpdateRole(existingRole)
	if err != nil {
		c.JSONErr(500, "更新角色状态失败: "+err.Error())
		return
	}

	c.JSONObject(map[string]interface{}{
		"success": true,
	})
}
