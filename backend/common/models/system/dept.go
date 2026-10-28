package system

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Dept struct {
	ID           string    `json:"id" orm:"column(id);pk"`
	ParentID     *string   `json:"parent_id" orm:"column(parent_id)"`
	DeptName     string    `json:"dept_name" orm:"column(dept_name)"`
	DeptCode     string    `json:"dept_code" orm:"column(dept_code)"`
	DeptFullName string    `json:"dept_full_name" orm:"column(dept_full_name)"`
	DeptType     string    `json:"dept_type" orm:"column(dept_type)"`
	SortNo       int       `json:"sort_no" orm:"column(sort_no)"`
	Remark       string    `json:"remark" orm:"column(remark)"`
	Status       int       `json:"status" orm:"column(status)"`
	IDPath       string    `json:"id_path" orm:"column(id_path)"`
	NamePath     string    `json:"name_path" orm:"column(name_path)"`
	CreateBy     string    `json:"create_by" orm:"column(create_by)"`
	CreateTime   time.Time `json:"create_time" orm:"column(create_time)"`
	UpdateBy     string    `json:"update_by" orm:"column(update_by)"`
	UpdateTime   time.Time `json:"update_time" orm:"column(update_time)"`
	DeleteFlag   int       `json:"delete_flag" orm:"column(delete_flag)"`
}

func (c *Dept) TableName() string {
	return "t_sys_dept"
}

func init() {
	orm.RegisterModel(new(Dept))
}
