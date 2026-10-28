package system

import (
	"api-service/controllers"
	systemDto "api-service/dtos/system"
	"api-service/services/system"
	"common/utils/strings"
)

type ConfigController struct {
	controllers.BaseController
}

var (
	configService = &system.ConfigService{}
)

// GetConfigList 获取系统配置列表
func (c *ConfigController) GetConfigList() {
	params := c.GetGetParams()

	configs, total := configService.GetConfigList(params)

	c.JSONList(controllers.ListData{
		List: configs,
		Pagination: controllers.PaginationData{
			Total:    total,
			PageNum:  params["page_num"],
			PageSize: params["page_size"],
		},
	})
}

// CreateConfig 创建系统配置
func (c *ConfigController) CreateConfig() {
	var config systemDto.ConfigDto
	if err := c.BindJSON(&config); err != nil {
		c.JSONErr(400, err.Error())
		return
	}

	config.ID = strings.GenerateID()

	// 设置创建人和更新人
	login_name := c.Ctx.Input.Header("login_name")
	config.CreateBy = login_name
	config.UpdateBy = login_name

	// 保存到数据库
	err := configService.CreateConfig(config)
	if err != nil {
		c.JSONErr(500, "创建系统配置失败: "+err.Error())
		return
	}

	c.JSONObject(map[string]interface{}{
		"success": true,
		"id":      config.ID,
	})
}

// UpdateConfig 更新系统配置
func (c *ConfigController) UpdateConfig() {
	id := c.Ctx.Input.Param(":id")

	// 检查系统配置是否存在
	existingConfig, err := configService.GetConfigDetail(id)
	if err != nil {
		c.JSONErr(404, "系统配置不存在")
		return
	}

	var config systemDto.ConfigDto
	if err := c.BindJSON(&config); err != nil {
		c.JSONErr(400, err.Error())
		return
	}

	// 复制必要的字段, 保留原有ID和创建信息
	existingConfig.ConfigName = config.ConfigName
	existingConfig.ConfigKey = config.ConfigKey
	existingConfig.ConfigValue = config.ConfigValue
	existingConfig.ConfigType = config.ConfigType
	existingConfig.Remark = config.Remark

	// 设置更新人和更新时间
	login_name := c.Ctx.Input.Header("login_name")
	existingConfig.UpdateBy = login_name

	// 更新到数据库
	err = configService.UpdateConfig(existingConfig)
	if err != nil {
		c.JSONErr(500, "更新系统配置失败: "+err.Error())
		return
	}

	c.JSONObject(map[string]interface{}{
		"success": true,
	})
}

// DeleteConfig 删除系统配置
func (c *ConfigController) DeleteConfig() {
	id := c.Ctx.Input.Param(":id")

	// 检查系统配置是否存在
	existingConfig, err := configService.GetConfigDetail(id)
	if err != nil {
		c.JSONErr(404, "系统配置不存在")
		return
	}

	login_name := c.Ctx.Input.Header("login_name")
	existingConfig.UpdateBy = login_name
	existingConfig.DeleteFlag = 1

	// 执行软删除
	err = configService.UpdateConfig(existingConfig)
	if err != nil {
		c.JSONErr(500, "删除系统配置失败: "+err.Error())
		return
	}

	c.JSONObject(map[string]interface{}{
		"success": true,
	})
}

// GetConfigDetail 获取系统配置详情
func (c *ConfigController) GetConfigDetail() {
	id := c.Ctx.Input.Param(":id")

	config, err := configService.GetConfigDetail(id)
	if err != nil {
		c.JSONErr(404, "系统配置不存在")
		return
	}

	c.JSONObject(config)
}
