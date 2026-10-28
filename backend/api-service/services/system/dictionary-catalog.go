package system

import (
	"common/utils/strings"
	"fmt"
	"time"

	systemDto "api-service/dtos/system"

	"github.com/beego/beego/v2/client/orm"
)

type DictionaryCatalogService struct {
}

func (s *DictionaryCatalogService) GetDictionaryCatalogList(params map[string]string) ([]systemDto.DictionaryCatalogDto, int64) {
	var dictionaryCatalogs []systemDto.DictionaryCatalogDto
	var total int64

	o := orm.NewOrm()

	// 构建 SQL 查询
	sql := `SELECT
				t1.id,
				t1.catalog_name,
				t1.catalog_code,
				t1.parent_id,
				t1.create_by,
				DATE_FORMAT(t1.create_time, '%Y-%m-%d %H:%i:%s') as create_time,
				t1.update_by,
				DATE_FORMAT(t1.update_time, '%Y-%m-%d %H:%i:%s') as update_time
			FROM t_sys_dictionary_catalog t1
			WHERE 1=1 `

	var args []interface{}

	// 添加搜索条件
	if params["keyword"] != "" {
		sql += ` AND CONCAT(t1.catalog_name, t1.catalog_code) LIKE ? `
		args = append(args, "%"+params["keyword"]+"%")
	}

	sqlCount := fmt.Sprintf(`SELECT COUNT(1) AS total FROM (%s) a`, sql)
	// 查询总数
	err := o.Raw(sqlCount, args...).QueryRow(&total)
	if err != nil || total == 0 {
		return []systemDto.DictionaryCatalogDto{}, 0
	}

	// 分页查询
	page_num := strings.ToInt(params["page_num"], 1)
	page_size := strings.ToInt(params["page_size"], 10)
	args = append(args, (page_num-1)*page_size, page_size)
	sql += ` ORDER BY t1.create_time DESC LIMIT ?, ? `
	_, err = o.Raw(sql, args...).QueryRows(&dictionaryCatalogs)
	if err != nil {
		return []systemDto.DictionaryCatalogDto{}, 0
	}
	return dictionaryCatalogs, total
}

func (s *DictionaryCatalogService) GetDictionaryCatalogDetail(id string) (systemDto.DictionaryCatalogDto, error) {
	var dictionaryCatalog systemDto.DictionaryCatalogDto

	// 构建 SQL 查询
	sql := `SELECT
				t1.id,
				t1.catalog_name,
				t1.catalog_code,
				t1.parent_id,
				t1.create_by,
				DATE_FORMAT(t1.create_time, '%Y-%m-%d %H:%i:%s') as create_time,
				t1.update_by,
				DATE_FORMAT(t1.update_time, '%Y-%m-%d %H:%i:%s') as update_time
			FROM t_sys_dictionary_catalog t1
			WHERE t1.id = ? `

	o := orm.NewOrm()
	err := o.Raw(sql, id).QueryRow(&dictionaryCatalog)
	if err != nil {
		return systemDto.DictionaryCatalogDto{}, err
	}
	return dictionaryCatalog, nil
}

func (s *DictionaryCatalogService) CreateDictionaryCatalog(dictionaryCatalogDto systemDto.DictionaryCatalogDto) error {
	dictionaryCatalog := dictionaryCatalogDto.DictionaryCatalog

	now := time.Now()
	dictionaryCatalog.CreateTime = now
	dictionaryCatalog.UpdateTime = now
	dictionaryCatalog.DeleteFlag = 0

	o := orm.NewOrm()
	_, err := o.Insert(&dictionaryCatalog)
	return err
}

func (s *DictionaryCatalogService) UpdateDictionaryCatalog(dictionaryCatalogDto systemDto.DictionaryCatalogDto) error {
	dictionaryCatalog := dictionaryCatalogDto.DictionaryCatalog

	now := time.Now()
	dictionaryCatalog.UpdateTime = now

	o := orm.NewOrm()
	_, err := o.Update(&dictionaryCatalog)
	return err
}
