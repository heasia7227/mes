package system

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type LoginLog struct {
	ID            string    `json:"id" orm:"column(id);pk"`
	LoginName     string    `json:"login_name" orm:"column(login_name)"`
	LoginIP       string    `json:"login_ip" orm:"column(login_ip)"`
	LoginLocation string    `json:"login_location" orm:"column(login_location)"`
	Browser       string    `json:"browser" orm:"column(browser)"`
	OS            string    `json:"os" orm:"column(os)"`
	LoginStatus   int       `json:"login_status" orm:"column(login_status)"`
	Message       string    `json:"message" orm:"column(message)"`
	LoginTime     time.Time `json:"login_time" orm:"column(login_time)"`
	CreateBy      string    `json:"create_by" orm:"column(create_by)"`
	CreateTime    time.Time `json:"create_time" orm:"column(create_time)"`
	UpdateBy      string    `json:"update_by" orm:"column(update_by)"`
	UpdateTime    time.Time `json:"update_time" orm:"column(update_time)"`
	DeleteFlag    int       `json:"delete_flag" orm:"column(delete_flag)"`
}

func (c *LoginLog) TableName() string {
	return "t_sys_login_log"
}

func init() {
	orm.RegisterModel(new(LoginLog))
}
