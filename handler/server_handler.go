package handler

import (
	"errors"
	"github.com/DrWrong/finalProject_securityserver/models"
	"github.com/DrWrong/finalProject_securityserver/thrift_interface"
	"github.com/DrWrong/finalProject_securityserver/utils"
	log "github.com/Sirupsen/logrus"
)

//  a security server implementation to implement the pre-defined interface
//  it contains two attributes
// + Cipher the encrypt and decrypt model
// + key the pre-defined decrypt key.
type SecurityServerImpl struct {
	Cipher *models.Cipher
	Key    string
}

//  create a new  security server implementation instance
func NewSecurityServerImpl() *SecurityServerImpl {
	key := utils.IniConf.String("key")
	cipher, err := models.NewCipher(key)
	if err != nil {
		panic(err)
	}
	return &SecurityServerImpl{cipher, key}
}

//  a commonly log
func (h *SecurityServerImpl) log(commonRequest *thrift_interface.CommonRequest) {
	log.WithFields(log.Fields{
		"requester": commonRequest.Requester,
		"requestId": commonRequest.RequestId,
	}).Info("processing request")
}

// when process ping request, it just return true, with error = nil
func (h *SecurityServerImpl) Ping(
	commonRequest *thrift_interface.CommonRequest) (
	bool, error) {
	h.log(commonRequest)
	return true, nil
}

//  process encrypt it just invoke the Cipher instance's Encrypt method
func (h *SecurityServerImpl) Encrypted(
	commonRequest *thrift_interface.CommonRequest,
	plainText string) (string, error) {
	h.log(commonRequest)
	return h.Cipher.Encrypt(plainText)

}

//  decrypt process, it firstly compare the key with pre-configured
// keys if they are the same then decrypt
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
