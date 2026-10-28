package system

import (
	"common/utils/strings"
	"fmt"
	"time"

	systemDto "api-service/dtos/system"

	"github.com/beego/beego/v2/client/orm"
)

type PostService struct {
}

func (s *PostService) GetPostList(params map[string]string) ([]systemDto.PostDto, int64) {
	var posts []systemDto.PostDto
	var total int64

	o := orm.NewOrm()

	// 构建 SQL 查询
	sql := `SELECT
				t1.id,
				t1.post_name,
				t1.post_code,
				t1.status,
				t1.sort,
				t1.remark,
				t1.create_by,
				DATE_FORMAT(t1.create_time, '%Y-%m-%d %H:%i:%s') as create_time,
				t1.update_by,
				DATE_FORMAT(t1.update_time, '%Y-%m-%d %H:%i:%s') as update_time
			FROM t_sys_post t1
			WHERE 1=1 `

	var args []interface{}

	// 添加搜索条件
	if params["keyword"] != "" {
		sql += ` AND CONCAT(t1.post_name, t1.post_code) LIKE ? `
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
		return []systemDto.PostDto{}, 0
	}

	// 分页查询
	page_num := strings.ToInt(params["page_num"], 1)
	page_size := strings.ToInt(params["page_size"], 10)
	args = append(args, (page_num-1)*page_size, page_size)
	sql += ` ORDER BY t1.create_time DESC LIMIT ?, ? `
	_, err = o.Raw(sql, args...).QueryRows(&posts)
	if err != nil {
		return []systemDto.PostDto{}, 0
	}
	return posts, total
}

func (s *PostService) GetPostDetail(id string) (systemDto.PostDto, error) {
	var post systemDto.PostDto

	// 构建 SQL 查询
	sql := `SELECT
				t1.id,
				t1.post_name,
				t1.post_code,
				t1.status,
				t1.sort,
				t1.remark,
				t1.create_by,
				DATE_FORMAT(t1.create_time, '%Y-%m-%d %H:%i:%s') as create_time,
				t1.update_by,
				DATE_FORMAT(t1.update_time, '%Y-%m-%d %H:%i:%s') as update_time
			FROM t_sys_post t1
			WHERE t1.id = ? `

	o := orm.NewOrm()
	err := o.Raw(sql, id).QueryRow(&post)
	if err != nil {
		return systemDto.PostDto{}, err
	}
	return post, nil
}

func (s *PostService) CreatePost(postDto systemDto.PostDto) error {
	post := postDto.Post

	now := time.Now()
	post.CreateTime = now
	post.UpdateTime = now
	post.DeleteFlag = 0

	o := orm.NewOrm()
	_, err := o.Insert(&post)
	return err
}

func (s *PostService) UpdatePost(postDto systemDto.PostDto) error {
	post := postDto.Post

	now := time.Now()
	post.UpdateTime = now

	o := orm.NewOrm()
	_, err := o.Update(&post)
	return err
}
