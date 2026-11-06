package system

import system "common/models/system"

type UserDto struct {
	system.User
	DeptName   string `json:"dept_name"`
	PostName   string `json:"post_name"`
	RankName   string `json:"rank_name"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

type UserChangePasswordDto struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type UserChangeStatusDto struct {
	Status string `json:"status"`
}
