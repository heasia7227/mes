package system

import system "common/models/system"

type LoginLogDto struct {
	system.LoginLog
	LoginTime string `json:"login_time"`
}
