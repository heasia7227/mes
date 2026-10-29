package system

import (
	"time"

	systemDto "api-service/dtos/system"

	"github.com/beego/beego/v2/client/orm"
)

type InfoService struct {
}

func (s *InfoService) GetInfoList() []systemDto.InfoDto {
	var infos []systemDto.InfoDto

	o := orm.NewOrm()

	// 构建 SQL 查询
	sql := `SELECT
				t1.id,
				t1.system_name,
				t1.system_logo,
				t1.login_background,
				t1.footer_text,
				t1.create_by,
				DATE_FORMAT(t1.create_time, '%Y-%m-%d %H:%i:%s') as create_time,
				t1.update_by,
				DATE_FORMAT(t1.update_time, '%Y-%m-%d %H:%i:%s') as update_time
			FROM t_sys_info t1`

	_, err := o.Raw(sql).QueryRows(&infos)
	if err != nil {
		return []systemDto.InfoDto{}
	}
	return infos
}

// GetInfoDetail 获取系统信息详情
func (s *InfoService) GetInfoDetail(id string) (systemDto.InfoDto, error) {
	var info systemDto.InfoDto

	o := orm.NewOrm()

	// 构建 SQL 查询
	sql := `SELECT
				t1.id,
				t1.system_name,
				t1.system_logo,
				t1.login_background,
				t1.footer_text,
				t1.create_by,
				DATE_FORMAT(t1.create_time, '%Y-%m-%d %H:%i:%s') as create_time,
				t1.update_by,
				DATE_FORMAT(t1.update_time, '%Y-%m-%d %H:%i:%s') as update_time
			FROM t_sys_info t1
			WHERE t1.id = ?`

	err := o.Raw(sql, id).QueryRow(&info)
	if err != nil {
		return systemDto.InfoDto{}, err
	}
	return info, nil
}

func (s *InfoService) UpdateInfo(infoDto systemDto.InfoDto) error {
	info := infoDto.Info

	now := time.Now()
	info.UpdateTime = now

	o := orm.NewOrm()
	_, err := o.Update(&info)
	return err
}
