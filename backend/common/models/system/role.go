package system

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Role struct {
	ID          string    `json:"id" orm:"column(id);pk"`
	Title       string    `json:"title" orm:"column(title)"`
	Code        string    `json:"code" orm:"column(code)"`
	Description string    `json:"description" orm:"column(description)"`
	Status      string    `json:"status" orm:"column(status)"`
	CreateBy    string    `json:"create_by" orm:"column(create_by)"`
	CreateTime  time.Time `json:"create_time" orm:"column(create_time)"`
	UpdateBy    string    `json:"update_by" orm:"column(update_by)"`
	UpdateTime  time.Time `json:"update_time" orm:"column(update_time)"`
	DeleteFlag  int       `json:"delete_flag" orm:"column(delete_flag)"`
}

func (c *Role) TableName() string {
	return "t_sys_role"
}
func init() {
	orm.RegisterModel(new(Role))
}
