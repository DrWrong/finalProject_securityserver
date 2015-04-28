package utils

import (
	"github.com/astaxie/beego/config"
)

var (
	IniConf config.ConfigContainer
)

func InitConfig(configFile string) {
	var err error
	IniConf, err = config.NewConfig("ini", configFile)
	if err != nil {
		panic(err)
	}
}
