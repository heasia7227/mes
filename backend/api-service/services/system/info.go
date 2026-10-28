package system

import (
	"time"

	systemDto "api-service/dtos/system"

	"github.com/beego/beego/v2/client/orm"
)

type InfoService struct {
}

func (s *InfoService) GetInfoList(params map[string]string) []systemDto.InfoDto {
	var infos []systemDto.InfoDto

	o := orm.NewOrm()

	// 构建 SQL 查询
	sql := `SELECT
				t1.id,
				t1.system_name,
				t1.system_logo,
				t1.login_background,
				t1.footer_text,
				t1.is_enabled,
				t1.create_by,
				DATE_FORMAT(t1.create_time, '%Y-%m-%d %H:%i:%s') as create_time,
				t1.update_by,
				DATE_FORMAT(t1.update_time, '%Y-%m-%d %H:%i:%s') as update_time
			FROM t_sys_info t1
			WHERE 1=1 `

	var args []interface{}

	_, err := o.Raw(sql, args...).QueryRows(&infos)
	if err != nil {
		return []systemDto.InfoDto{}
	}
	return infos
}

func (s *InfoService) UpdateInfo(infoDto systemDto.InfoDto) error {
	info := infoDto.Info

	now := time.Now()
	info.UpdateTime = now

	o := orm.NewOrm()
	_, err := o.Update(&info)
	return err
}
