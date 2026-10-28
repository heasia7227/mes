package system

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type DictionaryCatalog struct {
	ID          string    `json:"id" orm:"column(id);pk"`
	CatalogName string    `json:"catalog_name" orm:"column(catalog_name)"`
	CatalogCode string    `json:"catalog_code" orm:"column(catalog_code)"`
	ParentID    string    `json:"parent_id" orm:"column(parent_id)"`
	CreateBy    string    `json:"create_by" orm:"column(create_by)"`
	CreateTime  time.Time `json:"create_time" orm:"column(create_time)"`
	UpdateBy    string    `json:"update_by" orm:"column(update_by)"`
	UpdateTime  time.Time `json:"update_time" orm:"column(update_time)"`
	DeleteFlag  int       `json:"delete_flag" orm:"column(delete_flag)"`
}

// TableName 指定表名
func (c *DictionaryCatalog) TableName() string {
	return "t_sys_dictionary_catalog"
}

func init() {
	orm.RegisterModel(new(DictionaryCatalog))
}
