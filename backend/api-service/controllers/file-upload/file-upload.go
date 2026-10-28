package fileUpload

import (
	"api-service/controllers"
	fileUpload "api-service/services/file-upload"
	fileUploadModel "common/models/file-upload"
	"common/utils/conf"
	"common/utils/goroutine"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/beego/beego/v2/core/logs"
)

// UploadController 文件上传控制器
type UploadController struct {
	controllers.BaseController
}

var (
	fileUploadService = &fileUpload.FileUploadService{}
)

// UploadRequest 上传请求参数
type UploadRequest struct {
	ChunkNumber int    `json:"chunkNumber" form:"chunkNumber"`
	TotalChunks int    `json:"totalChunks" form:"totalChunks"`
	ChunkSize   int64  `json:"chunkSize" form:"chunkSize"`
	TotalSize   int64  `json:"totalSize" form:"totalSize"`
	FileName    string `json:"fileName" form:"fileName"`
	FileMd5     string `json:"fileMd5" form:"fileMd5"` // 文件MD5值
}

func (c *UploadController) GetUploadList() {
	params := c.GetGetParams()
	uploads, total := fileUploadService.GetUploadList(params)

	c.JSONList(controllers.ListData{
		List: uploads,
		Pagination: controllers.PaginationData{
			Total:    total,
			PageNum:  params["page_num"],
			PageSize: params["page_size"],
		},
	})
}

// UploadChunk 上传文件块
func (c *UploadController) UploadChunk() {
	var req UploadRequest
	if err := c.ParseForm(&req); err != nil {
		c.JSONErr(400, "参数解析失败: "+err.Error())
		return
	}

	// 获取文件
	file, _, err := c.GetFile("file")
	if err != nil {
		c.JSONErr(400, "获取文件失败: "+err.Error())
		return
	}
	defer file.Close()

	// 获取配置的文件保存路径
	processFilePath := conf.GetStringConf("process_host_file_path", "./process")
	if processFilePath == "" {
		c.JSONErr(500, "未配置文件保存路径")
		return
	}

	year := time.Now().Format("2006")
	month := time.Now().Format("01")
	day := time.Now().Format("02")
	relativePath := filepath.Join(year, month, day)
	saveDir := filepath.Join(processFilePath, relativePath)
	tempDir := filepath.Join(saveDir, "temp")

	// 确保目录存在
	if err := os.MkdirAll(tempDir, 0777); err != nil {
		c.JSONErr(500, "创建临时目录失败: "+err.Error())
		return
	}

	// 保存文件块
	chunkFilePath := filepath.Join(tempDir, fmt.Sprintf("chunk_%d", req.ChunkNumber))
	if err := c.SaveToFile("file", chunkFilePath); err != nil {
		c.JSONErr(500, "保存文件块失败: "+err.Error())
		return
	}

	// 检查是否所有块都已上传
	if isAllChunksUploaded(tempDir, req.TotalChunks) {
		// 合并文件块
		ext := filepath.Ext(req.FileName)
		finalName := req.FileMd5 + ext
		targetFilePath := filepath.Join(saveDir, finalName)
		if err := mergeChunks(tempDir, targetFilePath, req.TotalChunks); err != nil {
			c.JSONErr(500, "合并文件块失败: "+err.Error())
			return
		}

		// 清理临时文件
		goroutine.SafeGo(func() {
			if err := os.RemoveAll(tempDir); err != nil {
				logs.Error("清理临时文件失败:", err)
			}
		})

		// 保存文件上传记录
		err := fileUploadService.SaveFileUpload(&fileUploadModel.FileUpload{
			FileName:   req.FileName,
			FileSize:   req.TotalSize,
			FileFormat: ext[1:],
			FileMD5:    req.FileMd5,
			FilePath:   filepath.Join(relativePath, finalName),
		})
		if err != nil {
			c.JSONErr(500, "保存文件上传记录失败: "+err.Error())
			return
		}

		c.JSONObject(map[string]interface{}{
			"success":  true,
			"filePath": filepath.Join(relativePath, finalName),
			"fileMd5":  req.FileMd5,
			"fileUrl":  targetFilePath,
		})
	} else {
		c.JSONObject(map[string]interface{}{
			"success":     true,
			"chunkNumber": req.ChunkNumber,
			"message":     "文件块上传成功, 等待其他块",
		})
	}
}

// CheckChunk 检查文件块是否已存在
func (c *UploadController) CheckChunk() {
	params := c.GetGetParams()
	if params["fileMd5"] == "" {
		c.JSONErr(400, "参数fileMd5不能为空")
		return
	}

	fileMd5 := params["fileMd5"]
	file, _ := fileUploadService.GetFileUploadByFileMd5(fileMd5)

	if file == nil || file.ID == "" {
		c.JSONObject(map[string]interface{}{
			"uploaded": false,
		})
		return
	}

	processFilePath := conf.GetStringConf("process_host_file_path", "./process")
	c.JSONObject(map[string]interface{}{
		"uploaded": true,
		"id":       file.ID,
		"filePath": file.FilePath,
		"fileMd5":  file.FileMD5,
		"fileUrl":  processFilePath + file.FilePath,
	})
}

// isAllChunksUploaded 检查是否所有文件块都已上传
func isAllChunksUploaded(tempDir string, totalChunks int) bool {
	files, err := os.ReadDir(tempDir)
	if err != nil {
		return false
	}

	chunkCount := 0
	fileSet := make(map[string]struct{}, len(files))
	for _, file := range files {
		fileSet[file.Name()] = struct{}{}
	}
	prefix := "chunk_"
	for i := 1; i <= totalChunks; i++ {
		name := prefix + strconv.Itoa(i)
		if _, exists := fileSet[name]; exists {
			chunkCount++
		}
	}

	return chunkCount == totalChunks
}

// mergeChunks 合并文件块
func mergeChunks(tempDir, targetFilePath string, totalChunks int) error {
	// 确保目标文件的目录存在
	if err := os.MkdirAll(filepath.Dir(targetFilePath), 0777); err != nil {
		return err
	}

	// 创建目标文件
	targetFile, err := os.Create(targetFilePath)
	if err != nil {
		return err
	}
	defer targetFile.Close()

	// 按顺序合并文件块
	for i := 1; i <= totalChunks; i++ {
		chunkPath := filepath.Join(tempDir, fmt.Sprintf("chunk_%d", i))
		chunkFile, err := os.Open(chunkPath)
		if err != nil {
			return err
		}

		// 复制文件块内容到目标文件
		_, err = io.Copy(targetFile, chunkFile)
		chunkFile.Close()
		if err != nil {
			return err
		}
	}

	return nil
}
