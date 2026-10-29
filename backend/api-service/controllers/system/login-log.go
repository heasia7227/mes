package system

import (
	"api-service/controllers"
	"api-service/services/system"
)

type LoginLogController struct {
	controllers.BaseController
}

var (
	loginLogService = &system.LoginLogService{}
)

// GetLoginLogList 获取登录日志列表
func (c *LoginLogController) GetLoginLogList() {
	params := c.GetGetParams()

	loginlogs, total := loginLogService.GetLoginLogList(params)

	c.JSONList(controllers.ListData{
		List: loginlogs,
		Pagination: controllers.PaginationData{
			Total:    total,
			PageNum:  params["page_num"],
			PageSize: params["page_size"],
		},
	})
}
