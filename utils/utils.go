package utils

import (
	"github.com/astaxie/beego/config"
)

var (
	// a global configure container it contains the configure of the system.
	IniConf config.ConfigContainer
)

// read a configFile and return an IniConf instance
func InitConfig(configFile string) {
	var err error
	IniConf, err = config.NewConfig("ini", configFile)
	if err != nil {
		panic(err)
	}
}
