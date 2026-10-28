package system

import system "common/models/system"

type LogDto struct {
	system.Log
	OperateTime string `json:"operate_time"`
}
