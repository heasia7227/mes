package system

import system "common/models/system"

type DeptDto struct {
	system.Dept
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}
