package main

import (
	"flag"
	"github.com/DrWrong/finalProject_securityserver/utils"
	log "github.com/Sirupsen/logrus"
	"os"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	pwd, _ := os.Getwd()
	exeDir := flag.String("d", pwd, "Execute Directory")
	flag.Parse()
	utils.InitConfig(*exeDir + "/conf/securityserver.cfg")
	server := NewSecurityServer(*exeDir)
	log.Info("every thing is inited now start the thrift server")
	server.run()
}
