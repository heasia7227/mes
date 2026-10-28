package system

import system "common/models/system"

type LoginLogDto struct {
	system.LoginLog
	LoginTime  string `json:"login_time"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}
