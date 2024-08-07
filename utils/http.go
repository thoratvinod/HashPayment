package utils

import (
	"fmt"
	"net/http"
)

func GetBaseURL(r *http.Request) string {
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	return fmt.Sprintf(scheme + "://" + r.Host)
}
