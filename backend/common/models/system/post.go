package system

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Post struct {
	ID         string    `json:"id" orm:"column(id);pk"`
	PostName   string    `json:"post_name" orm:"column(post_name)"`
	PostCode   string    `json:"post_code" orm:"column(post_code)"`
	PostStatus int       `json:"post_status" orm:"column(post_status)"`
	SortNo     int       `json:"sort_no" orm:"column(sort_no)"`
	Remark     string    `json:"remark" orm:"column(remark)"`
	CreateBy   string    `json:"create_by" orm:"column(create_by)"`
	CreateTime time.Time `json:"create_time" orm:"column(create_time)"`
	UpdateBy   string    `json:"update_by" orm:"column(update_by)"`
	UpdateTime time.Time `json:"update_time" orm:"column(update_time)"`
	DeleteFlag int       `json:"delete_flag" orm:"column(delete_flag)"`
}

func (c *Post) TableName() string {
	return "t_sys_post"
}

func init() {
	orm.RegisterModel(new(Post))
}
