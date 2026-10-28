package system

import (
	"api-service/controllers"
	systemDto "api-service/dtos/system"
	"api-service/services/system"
	"common/utils/strings"
)

type DictionaryCatalogController struct {
	controllers.BaseController
}

var (
	dictionarycatalogService = &system.DictionaryCatalogService{}
)

// GetDictionaryCatalogList 获取字典目录列表
func (c *DictionaryCatalogController) GetDictionaryCatalogList() {
	params := c.GetGetParams()

	dictionarycatalogs, total := dictionarycatalogService.GetDictionaryCatalogList(params)

	c.JSONList(controllers.ListData{
		List: dictionarycatalogs,
		Pagination: controllers.PaginationData{
			Total:    total,
			PageNum:  params["page_num"],
			PageSize: params["page_size"],
		},
	})
}

// CreateDictionaryCatalog 创建字典目录
func (c *DictionaryCatalogController) CreateDictionaryCatalog() {
	var dictionarycatalog systemDto.DictionaryCatalogDto
	if err := c.BindJSON(&dictionarycatalog); err != nil {
		c.JSONErr(400, err.Error())
		return
	}

	dictionarycatalog.ID = strings.GenerateID()

	// 设置创建人和更新人
	login_name := c.Ctx.Input.Header("login_name")
	dictionarycatalog.CreateBy = login_name
	dictionarycatalog.UpdateBy = login_name

	// 保存到数据库
	err := dictionarycatalogService.CreateDictionaryCatalog(dictionarycatalog)
	if err != nil {
		c.JSONErr(500, "创建字典目录失败: "+err.Error())
		return
	}

	c.JSONObject(map[string]interface{}{
		"success": true,
		"id":      dictionarycatalog.ID,
	})
}

// UpdateDictionaryCatalog 更新字典目录
func (c *DictionaryCatalogController) UpdateDictionaryCatalog() {
	id := c.Ctx.Input.Param(":id")

	// 检查字典目录是否存在
	existingDictionaryCatalog, err := dictionarycatalogService.GetDictionaryCatalogDetail(id)
	if err != nil {
		c.JSONErr(404, "字典目录不存在")
		return
	}

	var dictionarycatalog systemDto.DictionaryCatalogDto
	if err := c.BindJSON(&dictionarycatalog); err != nil {
		c.JSONErr(400, err.Error())
		return
	}

	// 复制必要的字段, 保留原有ID和创建信息
	existingDictionaryCatalog.CatalogName = dictionarycatalog.CatalogName
	existingDictionaryCatalog.CatalogCode = dictionarycatalog.CatalogCode
	existingDictionaryCatalog.ParentID = dictionarycatalog.ParentID

	// 设置更新人和更新时间
	login_name := c.Ctx.Input.Header("login_name")
	existingDictionaryCatalog.UpdateBy = login_name

	// 更新到数据库
	err = dictionarycatalogService.UpdateDictionaryCatalog(existingDictionaryCatalog)
	if err != nil {
		c.JSONErr(500, "更新字典目录失败: "+err.Error())
		return
	}

	c.JSONObject(map[string]interface{}{
		"success": true,
	})
}

// DeleteDictionaryCatalog 删除字典目录
func (c *DictionaryCatalogController) DeleteDictionaryCatalog() {
	id := c.Ctx.Input.Param(":id")

	// 检查字典目录是否存在
	existingDictionaryCatalog, err := dictionarycatalogService.GetDictionaryCatalogDetail(id)
	if err != nil {
		c.JSONErr(404, "字典目录不存在")
		return
	}

	login_name := c.Ctx.Input.Header("login_name")
	existingDictionaryCatalog.UpdateBy = login_name
	existingDictionaryCatalog.DeleteFlag = 1

	// 执行软删除
	err = dictionarycatalogService.UpdateDictionaryCatalog(existingDictionaryCatalog)
	if err != nil {
		c.JSONErr(500, "删除字典目录失败: "+err.Error())
		return
	}

	c.JSONObject(map[string]interface{}{
		"success": true,
	})
}

// GetDictionaryCatalogDetail 获取字典目录详情
func (c *DictionaryCatalogController) GetDictionaryCatalogDetail() {
	id := c.Ctx.Input.Param(":id")

	dictionarycatalog, err := dictionarycatalogService.GetDictionaryCatalogDetail(id)
	if err != nil {
		c.JSONErr(404, "字典目录不存在")
		return
	}

	c.JSONObject(dictionarycatalog)
}
