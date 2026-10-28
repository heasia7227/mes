package system

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Config struct {
	ID          string    `json:"id" orm:"column(id);pk"`
	ConfigName  string    `json:"config_name" orm:"column(config_name)"`
	ConfigKey   string    `json:"config_key" orm:"column(config_key)"`
	ConfigValue string    `json:"config_value" orm:"column(config_value)"`
	ConfigType  string    `json:"config_type" orm:"column(config_type)"`
	CreateBy    string    `json:"create_by" orm:"column(create_by)"`
	CreateTime  time.Time `json:"create_time" orm:"column(create_time)"`
	UpdateBy    string    `json:"update_by" orm:"column(update_by)"`
	UpdateTime  time.Time `json:"update_time" orm:"column(update_time)"`
	Remark      string    `json:"remark" orm:"column(remark)"`
	DeleteFlag  int       `json:"delete_flag" orm:"column(delete_flag)"`
}

func (c *Config) TableName() string {
	return "t_sys_config"
}

func init() {
	orm.RegisterModel(new(Config))
}
