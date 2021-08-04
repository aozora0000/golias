package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"time"
)

type Cipher struct {
	key []byte
}

func NewCipher(key []byte) Cipher {
	return Cipher{key: key}
}

func CreateNewKey(t time.Time) []byte {
	return decodeBase64(t.String())[0:31]
}

func encodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func decodeBase64(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

func (c Cipher) Encrypt(text []byte) (string, error) {
	block, err := aes.NewCipher(c.key)
	if err != nil {
		return "", err
	}
	b := encodeBase64(text)
	cipt := make([]byte, aes.BlockSize+len(b))
	iv := cipt[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(cipt[aes.BlockSize:], []byte(b))
	return encodeBase64(cipt), nil
}

func (c Cipher) Decrypt(t string) (string, error) {
	block, err := aes.NewCipher(c.key)
	if err != nil {
		return "", err
	}
	text := decodeBase64(t)
	if len(text) < aes.BlockSize {
		return "", errors.New("too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	return string(decodeBase64(string(text))), nil
}
