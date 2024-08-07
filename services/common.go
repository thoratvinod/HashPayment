package services

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/thoratvinod/HashPayment/utils"
)

func getDecryptedAPIKey(cryptoKey string) (string, error) {
	hexSecret := os.Getenv("SECRET_KEY")

	secretKey, err := hex.DecodeString(hexSecret)
	if err != nil {
		return "", fmt.Errorf("failed to decode secret from hex to binary: %+v", err)
	}

	decryptedData, err := utils.Decrypt(secretKey, cryptoKey)
	if err != nil {
		return "", err
	}
	return decryptedData, nil
}
