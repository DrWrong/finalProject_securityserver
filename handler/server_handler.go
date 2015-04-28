package handler

import (
	"errors"
	log "github.com/Sirupsen/logrus"
	"securityserver/models"
	"securityserver/thrift_interface"
	"securityserver/utils"
)

type SecurityServerImpl struct {
	Cipher *models.Cipher
	Key    string
}

func NewSecurityServerImpl() *SecurityServerImpl {
	key := utils.IniConf.String("key")
	cipher, err := models.NewCipher(key)
	if err != nil {
		panic(err)
	}
	return &SecurityServerImpl{cipher, key}
}

func (h *SecurityServerImpl) log(commonRequest *thrift_interface.CommonRequest) {
	log.WithFields(log.Fields{
		"requester": commonRequest.Requester,
		"requestId": commonRequest.RequestId,
	}).Info("processing request")
}

func (h *SecurityServerImpl) Ping(
	commonRequest *thrift_interface.CommonRequest) (
	bool, error) {
	h.log(commonRequest)
	return true, nil
}

func (h *SecurityServerImpl) Encrypted(
	commonRequest *thrift_interface.CommonRequest,
	plainText string) (string, error) {
	h.log(commonRequest)
	return h.Cipher.Encrypt(plainText)

}

func (h *SecurityServerImpl) Decrypted(
	commonRequest *thrift_interface.CommonRequest,
	cipherText string,
	key string) (string, error) {
	if key != h.Key {
		err := errors.New("The key provided is wrong")
		return "", err
	}
	return h.Cipher.Decrypt(cipherText)
}
