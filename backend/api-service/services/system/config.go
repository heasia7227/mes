package system

import (
	"common/utils/strings"
	"fmt"
	"time"

	systemDto "api-service/dtos/system"

	"github.com/beego/beego/v2/client/orm"
)

type ConfigService struct {
}

func (s *ConfigService) GetConfigList(params map[string]string) ([]systemDto.ConfigDto, int64) {
	var configs []systemDto.ConfigDto
	var total int64

	o := orm.NewOrm()

	// 构建 SQL 查询
	sql := `SELECT
				t1.id,
				t1.config_name,
				t1.config_key,
				t1.config_value,
				t1.config_type,
				t1.remark,
				t1.create_by,
				DATE_FORMAT(t1.create_time, '%Y-%m-%d %H:%i:%s') as create_time,
				t1.update_by,
				DATE_FORMAT(t1.update_time, '%Y-%m-%d %H:%i:%s') as update_time
			FROM t_sys_config t1
			WHERE 1=1 `

	var args []interface{}

	// 添加搜索条件
	if params["keyword"] != "" {
		sql += ` AND CONCAT(t1.config_name, t1.config_key, t1.config_value) LIKE ? `
		args = append(args, "%"+params["keyword"]+"%")
	}

	sqlCount := fmt.Sprintf(`SELECT COUNT(1) AS total FROM (%s) a`, sql)
	// 查询总数
	err := o.Raw(sqlCount, args...).QueryRow(&total)
	if err != nil || total == 0 {
		return []systemDto.ConfigDto{}, 0
	}

	// 分页查询
	page_num := strings.ToInt(params["page_num"], 1)
	page_size := strings.ToInt(params["page_size"], 10)
	args = append(args, (page_num-1)*page_size, page_size)
	sql += ` ORDER BY t1.create_time DESC LIMIT ?, ? `
	_, err = o.Raw(sql, args...).QueryRows(&configs)
	if err != nil {
		return []systemDto.ConfigDto{}, 0
	}
	return configs, total
}

func (s *ConfigService) GetConfigDetail(id string) (systemDto.ConfigDto, error) {
	var config systemDto.ConfigDto

	// 构建 SQL 查询
	sql := `SELECT
				t1.id,
				t1.config_name,
				t1.config_key,
				t1.config_value,
				t1.config_type,
				t1.remark,
				t1.create_by,
				DATE_FORMAT(t1.create_time, '%Y-%m-%d %H:%i:%s') as create_time,
				t1.update_by,
				DATE_FORMAT(t1.update_time, '%Y-%m-%d %H:%i:%s') as update_time
			FROM t_sys_config t1
			WHERE t1.id = ? `

	o := orm.NewOrm()
	err := o.Raw(sql, id).QueryRow(&config)
	if err != nil {
		return systemDto.ConfigDto{}, err
	}
	return config, nil
}

func (s *ConfigService) CreateConfig(configDto systemDto.ConfigDto) error {
	config := configDto.Config

	now := time.Now()
	config.CreateTime = now
	config.UpdateTime = now
	config.DeleteFlag = 0

	o := orm.NewOrm()
	_, err := o.Insert(&config)
	return err
}

func (s *ConfigService) UpdateConfig(configDto systemDto.ConfigDto) error {
	config := configDto.Config

	now := time.Now()
	config.UpdateTime = now

	o := orm.NewOrm()
	_, err := o.Update(&config)
	return err
}
