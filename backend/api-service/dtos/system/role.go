package system

import system "common/models/system"

type RoleDto struct {
	system.Role
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}
