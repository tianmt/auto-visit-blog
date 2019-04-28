package config

import (
	"auto-visit-blog/tools"
	"github.com/spf13/viper"
)

// 用户配置文件信息结构体
type ConfigInfo struct {
	Bname      string
	Uname      string
	Visitrate  float64
	Loopdelay  int
	Visitdelay int
}

// 读取配置文件
func (config_info *ConfigInfo) ReadConfigInfo() {
	err := Init("conf/config.yaml")
	tools.CheckError(err)

	blog_name := viper.GetString("blog")

	config_info.Bname = blog_name

	config_info.Uname = viper.GetString(blog_name + ".uname")
	config_info.Loopdelay = viper.GetInt(blog_name + ".loopdelay")
	config_info.Visitdelay = viper.GetInt(blog_name + ".visitdelay")
	config_info.Visitrate = viper.GetFloat64(blog_name + ".visitrate")
}
