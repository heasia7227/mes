package system

import (
	systemDto "api-service/dtos/system"
	"common/utils/strings"
	"fmt"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type LoginLogService struct {
}

func (s *LoginLogService) GetLoginLogList(params map[string]string) ([]systemDto.LoginLogDto, int64) {
	var logs []systemDto.LoginLogDto
	var total int64

	o := orm.NewOrm()

	// 构建 SQL 查询
	sql := `SELECT
				t1.id,
				t1.login_name,
				t1.login_ip,
				t1.browser,
				t1.os,
				t1.login_status,
				t1.message,
				DATE_FORMAT(t1.login_time, '%Y-%m-%d %H:%i:%s') as login_time
			FROM t_sys_login_log t1
			WHERE 1=1 `

	var args []interface{}

	// 添加搜索条件
	if params["keyword"] != "" {
		sql += ` AND CONCAT(t1.login_name, t1.login_ip, t1.browser, t1.os) LIKE ? `
		args = append(args, "%"+params["keyword"]+"%")
	}

	if params["login_status"] != "" {
		sql += ` AND t1.login_status = ?`
		args = append(args, params["login_status"])
	}

	sqlCount := fmt.Sprintf(`SELECT COUNT(1) AS total FROM (%s) a`, sql)
	// 查询总数
	err := o.Raw(sqlCount, args...).QueryRow(&total)
	if err != nil || total == 0 {
		return []systemDto.LoginLogDto{}, 0
	}

	// 分页查询
	page_num := strings.ToInt(params["page_num"], 1)
	page_size := strings.ToInt(params["page_size"], 10)
	args = append(args, (page_num-1)*page_size, page_size)
	sql += ` ORDER BY t1.login_time DESC LIMIT ?, ? `
	_, err = o.Raw(sql, args...).QueryRows(&logs)
	if err != nil {
		return []systemDto.LoginLogDto{}, 0
	}
	return logs, total
}

func (s *LoginLogService) GetLoginLogDetail(id string) systemDto.LoginLogDto {
	var log systemDto.LoginLogDto
	// 构建 SQL 查询
	sql := `SELECT
				t1.id,
				t1.login_name,
				t1.login_ip,
				t1.browser,
				t1.os,
				t1.login_status,
				t1.message,
				DATE_FORMAT(t1.login_time, '%Y-%m-%d %H:%i:%s') as login_time
			FROM t_sys_login_log t1
			WHERE t1.id = ? `

	o := orm.NewOrm()
	err := o.Raw(sql, id).QueryRow(&log)
	if err != nil {
		return systemDto.LoginLogDto{}
	}
	return log
}

func (s *LoginLogService) CreateLoginLog(loginLogDto systemDto.LoginLogDto) error {
	log := loginLogDto.LoginLog

	now := time.Now()
	log.LoginTime = now

	o := orm.NewOrm()
	_, err := o.Insert(&log)
	return err
}
