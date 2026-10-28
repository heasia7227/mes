package fileUpload

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

// FileUpload 文件上传记录
type FileUpload struct {
	ID         string    `json:"id" orm:"column(id);pk"`                // 文件ID
	FileName   string    `json:"file_name" orm:"column(file_name)"`     // 文件名称
	FileSize   int64     `json:"file_size" orm:"column(file_size)"`     // 文件大小(字节)
	FileFormat string    `json:"file_format" orm:"column(file_format)"` // 文件格式(VBF/ZIP等)
	FileMD5    string    `json:"file_md5" orm:"column(file_md5)"`       // 文件MD5
	FilePath   string    `json:"file_path" orm:"column(file_path)"`     // 文件路径
	CreateTime time.Time `json:"create_time" orm:"column(create_time)"` // 创建时间
}

// TableName 指定表名
func (c *FileUpload) TableName() string {
	return "t_upload_file"
}

func init() {
	orm.RegisterModel(new(FileUpload))
}
