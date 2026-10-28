package system

import (
	"api-service/controllers"
	systemDto "api-service/dtos/system"
	"api-service/services/system"
	"common/utils/strings"
)

type DictionaryController struct {
	controllers.BaseController
}

var (
	dictionaryService = &system.DictionaryService{}
)

// GetDictionaryList 获取字典列表
func (c *DictionaryController) GetDictionaryList() {
	params := c.GetGetParams()

	dictionarys, total := dictionaryService.GetDictionaryList(params)

	c.JSONList(controllers.ListData{
		List: dictionarys,
		Pagination: controllers.PaginationData{
			Total:    total,
			PageNum:  params["page_num"],
			PageSize: params["page_size"],
		},
	})
}

// CreateDictionary 创建字典
func (c *DictionaryController) CreateDictionary() {
	var dictionary systemDto.DictionaryDto
	if err := c.BindJSON(&dictionary); err != nil {
		c.JSONErr(400, err.Error())
		return
	}

	dictionary.ID = strings.GenerateID()

	// 设置创建人和更新人
	login_name := c.Ctx.Input.Header("login_name")
	dictionary.CreateBy = login_name
	dictionary.UpdateBy = login_name

	// 保存到数据库
	err := dictionaryService.CreateDictionary(dictionary)
	if err != nil {
		c.JSONErr(500, "创建字典失败: "+err.Error())
		return
	}

	c.JSONObject(map[string]interface{}{
		"success": true,
		"id":      dictionary.ID,
	})
}

// UpdateDictionary 更新字典
func (c *DictionaryController) UpdateDictionary() {
	id := c.Ctx.Input.Param(":id")

	// 检查字典是否存在
	existingDictionary, err := dictionaryService.GetDictionaryDetail(id)
	if err != nil {
		c.JSONErr(404, "字典不存在")
		return
	}

	var dictionary systemDto.DictionaryDto
	if err := c.BindJSON(&dictionary); err != nil {
		c.JSONErr(400, err.Error())
		return
	}

	// 复制必要的字段, 保留原有ID和创建信息
	existingDictionary.DictionaryName = dictionary.DictionaryName
	existingDictionary.DictionaryValue = dictionary.DictionaryValue
	existingDictionary.AssociationDict = dictionary.AssociationDict
	existingDictionary.AssociationDictName = dictionary.AssociationDictName
	existingDictionary.AssociationDictValue = dictionary.AssociationDictValue

	// 设置更新人和更新时间
	login_name := c.Ctx.Input.Header("login_name")
	existingDictionary.UpdateBy = login_name

	// 更新到数据库
	err = dictionaryService.UpdateDictionary(existingDictionary)
	if err != nil {
		c.JSONErr(500, "更新字典失败: "+err.Error())
		return
	}

	c.JSONObject(map[string]interface{}{
		"success": true,
	})
}

// DeleteDictionary 删除字典
func (c *DictionaryController) DeleteDictionary() {
	id := c.Ctx.Input.Param(":id")

	// 检查字典是否存在
	existingDictionary, err := dictionaryService.GetDictionaryDetail(id)
	if err != nil {
		c.JSONErr(404, "字典不存在")
		return
	}

	login_name := c.Ctx.Input.Header("login_name")
	existingDictionary.UpdateBy = login_name
	existingDictionary.DeleteFlag = 1

	// 执行软删除
	err = dictionaryService.UpdateDictionary(existingDictionary)
	if err != nil {
		c.JSONErr(500, "删除字典失败: "+err.Error())
		return
	}

	c.JSONObject(map[string]interface{}{
		"success": true,
	})
}

// GetDictionaryDetail 获取字典详情
func (c *DictionaryController) GetDictionaryDetail() {
	id := c.Ctx.Input.Param(":id")

	dictionary, err := dictionaryService.GetDictionaryDetail(id)
	if err != nil {
		c.JSONErr(404, "字典不存在")
		return
	}

	c.JSONObject(dictionary)
}
