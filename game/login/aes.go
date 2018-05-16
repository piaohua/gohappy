package login

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"errors"

	"gohappy/glog"
)

func Decrypt(ciphertext, key []byte) ([]byte, error) {
	pkey := PaddingLeft(key, '0', 16)
	block, err := aes.NewCipher(pkey)
	if err != nil {
		return nil, err
	}
	blockModel := cipher.NewCBCDecrypter(block, pkey)
	plantText := make([]byte, len(ciphertext))
	blockModel.CryptBlocks(plantText, []byte(ciphertext))
	plantText, err = PKCS7UnPadding(plantText, block.BlockSize())
	if err != nil {
		return nil, err
	}
	return plantText, nil
}

func PKCS7UnPadding(plantText []byte, blockSize int) ([]byte, error) {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	if unpadding > length {
		return nil, errors.New("faild")
	}
	return plantText[:(length - unpadding)], nil
}

func PaddingLeft(ori []byte, pad byte, length int) []byte {
	if len(ori) >= length {
		return ori[:length]
	}
	pads := bytes.Repeat([]byte{pad}, length-len(ori))
	return append(pads, ori...)
}

func TouristAccount(account, key string) (string, error) {
	ciphertext, err1 := hex.DecodeString(account)
	if err1 != nil {
		return "", err1
	}
	b, err2 := Decrypt([]byte(ciphertext), []byte(key))
	if err2 != nil {
		return "", err2
	}
	return string(b), nil
}

//token

type TokenAesCFB struct {
	encStream cipher.Stream
	decStream cipher.Stream
}

var TokenBlock *TokenAesCFB

func TokenInit(password string) {
	key, iv := byteToKey(password)

	block, err := aes.NewCipher(key)
	if err != nil {
		glog.Panic(err)
	}
	encStream := cipher.NewCFBEncrypter(block, iv)
	decStream := cipher.NewCFBDecrypter(block, iv)
	TokenBlock = new(TokenAesCFB)
	TokenBlock.encStream = encStream
	TokenBlock.decStream = decStream
}

func byteToKey(password string) ([]byte, []byte) {
	pass := []byte(password)

	hash0 := []byte{}
	hash1 := md5.Sum(append(hash0, pass...))
	hash2 := md5.Sum(append(hash1[:], pass...))
	hash3 := md5.Sum(append(hash2[:], pass...))

	key := append(hash1[:], hash2[:]...)
	iv := hash3[:]

	return key, iv
}

//加密字符串
func (this *TokenAesCFB) Encrypt(plaintext []byte) (ciphertext []byte) {
	ciphertext = make([]byte, len(plaintext))

	this.encStream.XORKeyStream(ciphertext, plaintext)

	return
}

//解密字符串
func (this *TokenAesCFB) Decrypt(ciphertext []byte) (plaintext []byte) {
	plaintext = make([]byte, len(ciphertext))

	this.decStream.XORKeyStream(plaintext, ciphertext)

	return
}
