package system

import (
	"common/utils/strings"
	"fmt"
	"time"

	systemDto "api-service/dtos/system"

	"github.com/beego/beego/v2/client/orm"
)

type RoleService struct {
}

func (s *RoleService) GetRoleList(params map[string]string) ([]systemDto.RoleDto, int64) {
	var roles []systemDto.RoleDto
	var total int64

	o := orm.NewOrm()

	// 构建 SQL 查询
	sql := `SELECT
				t1.id,
				t1.title,
				t1.code,
				t1.description,
				t1.status,
				t1.create_by,
				DATE_FORMAT(t1.create_time, '%Y-%m-%d %H:%i:%s') as create_time,
				t1.update_by,
				DATE_FORMAT(t1.update_time, '%Y-%m-%d %H:%i:%s') as update_time
			FROM t_sys_role t1
			WHERE 1=1 `

	var args []interface{}

	// 添加搜索条件
	if params["keyword"] != "" {
		sql += ` AND CONCAT(t1.title, t1.code, t1.description) LIKE ? `
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
		return []systemDto.RoleDto{}, 0
	}

	// 分页查询
	page_num := strings.ToInt(params["page_num"], 1)
	page_size := strings.ToInt(params["page_size"], 10)
	args = append(args, (page_num-1)*page_size, page_size)
	sql += ` ORDER BY t1.create_time DESC LIMIT ?, ? `
	_, err = o.Raw(sql, args...).QueryRows(&roles)
	if err != nil {
		return []systemDto.RoleDto{}, 0
	}
	return roles, total
}

func (s *RoleService) GetRoleDetail(id string) (systemDto.RoleDto, error) {
	var role systemDto.RoleDto
	o := orm.NewOrm()

	// 构建 SQL 查询
	sql := `SELECT
				t1.id,
				t1.title,
				t1.code,
				t1.description,
				t1.status,
				t1.create_by,
				DATE_FORMAT(t1.create_time, '%Y-%m-%d %H:%i:%s') as create_time,
				t1.update_by,
				DATE_FORMAT(t1.update_time, '%Y-%m-%d %H:%i:%s') as update_time
			FROM t_sys_role t1
			WHERE t1.id = ? `

	err := o.Raw(sql, id).QueryRow(&role)
	if err != nil {
		return systemDto.RoleDto{}, err
	}
	return role, nil
}

func (s *RoleService) CreateRole(roleDto systemDto.RoleDto) error {
	role := roleDto.Role

	now := time.Now()
	role.CreateTime = now
	role.UpdateTime = now
	role.DeleteFlag = 0

	o := orm.NewOrm()
	_, err := o.Insert(&role)
	return err
}

func (s *RoleService) UpdateRole(roleDto systemDto.RoleDto) error {
	role := roleDto.Role

	now := time.Now()
	role.UpdateTime = now

	o := orm.NewOrm()
	_, err := o.Update(&role)
	return err
}
