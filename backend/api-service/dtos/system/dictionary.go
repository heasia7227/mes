package system

import system "common/models/system"

type DictionaryDto struct {
	system.Dictionary
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}
