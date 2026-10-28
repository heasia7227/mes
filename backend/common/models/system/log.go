package system

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Log struct {
	ID            string    `json:"id" orm:"column(id);pk"`
	Ip            string    `json:"ip" orm:"column(ip)"`
	LogContent    string    `json:"log_content" orm:"column(log_content)"`
	IsException   string    `json:"is_exception" orm:"column(is_exception)"`
	OperateType   string    `json:"action" orm:"column(operate_type)"`
	OperateModule string    `json:"module" orm:"column(operate_module)"`
	OperateBy     string    `json:"operate_by" orm:"column(operate_by)"`
	OperateTime   time.Time `json:"operate_time" orm:"column(operate_time)"`
	Status        string    `json:"status" orm:"column(status)"`
	Duration      int       `json:"duration" orm:"column(duration)"`
	Description   string    `json:"description" orm:"column(description)"`
}

// TableName 鎸囧畾琛ㄥ悕
func (c *Log) TableName() string {
	return "t_sys_log"
}

func init() {
	orm.RegisterModel(new(Log))
}
