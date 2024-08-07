package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

// Encryption of keys
func Encrypt(secret []byte, data string) (string, error) {
	block, err := aes.NewCipher(secret)
	if err != nil {
		return "", err
	}
	b := []byte(data)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], b)
	return base64.URLEncoding.EncodeToString(ciphertext), nil
}
