package client

import (
	"strings"
	"fmt"
)

// Derive the audience from base URL
func createAudience(baseUrl string) (string) {
	urlTokens := strings.Split(baseUrl, "/")
	host := urlTokens[2]
	api_name := urlTokens[4]
	return fmt.Sprintf("https://%s/%s", host, api_name)
}
