package fileUpload

import (
	fileUploadDto "api-service/dtos/file-upload"
	fileUpload "common/models/file-upload"
	"common/utils/strings"
	"fmt"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type FileUploadService struct {
}

// GetUploadList 获取文件上传记录列表
func (c *FileUploadService) GetUploadList(params map[string]string) ([]fileUploadDto.FileUploadDto, int64) {
	var uploads []fileUploadDto.FileUploadDto
	var total int64

	o := orm.NewOrm()

	// 构建 SQL 查询
	sql := `SELECT
				t1.id,
				t1.file_name,
				t1.file_size,
				t1.file_format,
				t1.file_md5,
				t1.file_path,
				DATE_FORMAT(t1.create_time, '%Y-%m-%d %H:%i:%s') as create_time
			FROM t_upload_file t1
			WHERE 1=1 `

	var args []interface{}

	// 添加搜索条件
	if params["file_name"] != "" {
		sql += ` AND t1.file_name LIKE ?`
		args = append(args, "%"+params["file_name"]+"%")
	}
	if params["file_md5"] != "" {
		sql += ` AND t1.file_md5 LIKE ?`
		args = append(args, "%"+params["file_md5"]+"%")
	}

	if params["start_date"] != "" && params["end_date"] != "" {
		sql += ` AND t1.create_time BETWEEN ? AND ?`
		args = append(args, fmt.Sprintf("%s 00:00:00", params["start_date"]), fmt.Sprintf("%s 23:59:59", params["end_date"]))
	} else if params["start_date"] != "" {
		sql += ` AND t1.create_time >= ?`
		args = append(args, fmt.Sprintf("%s 00:00:00", params["start_date"]))
	} else if params["end_date"] != "" {
		sql += ` AND t1.create_time <= ?`
		args = append(args, fmt.Sprintf("%s 23:59:59", params["end_date"]))
	}

	sqlCount := fmt.Sprintf(`SELECT COUNT(1) AS total FROM (%s) a`, sql)
	// 查询总数
	err := o.Raw(sqlCount, args...).QueryRow(&total)
	if err != nil || total == 0 {
		return []fileUploadDto.FileUploadDto{}, 0
	}

	// 分页查询
	page_num := strings.ToInt(params["page_num"], 1)
	page_size := strings.ToInt(params["page_size"], 10)
	args = append(args, (page_num-1)*page_size, page_size)
	sql += ` ORDER BY t1.create_time DESC LIMIT ?, ? `
	_, err = o.Raw(sql, args...).QueryRows(&uploads)
	if err != nil {
		return []fileUploadDto.FileUploadDto{}, 0
	}
	return uploads, total
}

// SaveFileUpload 保存文件上传记录
func (c *FileUploadService) SaveFileUpload(fileUpload *fileUpload.FileUpload) error {
	o := orm.NewOrm()

	fileUpload.ID = strings.GenerateID()
	fileUpload.CreateTime = time.Now()

	_, err := o.Insert(fileUpload)
	if err != nil {
		return err
	}

	return nil
}

// GetFileUploadByFileMd5 获取文件上传记录
func (c *FileUploadService) GetFileUploadByFileMd5(fileMd5 string) (*fileUpload.FileUpload, error) {
	var file fileUpload.FileUpload
	o := orm.NewOrm()
	err := o.QueryTable(new(fileUpload.FileUpload)).Filter("file_md5", fileMd5).Limit(1).One(&file)
	if err != nil {
		return &fileUpload.FileUpload{}, err
	}
	return &file, nil
}
