package routers

import fileUpload "api-service/controllers/file-upload"

func init() {
	// 文件上传接口, 支持断点续传
	RegisterRouter("/file/upload/chunk", &fileUpload.UploadController{}, "post:UploadChunk")
	RegisterRouter("/file/upload/check", &fileUpload.UploadController{}, "get:CheckChunk")
	RegisterRouter("/file/upload/list", &fileUpload.UploadController{}, "get:GetUploadList")
}
