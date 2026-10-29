package system

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type LoginLog struct {
	ID        string    `json:"id" orm:"column(id);pk"`
	LoginName string    `json:"login_name" orm:"column(login_name)"`
	LoginIP   string    `json:"login_ip" orm:"column(login_ip)"`
	Browser   string    `json:"browser" orm:"column(browser)"`
	OS        string    `json:"os" orm:"column(os)"`
	Status    string    `json:"status" orm:"column(status)"`
	Message   string    `json:"message" orm:"column(message)"`
	LoginTime time.Time `json:"login_time" orm:"column(login_time)"`
}

func (c *LoginLog) TableName() string {
	return "t_sys_login_log"
}

func init() {
	orm.RegisterModel(new(LoginLog))
}
