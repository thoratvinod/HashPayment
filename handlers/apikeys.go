package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/thoratvinod/HashPayment/services"
	"github.com/thoratvinod/HashPayment/specs"
)

func SetAPIKeysHandler(w http.ResponseWriter, r *http.Request) {
	request := specs.SetAPIKeysRequest{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	for gateway, apiKey := range request.APIKeys {
		if gateway == "" {
			http.Error(w, "Gateway name cannot be empty", http.StatusBadRequest)
			return
		}
		err := services.GetAPIKeyManager().Set(gateway, apiKey)
		if err != nil {
			http.Error(w, fmt.Sprintf("decoding of API key failed for %v: %+v", gateway, err), http.StatusBadRequest)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "API keys updated successfully.",
	})
}
