package system

import system "common/models/system"

type PostDto struct {
	system.Post
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}
