package system

import (
	"common/utils/strings"
	"fmt"
	"time"

	systemDto "api-service/dtos/system"

	"github.com/beego/beego/v2/client/orm"
)

type MenuService struct {
}

func (s *MenuService) GetMenuList(params map[string]string) ([]systemDto.MenuDto, int64) {
	var menus []systemDto.MenuDto
	var total int64

	o := orm.NewOrm()

	// 构建 SQL 查询
	sql := `SELECT
				t1.id,
				t1.title,
				t1.icon,
				t1.parent_id,
				t1.ranking,
				t1.menu_type,
				t1.design_type,
				t1.menu_path,
				t1.page_type,
				t1.permission,
				t1.status,
				t1.create_by,
				DATE_FORMAT(t1.create_time, '%Y-%m-%d %H:%i:%s') as create_time,
				t1.update_by,
				DATE_FORMAT(t1.update_time, '%Y-%m-%d %H:%i:%s') as update_time
			FROM t_sys_menu t1
			WHERE 1=1 `

	var args []interface{}

	// 添加搜索条件
	if params["keyword"] != "" {
		sql += ` AND CONCAT(t1.title) LIKE ? `
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
		return []systemDto.MenuDto{}, 0
	}

	// 分页查询
	page_num := strings.ToInt(params["page_num"], 1)
	page_size := strings.ToInt(params["page_size"], 10)
	args = append(args, (page_num-1)*page_size, page_size)
	sql += ` ORDER BY t1.create_time DESC LIMIT ?, ? `
	_, err = o.Raw(sql, args...).QueryRows(&menus)
	if err != nil {
		return []systemDto.MenuDto{}, 0
	}
	return menus, total
}

func (s *MenuService) GetMenuDetail(id string) (systemDto.MenuDto, error) {
	var menu systemDto.MenuDto

	// 构建 SQL 查询
	sql := `SELECT
				t1.id,
				t1.title,
				t1.icon,
				t1.parent_id,
				t1.ranking,
				t1.menu_type,
				t1.design_type,
				t1.menu_path,
				t1.page_type,
				t1.permission,
				t1.status,
				t1.create_by,
				DATE_FORMAT(t1.create_time, '%Y-%m-%d %H:%i:%s') as create_time,
				t1.update_by,
				DATE_FORMAT(t1.update_time, '%Y-%m-%d %H:%i:%s') as update_time
			FROM t_sys_menu t1
			WHERE t1.id = ? `

	o := orm.NewOrm()
	err := o.Raw(sql, id).QueryRow(&menu)
	if err != nil {
		return systemDto.MenuDto{}, err
	}
	return menu, nil
}

func (s *MenuService) CreateMenu(menuDto systemDto.MenuDto) error {
	menu := menuDto.Menu

	now := time.Now()
	menu.CreateTime = now
	menu.UpdateTime = now
	menu.DeleteFlag = 0

	o := orm.NewOrm()
	_, err := o.Insert(&menu)
	return err
}

func (s *MenuService) UpdateMenu(menuDto systemDto.MenuDto) error {
	menu := menuDto.Menu

	now := time.Now()
	menu.UpdateTime = now

	o := orm.NewOrm()
	_, err := o.Update(&menu)
	return err
}
