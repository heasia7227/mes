package system

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type RoleMenu struct {
	ID         string    `json:"id" orm:"column(id);pk"`
	RoleID     string    `json:"role_id" orm:"column(role_id)"`
	MenuID     string    `json:"menu_id" orm:"column(menu_id)"`
	CreateBy   string    `json:"create_by" orm:"column(create_by)"`
	CreateTime time.Time `json:"create_time" orm:"column(create_time)"`
}

// TableName 指定表名
func (c *RoleMenu) TableName() string {
	return "t_sys_role_menu"
}

func init() {
	orm.RegisterModel(new(RoleMenu))
}
