package system

import (
	systemDto "api-service/dtos/system"
	"common/utils/strings"
	"fmt"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type LogService struct {
}

func (s *LogService) GetLogList(params map[string]string) ([]systemDto.LogDto, int64) {
	var logs []systemDto.LogDto
	var total int64

	o := orm.NewOrm()

	// 构建 SQL 查询
	sql := `SELECT
				t1.id,
				t1.ip,
				t1.log_content,
				t1.is_exception,
				t1.operate_type,
				t1.operate_module,
				t1.operate_by,
				DATE_FORMAT(t1.operate_time, '%Y-%m-%d %H:%i:%s') as operate_time,
				t1.status,
				t1.duration,
				t1.description
			FROM t_sys_log t1
			WHERE 1=1 `

	var args []interface{}

	// 添加搜索条件
	if params["keyword"] != "" {
		sql += ` AND CONCAT(t1.ip, t1.log_content, t1.operate_type, t1.operate_module, t1.operate_by) LIKE ? `
		args = append(args, "%"+params["keyword"]+"%")
	}

	if params["status"] != "" {
		sql += ` AND t1.status = ?`
		args = append(args, params["status"])
	}

	sqlCount := fmt.Sprintf(`SELECT COUNT(1) AS total FROM (%s) a`, sql)
	// 查询总数
	err := o.Raw(sqlCount, args...).QueryRow(&total)
	if err != nil || total == 0 {
		return []systemDto.LogDto{}, 0
	}

	// 分页查询
	page_num := strings.ToInt(params["page_num"], 1)
	page_size := strings.ToInt(params["page_size"], 10)
	args = append(args, (page_num-1)*page_size, page_size)
	sql += ` ORDER BY t1.operate_time DESC LIMIT ?, ? `
	_, err = o.Raw(sql, args...).QueryRows(&logs)
	if err != nil {
		return []systemDto.LogDto{}, 0
	}
	return logs, total
}

func (s *LogService) GetLogDetail(id string) (systemDto.LogDto, error) {
	var log systemDto.LogDto
	o := orm.NewOrm()

	// 构建 SQL 查询
	sql := `SELECT
				t1.id,
				t1.ip,
				t1.log_content,
				t1.is_exception,
				t1.operate_type,
				t1.operate_module,
				t1.operate_by,
				DATE_FORMAT(t1.operate_time, '%Y-%m-%d %H:%i:%s') as operate_time,
				t1.status,
				t1.duration,
				t1.description
			FROM t_sys_log t1
			WHERE t1.id = ? `

	err := o.Raw(sql, id).QueryRow(&log)
	if err != nil {
		return systemDto.LogDto{}, err
	}
	return log, nil
}

func (s *LogService) CreateLog(logDto systemDto.LogDto) error {
	log := logDto.Log

	now := time.Now()
	log.OperateTime = now

	o := orm.NewOrm()
	_, err := o.Insert(&log)
	return err
}
