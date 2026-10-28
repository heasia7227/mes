package system

import (
	"api-service/controllers"
	systemDto "api-service/dtos/system"
	"api-service/services/system"
	"common/utils/strings"
)

type RankController struct {
	controllers.BaseController
}

var (
	rankService = &system.RankService{}
)

// GetRankList 获取岗位列表
func (c *RankController) GetRankList() {
	params := c.GetGetParams()

	ranks, total := rankService.GetRankList(params)

	c.JSONList(controllers.ListData{
		List: ranks,
		Pagination: controllers.PaginationData{
			Total:    total,
			PageNum:  params["page_num"],
			PageSize: params["page_size"],
		},
	})
}

// CreateRank 创建岗位
func (c *RankController) CreateRank() {
	var rank systemDto.RankDto
	if err := c.BindJSON(&rank); err != nil {
		c.JSONErr(400, err.Error())
		return
	}

	rank.ID = strings.GenerateID()
	rank.Status = "enable"

	// 设置创建人和更新人
	login_name := c.Ctx.Input.Header("login_name")
	rank.CreateBy = login_name
	rank.UpdateBy = login_name

	// 保存到数据库
	err := rankService.CreateRank(rank)
	if err != nil {
		c.JSONErr(500, "创建岗位失败: "+err.Error())
		return
	}

	c.JSONObject(map[string]interface{}{
		"success": true,
		"id":      rank.ID,
	})
}

// UpdateRank 更新岗位
func (c *RankController) UpdateRank() {
	id := c.Ctx.Input.Param(":id")

	// 检查岗位是否存在
	existingRank, err := rankService.GetRankDetail(id)
	if err != nil {
		c.JSONErr(404, "岗位不存在")
		return
	}

	var rank systemDto.RankDto
	if err := c.BindJSON(&rank); err != nil {
		c.JSONErr(400, err.Error())
		return
	}

	// 复制必要的字段, 保留原有ID和创建信息
	existingRank.RankName = rank.RankName
	existingRank.Sort = rank.Sort
	existingRank.Remark = rank.Remark

	// 设置更新人和更新时间
	login_name := c.Ctx.Input.Header("login_name")
	existingRank.UpdateBy = login_name

	// 更新到数据库
	err = rankService.UpdateRank(existingRank)
	if err != nil {
		c.JSONErr(500, "更新岗位失败: "+err.Error())
		return
	}

	c.JSONObject(map[string]interface{}{
		"success": true,
	})
}

// DeleteRank 删除岗位
func (c *RankController) DeleteRank() {
	id := c.Ctx.Input.Param(":id")

	// 检查岗位是否存在
	existingRank, err := rankService.GetRankDetail(id)
	if err != nil {
		c.JSONErr(404, "岗位不存在")
		return
	}
	if existingRank.Status == "enable" {
		c.JSONErr(400, "已启用的岗位不能删除")
		return
	}

	login_name := c.Ctx.Input.Header("login_name")
	existingRank.UpdateBy = login_name
	existingRank.DeleteFlag = 1

	// 执行软删除
	err = rankService.UpdateRank(existingRank)
	if err != nil {
		c.JSONErr(500, "删除岗位失败: "+err.Error())
		return
	}

	c.JSONObject(map[string]interface{}{
		"success": true,
	})
}

// GetRankDetail 获取岗位详情
func (c *RankController) GetRankDetail() {
	id := c.Ctx.Input.Param(":id")

	rank, err := rankService.GetRankDetail(id)
	if err != nil {
		c.JSONErr(404, "岗位不存在")
		return
	}

	c.JSONObject(rank)
}

// UpdateRankStatus 更新岗位状态
func (c *RankController) UpdateRankStatus() {
	id := c.Ctx.Input.Param(":id")

	// 检查岗位是否存在
	existingRank, err := rankService.GetRankDetail(id)
	if err != nil {
		c.JSONErr(404, "岗位不存在")
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
	existingRank.UpdateBy = login_name
	existingRank.Status = status

	// 更新状态
	err = rankService.UpdateRank(existingRank)
	if err != nil {
		c.JSONErr(500, "更新岗位状态失败: "+err.Error())
		return
	}

	c.JSONObject(map[string]interface{}{
		"success": true,
	})
}
