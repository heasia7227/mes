package system

import (
	"api-service/controllers"
	systemDto "api-service/dtos/system"
	"api-service/services/system"
	"common/utils/strings"
)

type MenuController struct {
	controllers.BaseController
}

var (
	menuService = &system.MenuService{}
)

// GetMenuList 获取菜单列表
func (c *MenuController) GetMenuList() {
	params := c.GetGetParams()

	menus, total := menuService.GetMenuList(params)

	c.JSONList(controllers.ListData{
		List: menus,
		Pagination: controllers.PaginationData{
			Total:    total,
			PageNum:  params["page_num"],
			PageSize: params["page_size"],
		},
	})
}

// CreateMenu 创建菜单
func (c *MenuController) CreateMenu() {
	var menu systemDto.MenuDto
	if err := c.BindJSON(&menu); err != nil {
		c.JSONErr(400, err.Error())
		return
	}

	menu.ID = strings.GenerateID()
	menu.Status = "enable"

	// 设置创建人和更新人
	login_name := c.Ctx.Input.Header("login_name")
	menu.CreateBy = login_name
	menu.UpdateBy = login_name

	// 保存到数据库
	err := menuService.CreateMenu(menu)
	if err != nil {
		c.JSONErr(500, "创建菜单失败: "+err.Error())
		return
	}

	c.JSONObject(map[string]interface{}{
		"success": true,
		"id":      menu.ID,
	})
}

// UpdateMenu 更新菜单
func (c *MenuController) UpdateMenu() {
	id := c.Ctx.Input.Param(":id")

	// 检查菜单是否存在
	existingMenu, err := menuService.GetMenuDetail(id)
	if err != nil {
		c.JSONErr(404, "菜单不存在")
		return
	}

	var menu systemDto.MenuDto
	if err := c.BindJSON(&menu); err != nil {
		c.JSONErr(400, err.Error())
		return
	}

	// 复制必要的字段, 保留原有ID和创建信息
	existingMenu.Title = menu.Title
	existingMenu.Icon = menu.Icon
	existingMenu.ParentID = menu.ParentID
	existingMenu.Ranking = menu.Ranking
	existingMenu.MenuType = menu.MenuType
	existingMenu.DesignType = menu.DesignType
	existingMenu.PageType = menu.PageType
	existingMenu.MenuPath = menu.MenuPath
	existingMenu.Permission = menu.Permission

	// 设置更新人和更新时间
	login_name := c.Ctx.Input.Header("login_name")
	existingMenu.UpdateBy = login_name

	// 更新到数据库
	err = menuService.UpdateMenu(existingMenu)
	if err != nil {
		c.JSONErr(500, "更新菜单失败: "+err.Error())
		return
	}

	c.JSONObject(map[string]interface{}{
		"success": true,
	})
}

// DeleteMenu 删除菜单
func (c *MenuController) DeleteMenu() {
	id := c.Ctx.Input.Param(":id")

	// 检查菜单是否存在
	existingMenu, err := menuService.GetMenuDetail(id)
	if err != nil {
		c.JSONErr(404, "菜单不存在")
		return
	}
	if existingMenu.Status == "enable" {
		c.JSONErr(400, "已启用的菜单不能删除")
		return
	}

	login_name := c.Ctx.Input.Header("login_name")
	existingMenu.UpdateBy = login_name
	existingMenu.DeleteFlag = 1

	// 执行软删除
	err = menuService.UpdateMenu(existingMenu)
	if err != nil {
		c.JSONErr(500, "删除菜单失败: "+err.Error())
		return
	}

	c.JSONObject(map[string]interface{}{
		"success": true,
	})
}

// GetMenuDetail 获取菜单详情
func (c *MenuController) GetMenuDetail() {
	id := c.Ctx.Input.Param(":id")

	menu, err := menuService.GetMenuDetail(id)
	if err != nil {
		c.JSONErr(404, "菜单不存在")
		return
	}

	c.JSONObject(menu)
}

// UpdateMenuStatus 更新菜单状态
func (c *MenuController) UpdateMenuStatus() {
	id := c.Ctx.Input.Param(":id")

	// 检查菜单是否存在
	existingMenu, err := menuService.GetMenuDetail(id)
	if err != nil {
		c.JSONErr(404, "菜单不存在")
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
	existingMenu.UpdateBy = login_name
	existingMenu.Status = status

	// 更新状态
	err = menuService.UpdateMenu(existingMenu)
	if err != nil {
		c.JSONErr(500, "更新菜单状态失败: "+err.Error())
		return
	}

	c.JSONObject(map[string]interface{}{
		"success": true,
	})
}
