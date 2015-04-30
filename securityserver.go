package main

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/DrWrong/finalProject_securityserver/handler"
	"github.com/DrWrong/finalProject_securityserver/thrift_interface"
	"github.com/DrWrong/finalProject_securityserver/utils"
	log "github.com/Sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//  a security server class
//  it contains two attributes
//  + ExeDir the execute directory of the server
//  + server the thrift server
type SecurityServer struct {
	ExeDir string
	server *thrift.TSimpleServer
}

//  create a new thrift server and have it initialized
func NewSecurityServer(exeDir string) (Securityserver *SecurityServer) {
	Securityserver = &SecurityServer{ExeDir: exeDir}
	Securityserver.init()
	return Securityserver
}

// initialize the server and register a signal handler to handle terminate signal
func (self *SecurityServer) init() {
	self.registerSignalHandler()
	log.Info("security server has been inited")
}

//  when received the terminated signal stop the server
func (s *SecurityServer) registerSignalHandler() {
	go func() {
		for {
			c := make(chan os.Signal)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			//sig is blocked as c is no buffer
			sig := <-c
			log.Infof("Signal %d received", sig)
			s.stop()
			time.Sleep(time.Second)
			os.Exit(0)
		}
	}()
}

// run the server
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

// stop the server
func (s *SecurityServer) stop() {
	log.Info("security server is about to stop...")
	s.server.Stop()
	log.Info("security server has gone away")
}
