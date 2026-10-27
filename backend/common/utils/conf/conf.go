package conf

import (
	"log"
	"path/filepath"
	"sync"

	"github.com/beego/beego/v2/core/config"
)

var (
	configer config.Configer
	readerMu sync.RWMutex
)

func initConfiger() {
	readerMu.Lock()
	defer readerMu.Unlock()

	if configer != nil {
		return
	}

	files := []string{
		filepath.Join("./conf", "app.conf"),
		filepath.Join("../conf", "app.conf"),
	}

	confPath := filepath.Dir(files[0])
	c, err := config.NewConfig("ini", files[0])
	if err != nil {
		log.Println("加载配置文件失败: ", err)
		c, err = config.NewConfig("ini", files[1])
		if err != nil {
			log.Println("加载配置文件失败: ", err)
		} else {
			log.Println("加载配置文件成功: ", files[1])
			confPath = filepath.Dir(files[1])
		}
	} else {
		log.Println("加载配置文件成功: ", files[0])
	}

	if c == nil {
		log.Println("配置文件加载失败")
		return
	}

	env, err := c.String("env")
	if err != nil {
		log.Println("读取env失败, 错误信息: ", err)
		env = "prod"
	}
	log.Println("当前环境: ", env)

	envConf := filepath.Join(confPath, "app."+env+".conf")
	log.Println("环境配置文件路径: ", envConf)

	c, err = config.NewConfig("ini", envConf)
	if err != nil {
		log.Println("加载环境配置文件失败: ", err)
		return
	}
	configer = c
	log.Println("加载环境配置文件成功: ", envConf)
}

// 读取配置文件 app.conf
func GetStringConf(confKey, def string) string {
	initConfiger()
	if configer == nil {
		return def
	}
	value, err := configer.String(confKey)
	if err != nil {
		return def
	}
	return value
}

// 读取配置文件 app.conf
func GetIntConf(confKey string, def int) int {
	initConfiger()
	if configer == nil {
		return def
	}
	value, err := configer.Int(confKey)
	if err != nil {
		return def
	}
	return value
}
