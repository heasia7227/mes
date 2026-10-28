package system

import (
	systemDto "api-service/dtos/system"
	"common/models/system"
	"common/utils/strings"
	"common/utils/token"
	"fmt"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type UserService struct {
}

func (c *UserService) GetUserList(params map[string]string) ([]systemDto.UserDto, int64) {
	var users []systemDto.UserDto
	var total int64

	o := orm.NewOrm()

	// 构建 SQL 查询
	sql := `SELECT
				t1.id,
				t1.user_name,
				t1.login_name,
				t1.password,
				t1.sex,
				t1.post_code,
				t1.phone,
				t1.email,
				t1.avatar_url,
				t1.introduction,
				t1.status,
				t1.create_by,
				DATE_FORMAT(t1.create_time, '%Y-%m-%d %H:%i:%s') as create_time,
				t1.update_by,
				DATE_FORMAT(t1.update_time, '%Y-%m-%d %H:%i:%s') as update_time
			FROM t_sys_user t1
			WHERE 1=1 `

	var args []interface{}

	// 添加搜索条件
	if params["keyword"] != "" {
		sql += ` AND CONCAT(t1.user_name, t1.login_name, t1.phone, t1.email) LIKE ? `
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
		return []systemDto.UserDto{}, 0
	}

	// 分页查询
	page_num := strings.ToInt(params["page_num"], 1)
	page_size := strings.ToInt(params["page_size"], 10)
	args = append(args, (page_num-1)*page_size, page_size)
	sql += ` ORDER BY t1.create_time DESC LIMIT ?, ? `
	_, err = o.Raw(sql, args...).QueryRows(&users)
	if err != nil {
		return []systemDto.UserDto{}, 0
	}
	return users, total
}

func (c *UserService) GetUserDetail(id string) (systemDto.UserDto, error) {
	var user systemDto.UserDto
	o := orm.NewOrm()

	// 构建 SQL 查询
	sql := `SELECT
				t1.id,
				t1.user_name,
				t1.login_name,
				t1.password,
				t1.sex,
				t1.post_code,
				t1.phone,
				t1.email,
				t1.avatar_url,
				t1.introduction,
				t1.status,
				t1.create_by,
				DATE_FORMAT(t1.create_time, '%Y-%m-%d %H:%i:%s') as create_time,
				t1.update_by,
				DATE_FORMAT(t1.update_time, '%Y-%m-%d %H:%i:%s') as update_time
			FROM t_sys_user t1
			WHERE t1.id = ? `

	err := o.Raw(sql, id).QueryRow(&user)
	if err != nil {
		return systemDto.UserDto{}, err
	}
	return user, nil
}

func (c *UserService) CreateUser(userDto systemDto.UserDto) error {
	user := userDto.User

	now := time.Now()
	user.CreateTime = now
	user.UpdateTime = now
	user.DeleteFlag = 0

	// 密码加密处理
	user.Password = strings.MD5Hash(user.Password)

	o := orm.NewOrm()
	_, err := o.Insert(&user)
	return err
}

func (c *UserService) UpdateUser(userDto systemDto.UserDto) error {
	user := userDto.User

	now := time.Now()
	user.UpdateTime = now

	o := orm.NewOrm()
	_, err := o.Update(&user)
	return err
}

func (s *UserService) GetUserByToken(tokenStr string) *system.User {
	claims, err := token.ParseToken(tokenStr)
	if err != nil {
		return nil
	}
	return &system.User{
		ID:        claims.UserId,
		UserName:  claims.UserName,
		LoginName: claims.LoginName,
	}
}
