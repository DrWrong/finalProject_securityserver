package main

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	log "github.com/Sirupsen/logrus"
	"os"
	"os/signal"
	"securityserver/handler"
	"securityserver/thrift_interface"
	"securityserver/utils"
	"syscall"
	"time"
)

type SecurityServer struct {
	ExeDir string
	server *thrift.TSimpleServer
}

func NewSecurityServer(exeDir string) (Securityserver *SecurityServer) {
	Securityserver = &SecurityServer{ExeDir: exeDir}
	Securityserver.init()
	return Securityserver
}

func (self *SecurityServer) init() {
	self.registerSignalHandler()
	log.Info("security server has been inited")
}

func (s *SecurityServer) registerSignalHandler() {
	go func() {
		for {
			c := make(chan os.Signal)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			//sig is blocked as c is 没缓冲
			sig := <-c
			log.Infof("Signal %d received", sig)
			s.stop()
			time.Sleep(time.Second)
			os.Exit(0)
		}
	}()
}
func (s *SecurityServer) run() {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	networkAddr := fmt.Sprintf(":%s", utils.IniConf.String("serverport"))
	serverTransport, err := thrift.NewTServerSocket(networkAddr)
	if err != nil {
		log.Error("Error! %s", err)
		os.Exit(1)
	}
	handler := handler.NewSecurityServerImpl()
	processor := thrift_interface.NewSecureServiceProcessor(handler)
	s.server = thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	log.Infof("thrift server in %s", networkAddr)
	err = s.server.Serve()
	if err != nil {
		log.Errorf("Server start error: %s", err)
	}
}

func (s *SecurityServer) stop() {
	log.Info("security server is about to stop...")
	s.server.Stop()
	log.Info("security server has gone away")
}
