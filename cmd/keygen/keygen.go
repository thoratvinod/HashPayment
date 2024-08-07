package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/thoratvinod/HashPayment/utils"
)

func hexToBytes(hexStr string) ([]byte, error) {
	// Decode hex string to bytes
	data, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func getEncrptedKey(hexSecret, key string) {

	secret, err := hexToBytes(hexSecret)
	if err != nil {
		fmt.Println(err.Error())
	}

	encrypted, err := utils.Encrypt(secret, key)
	if err != nil {
		fmt.Println(err.Error())
	}

	println("Encrypted key: ", encrypted)
}

// first argument is secret and second argument key which we want to encrypt
func main() {
	if len(os.Args) < 3 {
		println("Provide secret and key in command line arguments")
		os.Exit(1)
		return
	}
	getEncrptedKey(os.Args[1], os.Args[2])
}
