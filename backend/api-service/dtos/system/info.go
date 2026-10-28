package system

import system "common/models/system"

type InfoDto struct {
	system.Info
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}
