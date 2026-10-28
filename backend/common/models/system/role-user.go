package system

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type RoleUser struct {
	ID         string    `json:"id" orm:"column(id);pk"`
	RoleID     string    `json:"role_id" orm:"column(role_id)"`
	UserID     string    `json:"user_id" orm:"column(user_id)"`
	CreateBy   string    `json:"create_by" orm:"column(create_by)"`
	CreateTime time.Time `json:"create_time" orm:"column(create_time)"`
}

// TableName 指定表名
func (ru *RoleUser) TableName() string {
	return "t_sys_role_user"
}

func init() {
	orm.RegisterModel(new(RoleUser))
}
