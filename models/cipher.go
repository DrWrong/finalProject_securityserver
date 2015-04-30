package models

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"io"
)

// Cipher Class it contains a cipher block attribute.
type Cipher struct {
	Block cipher.Block
}

// Create a new cipher instance
func NewCipher(key string) (*Cipher, error) {
	data := []byte(key)
	// md5.Size = 32
	sumdata := md5.Sum(data)
	Block, err := aes.NewCipher(sumdata[:md5.Size])
	if err != nil {
		return nil, err
	}
	return &Cipher{Block}, nil
}

// the encrypt function
//  it is a realization of AES encrypt algorithm
func (c *Cipher) Encrypt(plainText string) (string, error) {
	text := []byte(plainText)
	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	cfb := cipher.NewCFBEncrypter(c.Block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	// fmt.Printf("%0x", ciphertext)
	encodetext := base64.URLEncoding.EncodeToString(ciphertext)
	// fmt.Print(encodetext)
	return encodetext, nil
}

// the decrypt function
//  it is  a realization of AES decrypt algorithm
func (c *Cipher) Decrypt(cipherText string) (string, error) {
	var data []byte
	text, err := base64.URLEncoding.DecodeString(cipherText)
	if err != nil {
		log.Errorf("base64 decoding cipher error: %s", err)
		return "", err
	}
	if len(text) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(c.Block, iv)
	cfb.XORKeyStream(text, text)

	data, err = base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		log.Errorf("base64 decoding error: %s", err)
		return "", err
	}
	return fmt.Sprintf("%s", data), nil

}
