package system

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type User struct {
	ID           string    `json:"id" orm:"column(id);pk"`
	UserName     string    `json:"user_name" orm:"column(user_name)"`
	LoginName    string    `json:"login_name" orm:"column(login_name)"`
	Password     string    `json:"password" orm:"column(password)"`
	PostCode     string    `json:"post_code" orm:"column(post_code)"`
	Sex          string    `json:"sex" orm:"column(sex)"`
	Phone        string    `json:"phone" orm:"column(phone)"`
	Email        string    `json:"email" orm:"column(email)"`
	AvatarUrl    string    `json:"avatar_url" orm:"column(avatar_url)"`
	Introduction string    `json:"introduction" orm:"column(introduction)"`
	Status       string    `json:"status" orm:"column(status)"`
	CreateBy     string    `json:"create_by" orm:"column(create_by)"`
	CreateTime   time.Time `json:"create_time" orm:"column(create_time)"`
	UpdateBy     string    `json:"update_by" orm:"column(update_by)"`
	UpdateTime   time.Time `json:"update_time" orm:"column(update_time)"`
	DeleteFlag   int       `json:"delete_flag" orm:"column(delete_flag)"`
}

// TableName 指定表名
func (c *User) TableName() string {
	return "t_sys_user"
}

func init() {
	orm.RegisterModel(new(User))
}
