package system

import system "common/models/system"

type ConfigDto struct {
	system.Config
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}
