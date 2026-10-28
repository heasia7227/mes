package system

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Info struct {
	Id              string    `json:"id" orm:"column(id);pk)"`
	SystemName      string    `json:"system_name" orm:"column(system_name)"`
	SystemLogo      string    `json:"system_logo" orm:"column(system_logo)"`
	LoginBackground string    `json:"login_background" orm:"column(login_background)"`
	FooterText      string    `json:"footer_text" orm:"column(footer_text)"`
	IsEnabled       int       `json:"is_enabled" orm:"column(is_enabled)"`
	CreateBy        string    `json:"create_by" orm:"column(create_by)"`
	CreateTime      time.Time `json:"create_time" orm:"column(create_time)"`
	UpdateBy        string    `json:"update_by" orm:"column(update_by)"`
	UpdateTime      time.Time `json:"update_time" orm:"column(update_time)"`
	DeleteFlag      int       `json:"delete_flag" orm:"column(delete_flag)"`
}

func (c *Info) TableName() string {
	return "t_sys_info"
}

func init() {
	orm.RegisterModel(new(Info))
}
