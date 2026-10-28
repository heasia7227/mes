package system

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Dictionary struct {
	ID                   string    `json:"id" orm:"column(id);pk"`
	CatalogID            string    `json:"catalog_id" orm:"column(catalog_id)"`
	CatalogCode          string    `json:"catalog_code" orm:"column(catalog_code)"`
	DictionaryValue      string    `json:"dictionary_value" orm:"column(dictionary_value)"`
	DictionaryName       string    `json:"dictionary_name" orm:"column(dictionary_name)"`
	Sort                 int       `json:"sort" orm:"column(sort)"`
	CreateBy             string    `json:"create_by" orm:"column(create_by)"`
	CreateTime           time.Time `json:"create_time" orm:"column(create_time)"`
	UpdateBy             string    `json:"update_by" orm:"column(update_by)"`
	UpdateTime           time.Time `json:"update_time" orm:"column(update_time)"`
	DeleteFlag           int       `json:"delete_flag" orm:"column(delete_flag)"`
	AssociationDict      string    `json:"association_dict" orm:"column(association_dict)"`
	AssociationDictName  string    `json:"association_dict_name" orm:"column(association_dict_name)"`
	AssociationDictValue string    `json:"association_dict_value" orm:"column(association_dict_value)"`
}

// TableName 指定表名
func (c *Dictionary) TableName() string {
	return "t_sys_dictionary"
}

func init() {
	orm.RegisterModel(new(Dictionary))
}
