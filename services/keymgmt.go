package services

import (
	"fmt"
	"sync"
)

// APIKeyManager manages API keys for different gateways.
type APIKeyManager struct {
	apiKeys map[string]string
	mu      sync.RWMutex // Mutex to handle concurrent access
}

// Singleton instance
var (
	apiKeyManager *APIKeyManager
	once          sync.Once
)

// GetAPIKeyManager returns the singleton instance of APIKeyManager.
func GetAPIKeyManager() *APIKeyManager {
	once.Do(func() {
		apiKeyManager = &APIKeyManager{
			apiKeys: make(map[string]string),
		}
	})
	return apiKeyManager
}

// Get retrieves the API key for the specified gateway.
func (keyMgmt *APIKeyManager) Get(gateway string) (string, error) {
	keyMgmt.mu.RLock()
	defer keyMgmt.mu.RUnlock()
	apiKey, exists := keyMgmt.apiKeys[gateway]
	if !exists {
		return "", fmt.Errorf("API key not found for gateway: %v", gateway)
	}
	return apiKey, nil
}

// Set stores the API key for the specified gateway.
func (keyMgmt *APIKeyManager) Set(gateway, encryptedAPIKey string) error {
	if gateway == "" {
		return fmt.Errorf("gateway cannot be empty")
	}
	keyMgmt.mu.Lock()
	defer keyMgmt.mu.Unlock()

	plainAPIKey, err := getDecryptedAPIKey(encryptedAPIKey)
	if err != nil {
		return fmt.Errorf("failed to decrypt %v key: %+v", gateway, err)
	}

	keyMgmt.apiKeys[gateway] = plainAPIKey
	return nil
}
