package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"

	"github.com/thoratvinod/HashPayment/utils"
)

// Encrypt function to encrypt data with the given secret
func Encrypt(secret []byte, data string) (string, error) {
	// Create a new AES cipher with the provided secret key
	block, err := aes.NewCipher(secret)
	if err != nil {
		return "", err
	}

	// Convert data to byte array
	b := []byte(data)

	// Create a byte array to hold the ciphertext (IV + actual data)
	ciphertext := make([]byte, aes.BlockSize+len(b))

	// Generate a random Initialization Vector (IV) for the AES encryption
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// Create a new CFB encrypter
	stream := cipher.NewCFBEncrypter(block, iv)

	// Encrypt the data and store it in the ciphertext array
	stream.XORKeyStream(ciphertext[aes.BlockSize:], b)

	// Encode the ciphertext to a Base64 string and return it
	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// Decrypt function to decrypt data with the given secret
func Decrypt(secret []byte, cryptoData string) (string, error) {
	// Create a new AES cipher with the provided secret key
	block, err := aes.NewCipher(secret)
	if err != nil {
		return "", err
	}

	// Decode the Base64 encoded ciphertext
	ciphertext, err := base64.URLEncoding.DecodeString(cryptoData)
	if err != nil {
		return "", err
	}

	// Check if the ciphertext length is valid
	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	// Extract the Initialization Vector (IV) from the ciphertext
	iv := ciphertext[:aes.BlockSize]

	// Extract the actual encrypted data
	ciphertext = ciphertext[aes.BlockSize:]

	// Create a new CFB decrypter
	stream := cipher.NewCFBDecrypter(block, iv)

	// Decrypt the data in place
	stream.XORKeyStream(ciphertext, ciphertext)

	// Convert the decrypted data to string and return it
	return string(ciphertext), nil
}

func main() {
	// Hex encoded secret key
	hexSecretKey := "b210e6ae7920d010dc7ea36fda78be24f3fc81adb81db87ffa3324b1e4ea1538"

	// Decode the hex secret key to bytes
	secretKey, err := hex.DecodeString(hexSecretKey)
	if err != nil {
		fmt.Println("Hex decode error:", err)
		return
	}

	// Ensure the secret key length is correct for AES-256 (32 bytes)
	if len(secretKey) != 32 {
		fmt.Println("Invalid key size:", len(secretKey))
		return
	}

	// Data to encrypt
	data := "AQExhmfxK4zObxxLw0m/n3Q5qf3Vb4ZMBJ9rW2ZZ03a/zTUeL2Vi2ZEeWzsTT2G96p8q+xDBXVsNvuR83LVYjEgiTGAH-aUNKkHPjul/Z9yhWTLHDIYTkrTSI922rhW+UDuZoarM=-i1i7>$S98tZkhgw$g{F"

	// Encrypt the data
	// encryptedData, err := Encrypt(secretKey, data)
	// encryptedData, err := utils.Encrypt(secretKey, data)
	// if err != nil {
	// 	fmt.Println("Encryption error:", err)
	// 	return
	// }
	// fmt.Println("Encrypted data:", encryptedData)


	encryptedData := "vDsdBmd4olqB_lTfpiRYgksxNq9nHxwzXXpWZfZv__xSiiJvRDL3kcoOZRJMJYHKJS5Y2KXuqwuRVYr-p41N6T7xgapR-Cfhg1TfndLjATknEVzDQYqIo-JHCQyk9Gabpr1MTIJUh8F8dpPFNVvlQLp4tcQxMFix-VuS2h8_2IscWvVSEQCzOqM5XWL6sgMWxOfPk-gheTqMyMh0HKSOVIYNciuNKiUqmFEK3LI="
	// Decrypt the data
	decryptedData, err := utils.Decrypt(secretKey, encryptedData)
	if err != nil {
		fmt.Println("Decryption error:", err)
		return
	}
	fmt.Println("Decrypted data:", decryptedData)

	if data != decryptedData {
		fmt.Print("data is not matched")
	}
}
