package system

import system "common/models/system"

type UserDto struct {
	system.User
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

type UserChangePasswordDto struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
