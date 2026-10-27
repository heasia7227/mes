package db

import (
	"common/utils/conf"
	"common/utils/strings"
	"common/utils/wlogs"
	"fmt"
	"path/filepath"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	taskLogPath := conf.GetStringConf("logs.task_file_path", "logs/")
	logFilePath := filepath.Join(taskLogPath, "db", "db.log")
	logWriter := wlogs.NewOrmFileWriter(logFilePath, 5)
	// 替换ORM内部日志输出对象
	orm.DebugLog = orm.NewLog(logWriter)
	// 开启ORM调试模式, 否则不会输出SQL
	orm.Debug = true

	driver := conf.GetStringConf("db_driver", "mysql")

	host := conf.GetStringConf("db_host", "127.0.0.1")
	port := conf.GetStringConf("db_port", "3306")
	user := conf.GetStringConf("db_user", "root")
	pass := conf.GetStringConf("db_pass", "")
	name := conf.GetStringConf("db_name", "3d_auto_pro_line")
	charset := conf.GetStringConf("db_charset", "utf8mb4")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local", user, pass, host, port, name, charset)
	orm.DebugLog.Printf("连接数据库: %s", dsn)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	if err := orm.RegisterDataBase("default", driver, dsn); err != nil {
		orm.DebugLog.Printf("连接数据库失败: %s", err.Error())
	} else {
		orm.DebugLog.Printf("连接数据库成功")
	}

	// Configure connection pool
	maxOpenConns := web.AppConfig.DefaultInt("db_max_conn", 100)
	maxIdleConns := web.AppConfig.DefaultInt("db_max_idle", 50)
	orm.SetMaxOpenConns("default", maxOpenConns)
	orm.SetMaxIdleConns("default", maxIdleConns)
	orm.DebugLog.Printf("配置数据库连接池, 最大连接数: %s, 最大空闲连接数: %s", strings.ToString(maxOpenConns), strings.ToString(maxIdleConns))
}
