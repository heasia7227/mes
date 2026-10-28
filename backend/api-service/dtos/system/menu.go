package system

import system "common/models/system"

type MenuDto struct {
	system.Menu
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}
