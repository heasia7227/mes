package system

import (
	"api-service/controllers"
	systemDto "api-service/dtos/system"
	"api-service/services/system"
)

type InfoController struct {
	controllers.BaseController
}

var (
	infoService = &system.InfoService{}
)

// GetSystemInfo 获取系统信息
func (c *InfoController) GetSystemInfo() {
	infos := infoService.GetInfoList()

	if len(infos) == 0 {
		c.JSONErr(404, "系统信息不存在")
		return
	}

	c.JSONObject(infos[0])
}

// UpdateSystemInfo 更新系统信息
func (c *InfoController) UpdateSystemInfo() {
	id := c.Ctx.Input.Param(":id")

	// 检查系统信息是否存在
	existingInfo, err := infoService.GetInfoDetail(id)
	if err != nil {
		c.JSONErr(404, "系统信息不存在")
		return
	}

	var info systemDto.InfoDto
	if err := c.BindJSON(&info); err != nil {
		c.JSONErr(400, err.Error())
		return
	}

	// 复制必要的字段, 保留原有ID和创建信息
	existingInfo.SystemName = info.SystemName
	existingInfo.SystemLogo = info.SystemLogo
	existingInfo.LoginBackground = info.LoginBackground
	existingInfo.FooterText = info.FooterText

	// 设置更新人和更新时间
	login_name := c.Ctx.Input.Header("login_name")
	existingInfo.UpdateBy = login_name

	// 更新到数据库
	err = infoService.UpdateInfo(existingInfo)
	if err != nil {
		c.JSONErr(500, "更新系统信息失败: "+err.Error())
		return
	}

	c.JSONObject(map[string]interface{}{
		"success": true,
	})
}
