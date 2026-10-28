package fileUpload

import fileUpload "common/models/file-upload"

type FileUploadDto struct {
	fileUpload.FileUpload
	CreateTime string `json:"create_time"`
}
