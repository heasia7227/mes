package system

import (
	"common/utils/strings"
	"fmt"
	"time"

	systemDto "api-service/dtos/system"

	"github.com/beego/beego/v2/client/orm"
)

type DictionaryService struct {
}

func (s *DictionaryService) GetDictionaryList(params map[string]string) ([]systemDto.DictionaryDto, int64) {
	var dictionaries []systemDto.DictionaryDto
	var total int64

	o := orm.NewOrm()

	// 构建 SQL 查询
	sql := `SELECT
				t1.id,
				t1.catalog_id,
				t1.catalog_code,
				t1.dictionary_value,
				t1.dictionary_name,
				t1.sort,
				t1.association_dict,
				t1.association_dict_name,
				t1.association_dict_value,
				t1.create_by,
				DATE_FORMAT(t1.create_time, '%Y-%m-%d %H:%i:%s') as create_time,
				t1.update_by,
				DATE_FORMAT(t1.update_time, '%Y-%m-%d %H:%i:%s') as update_time
			FROM t_sys_dictionary t1
			WHERE 1=1 `

	var args []interface{}

	// 添加搜索条件
	if params["keyword"] != "" {
		sql += ` AND CONCAT(t1.dictionary_name, t1.dictionary_value) LIKE ? `
		args = append(args, "%"+params["keyword"]+"%")
	}

	sqlCount := fmt.Sprintf(`SELECT COUNT(1) AS total FROM (%s) a`, sql)
	// 查询总数
	err := o.Raw(sqlCount, args...).QueryRow(&total)
	if err != nil || total == 0 {
		return []systemDto.DictionaryDto{}, 0
	}

	// 分页查询
	page_num := strings.ToInt(params["page_num"], 1)
	page_size := strings.ToInt(params["page_size"], 10)
	args = append(args, (page_num-1)*page_size, page_size)
	sql += ` ORDER BY t1.create_time DESC LIMIT ?, ? `
	_, err = o.Raw(sql, args...).QueryRows(&dictionaries)
	if err != nil {
		return []systemDto.DictionaryDto{}, 0
	}
	return dictionaries, total
}

func (s *DictionaryService) GetDictionaryDetail(id string) (systemDto.DictionaryDto, error) {
	var dictionary systemDto.DictionaryDto

	// 构建 SQL 查询
	sql := `SELECT
				t1.id,
				t1.catalog_id,
				t1.catalog_code,
				t1.dictionary_value,
				t1.dictionary_name,
				t1.sort,
				t1.association_dict,
				t1.association_dict_name,
				t1.association_dict_value,
				t1.create_by,
				DATE_FORMAT(t1.create_time, '%Y-%m-%d %H:%i:%s') as create_time,
				t1.update_by,
				DATE_FORMAT(t1.update_time, '%Y-%m-%d %H:%i:%s') as update_time
			FROM t_sys_dictionary t1
			WHERE t1.id = ? `

	o := orm.NewOrm()
	err := o.Raw(sql, id).QueryRow(&dictionary)
	if err != nil {
		return systemDto.DictionaryDto{}, err
	}
	return dictionary, nil
}

func (s *DictionaryService) CreateDictionary(dictionaryDto systemDto.DictionaryDto) error {
	dictionary := dictionaryDto.Dictionary

	now := time.Now()
	dictionary.CreateTime = now
	dictionary.UpdateTime = now
	dictionary.DeleteFlag = 0

	o := orm.NewOrm()
	_, err := o.Insert(&dictionary)
	return err
}

func (s *DictionaryService) UpdateDictionary(dictionaryDto systemDto.DictionaryDto) error {
	dictionary := dictionaryDto.Dictionary

	now := time.Now()
	dictionary.UpdateTime = now

	o := orm.NewOrm()
	_, err := o.Update(&dictionary)
	return err
}
