package system

import (
	"common/utils/strings"
	"fmt"
	"time"

	systemDto "api-service/dtos/system"

	"github.com/beego/beego/v2/client/orm"
)

type DeptService struct {
}

func (s *DeptService) GetDeptList(params map[string]string) ([]systemDto.DeptDto, int64) {
	var depts []systemDto.DeptDto
	var total int64

	o := orm.NewOrm()

	// 构建 SQL 查询
	sql := `SELECT
				t1.id,
				t1.parent_id,
				t1.dept_name,
				t1.dept_code,
				t1.dept_full_name,
				t1.dept_type,
				t1.sort,
				t1.remark,
				t1.status,
				t1.id_path,
				t1.name_path,
				t1.create_by,
				DATE_FORMAT(t1.create_time, '%Y-%m-%d %H:%i:%s') as create_time,
				t1.update_by,
				DATE_FORMAT(t1.update_time, '%Y-%m-%d %H:%i:%s') as update_time
			FROM t_sys_dept t1
			WHERE 1=1 `

	var args []interface{}

	// 添加搜索条件
	if params["keyword"] != "" {
		sql += ` AND CONCAT(t1.dept_name, t1.dept_full_name, t1.dept_code) LIKE ? `
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
		return []systemDto.DeptDto{}, 0
	}

	// 分页查询
	page_num := strings.ToInt(params["page_num"], 1)
	page_size := strings.ToInt(params["page_size"], 10)
	args = append(args, (page_num-1)*page_size, page_size)
	sql += ` ORDER BY t1.create_time DESC LIMIT ?, ? `
	_, err = o.Raw(sql, args...).QueryRows(&depts)
	if err != nil {
		return []systemDto.DeptDto{}, 0
	}
	return depts, total
}

func (s *DeptService) GetDeptDetail(id string) (systemDto.DeptDto, error) {
	var dept systemDto.DeptDto

	// 构建 SQL 查询
	sql := `SELECT
				t1.id,
				t1.parent_id,
				t1.dept_name,
				t1.dept_code,
				t1.dept_full_name,
				t1.dept_type,
				t1.sort,
				t1.remark,
				t1.status,
				t1.id_path,
				t1.name_path,
				t1.create_by,
				DATE_FORMAT(t1.create_time, '%Y-%m-%d %H:%i:%s') as create_time,
				t1.update_by,
				DATE_FORMAT(t1.update_time, '%Y-%m-%d %H:%i:%s') as update_time
			FROM t_sys_dept t1
			WHERE t1.id = ? `

	o := orm.NewOrm()
	err := o.Raw(sql, id).QueryRow(&dept)
	if err != nil {
		return systemDto.DeptDto{}, err
	}
	return dept, nil
}

func (s *DeptService) CreateDept(deptDto systemDto.DeptDto) error {
	dept := deptDto.Dept

	now := time.Now()
	dept.CreateTime = now
	dept.UpdateTime = now
	dept.DeleteFlag = 0

	o := orm.NewOrm()
	_, err := o.Insert(&dept)
	return err
}

func (s *DeptService) UpdateDept(deptDto systemDto.DeptDto) error {
	dept := deptDto.Dept

	now := time.Now()
	dept.UpdateTime = now

	o := orm.NewOrm()
	_, err := o.Update(&dept)
	return err
}
