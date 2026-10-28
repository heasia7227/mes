package system

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type UserRank struct {
	ID         string    `json:"id" orm:"column(id);pk"`
	UserID     string    `json:"user_id" orm:"column(user_id)"`
	RankID     string    `json:"rank_id" orm:"column(rank_id)"`
	CreateBy   string    `json:"create_by" orm:"column(create_by)"`
	CreateTime time.Time `json:"create_time" orm:"column(create_time)"`
	UpdateBy   string    `json:"update_by" orm:"column(update_by)"`
	UpdateTime time.Time `json:"update_time" orm:"column(update_time)"`
	DeleteFlag int       `json:"delete_flag" orm:"column(delete_flag)"`
}

func (ur *UserRank) TableName() string {
	return "t_sys_user_rank"
}

func init() {
	orm.RegisterModel(new(UserRank))
}
