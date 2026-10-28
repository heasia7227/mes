package system

import (
	"common/utils/strings"
	"fmt"
	"time"

	systemDto "api-service/dtos/system"

	"github.com/beego/beego/v2/client/orm"
)

type RankService struct {
}

func (s *RankService) GetRankList(params map[string]string) ([]systemDto.RankDto, int64) {
	var ranks []systemDto.RankDto
	var total int64

	o := orm.NewOrm()

	// 构建 SQL 查询
	sql := `SELECT
				t1.id,
				t1.rank_name,
				t1.status,
				t1.sort,
				t1.remark,
				t1.create_by,
				DATE_FORMAT(t1.create_time, '%Y-%m-%d %H:%i:%s') as create_time,
				t1.update_by,
				DATE_FORMAT(t1.update_time, '%Y-%m-%d %H:%i:%s') as update_time
			FROM t_sys_rank t1
			WHERE 1=1 `

	var args []interface{}

	// 添加搜索条件
	if params["keyword"] != "" {
		sql += ` AND CONCAT(t1.rank_name) LIKE ? `
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
		return []systemDto.RankDto{}, 0
	}

	// 分页查询
	page_num := strings.ToInt(params["page_num"], 1)
	page_size := strings.ToInt(params["page_size"], 10)
	args = append(args, (page_num-1)*page_size, page_size)
	sql += ` ORDER BY t1.create_time DESC LIMIT ?, ? `
	_, err = o.Raw(sql, args...).QueryRows(&ranks)
	if err != nil {
		return []systemDto.RankDto{}, 0
	}
	return ranks, total
}

func (s *RankService) GetRankDetail(id string) (systemDto.RankDto, error) {
	var rank systemDto.RankDto

	// 构建 SQL 查询
	sql := `SELECT
				t1.id,
				t1.rank_name,
				t1.status,
				t1.sort,
				t1.remark,
				t1.create_by,
				DATE_FORMAT(t1.create_time, '%Y-%m-%d %H:%i:%s') as create_time,
				t1.update_by,
				DATE_FORMAT(t1.update_time, '%Y-%m-%d %H:%i:%s') as update_time
			FROM t_sys_rank t1
			WHERE t1.id = ? `

	o := orm.NewOrm()
	err := o.Raw(sql, id).QueryRow(&rank)
	if err != nil {
		return systemDto.RankDto{}, err
	}
	return rank, nil
}

func (s *RankService) CreateRank(rankDto systemDto.RankDto) error {
	rank := rankDto.Rank

	now := time.Now()
	rank.CreateTime = now
	rank.UpdateTime = now
	rank.DeleteFlag = 0

	o := orm.NewOrm()
	_, err := o.Insert(&rank)
	return err
}

func (s *RankService) UpdateRank(rankDto systemDto.RankDto) error {
	rank := rankDto.Rank

	now := time.Now()
	rank.UpdateTime = now

	o := orm.NewOrm()
	_, err := o.Update(&rank)
	return err
}
