package system

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Rank struct {
	ID         string    `json:"id" orm:"column(id);pk"`
	RankName   string    `json:"rank_name" orm:"column(rank_name)"`
	RankStatus int       `json:"rank_status" orm:"column(rank_status)"`
	SortNo     int       `json:"sort_no" orm:"column(sort_no)"`
	Remark     string    `json:"remark" orm:"column(remark)"`
	CreateBy   string    `json:"create_by" orm:"column(create_by)"`
	CreateTime time.Time `json:"create_time" orm:"column(create_time)"`
	UpdateBy   string    `json:"update_by" orm:"column(update_by)"`
	UpdateTime time.Time `json:"update_time" orm:"column(update_time)"`
	DeleteFlag int       `json:"delete_flag" orm:"column(delete_flag)"`
}

func (c *Rank) TableName() string {
	return "t_sys_rank"
}

func init() {
	orm.RegisterModel(new(Rank))
}
