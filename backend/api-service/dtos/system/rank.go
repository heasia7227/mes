package system

import system "common/models/system"

type RankDto struct {
	system.Rank
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}
