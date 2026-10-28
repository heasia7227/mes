package system

import system "common/models/system"

type DictionaryCatalogDto struct {
	system.DictionaryCatalog
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}
