package system

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Menu struct {
	ID         string    `json:"id" orm:"column(id);pk"`
	Title      string    `json:"title" orm:"column(title)"`
	Icon       string    `json:"icon" orm:"column(icon)"`
	ParentID   string    `json:"parent_id" orm:"column(parent_id)"`
	Ranking    int       `json:"ranking" orm:"column(ranking)"`
	MenuType   string    `json:"menu_type" orm:"column(menu_type)"`
	DesignType string    `json:"design_type" orm:"column(design_type)"`
	MenuPath   string    `json:"menu_path" orm:"column(menu_path)"`
	Permission string    `json:"permission" orm:"column(permission)"`
	PageType   string    `json:"page_type" orm:"column(page_type)"`
	Status     string    `json:"status" orm:"column(status)"`
	CreateBy   string    `json:"create_by" orm:"column(create_by)"`
	CreateTime time.Time `json:"create_time" orm:"column(create_time)"`
	UpdateBy   string    `json:"update_by" orm:"column(update_by)"`
	UpdateTime time.Time `json:"update_time" orm:"column(update_time)"`
	DeleteFlag int       `json:"delete_flag" orm:"column(delete_flag)"`
}

// TableName 指定表名
func (c *Menu) TableName() string {
	return "t_sys_menu"
}

func init() {
	orm.RegisterModel(new(Menu))
}
