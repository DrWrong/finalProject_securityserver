package main

import (
	"flag"
	"github.com/DrWrong/finalProject_securityserver/utils"
	log "github.com/Sirupsen/logrus"
	"os"
	"runtime"
)

// the entrance of the server
// firstly it initialize the system's configure
// secondly it start the server.
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
