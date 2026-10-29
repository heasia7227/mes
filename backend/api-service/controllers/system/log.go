package system

import (
	"api-service/controllers"
	"api-service/services/system"
)

type LogController struct {
	controllers.BaseController
}

var (
	logService = &system.LogService{}
)

// GetLogList 获取日志列表
func (c *LogController) GetLogList() {
	params := c.GetGetParams()

	logs, total := logService.GetLogList(params)

	c.JSONList(controllers.ListData{
		List: logs,
		Pagination: controllers.PaginationData{
			Total:    total,
			PageNum:  params["page_num"],
			PageSize: params["page_size"],
		},
	})
}
