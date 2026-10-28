package system

import (
	"api-service/controllers"
	systemDto "api-service/dtos/system"
	"api-service/services/system"
	"common/utils/strings"
)

type DeptController struct {
	controllers.BaseController
}

var (
	deptService = &system.DeptService{}
)

// GetDeptList 获取部门列表
func (c *DeptController) GetDeptList() {
	params := c.GetGetParams()

	depts, total := deptService.GetDeptList(params)

	c.JSONList(controllers.ListData{
		List: depts,
		Pagination: controllers.PaginationData{
			Total:    total,
			PageNum:  params["page_num"],
			PageSize: params["page_size"],
		},
	})
}

// CreateDept 创建部门
func (c *DeptController) CreateDept() {
	var dept systemDto.DeptDto
	if err := c.BindJSON(&dept); err != nil {
		c.JSONErr(400, err.Error())
		return
	}

	dept.ID = strings.GenerateID()
	dept.Status = "enable"

	// 设置创建人和更新人
	login_name := c.Ctx.Input.Header("login_name")
	dept.CreateBy = login_name
	dept.UpdateBy = login_name

	// 保存到数据库
	err := deptService.CreateDept(dept)
	if err != nil {
		c.JSONErr(500, "创建部门失败: "+err.Error())
		return
	}

	c.JSONObject(map[string]interface{}{
		"success": true,
		"id":      dept.ID,
	})
}

// UpdateDept 更新部门
func (c *DeptController) UpdateDept() {
	id := c.Ctx.Input.Param(":id")

	// 检查部门是否存在
	existingDept, err := deptService.GetDeptDetail(id)
	if err != nil {
		c.JSONErr(404, "部门不存在")
		return
	}

	var dept systemDto.DeptDto
	if err := c.BindJSON(&dept); err != nil {
		c.JSONErr(400, err.Error())
		return
	}

	// 复制必要的字段, 保留原有ID和创建信息
	existingDept.ParentID = dept.ParentID
	existingDept.DeptName = dept.DeptName
	existingDept.DeptCode = dept.DeptCode
	existingDept.DeptFullName = dept.DeptFullName
	existingDept.DeptType = dept.DeptType
	existingDept.Sort = dept.Sort
	existingDept.Remark = dept.Remark

	// 设置更新人和更新时间
	login_name := c.Ctx.Input.Header("login_name")
	existingDept.UpdateBy = login_name

	// 更新到数据库
	err = deptService.UpdateDept(existingDept)
	if err != nil {
		c.JSONErr(500, "更新部门失败: "+err.Error())
		return
	}

	c.JSONObject(map[string]interface{}{
		"success": true,
	})
}

// DeleteDept 删除部门
func (c *DeptController) DeleteDept() {
	id := c.Ctx.Input.Param(":id")

	// 检查部门是否存在
	existingDept, err := deptService.GetDeptDetail(id)
	if err != nil {
		c.JSONErr(404, "部门不存在")
		return
	}
	if existingDept.Status == "enable" {
		c.JSONErr(400, "已启用的部门不能删除")
		return
	}

	login_name := c.Ctx.Input.Header("login_name")
	existingDept.UpdateBy = login_name
	existingDept.DeleteFlag = 1

	// 执行软删除
	err = deptService.UpdateDept(existingDept)
	if err != nil {
		c.JSONErr(500, "删除部门失败: "+err.Error())
		return
	}

	c.JSONObject(map[string]interface{}{
		"success": true,
	})
}

// GetDeptDetail 获取部门详情
func (c *DeptController) GetDeptDetail() {
	id := c.Ctx.Input.Param(":id")

	dept, err := deptService.GetDeptDetail(id)
	if err != nil {
		c.JSONErr(404, "部门不存在")
		return
	}

	c.JSONObject(dept)
}

// UpdateDeptStatus 更新部门状态
func (c *DeptController) UpdateDeptStatus() {
	id := c.Ctx.Input.Param(":id")

	// 检查部门是否存在
	existingDept, err := deptService.GetDeptDetail(id)
	if err != nil {
		c.JSONErr(404, "部门不存在")
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
	existingDept.UpdateBy = login_name
	existingDept.Status = status

	// 更新状态
	err = deptService.UpdateDept(existingDept)
	if err != nil {
		c.JSONErr(500, "更新部门状态失败: "+err.Error())
		return
	}

	c.JSONObject(map[string]interface{}{
		"success": true,
	})
}
