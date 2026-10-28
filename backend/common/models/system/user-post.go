package system

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type UserPost struct {
	ID         string    `json:"id" orm:"column(id);pk"`
	UserID     string    `json:"user_id" orm:"column(user_id)"`
	PostID     string    `json:"post_id" orm:"column(post_id)"`
	CreateBy   string    `json:"create_by" orm:"column(create_by)"`
	CreateTime time.Time `json:"create_time" orm:"column(create_time)"`
	UpdateBy   string    `json:"update_by" orm:"column(update_by)"`
	UpdateTime time.Time `json:"update_time" orm:"column(update_time)"`
	DeleteFlag int       `json:"delete_flag" orm:"column(delete_flag)"`
}

func (c *UserPost) TableName() string {
	return "t_sys_user_post"
}

func init() {
	orm.RegisterModel(new(UserPost))
}
