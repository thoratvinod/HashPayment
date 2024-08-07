package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

// Decryption of data
func Decrypt(secret []byte, cryptoData string) (string, error) {
	block, err := aes.NewCipher(secret)
	if err != nil {
		return "", err
	}
	ciphertext, err := base64.URLEncoding.DecodeString(cryptoData)
	if err != nil {
		return "", err
	}
	if len(ciphertext) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)
	return string(ciphertext), nil
}
