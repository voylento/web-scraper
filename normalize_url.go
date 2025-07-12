package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("error: could not parse URL: %w", err)
	}

	fullPath := parsedURL.Host + parsedURL.Path
	fullPath = strings.TrimSuffix(fullPath, "/")
	fullPath = strings.ToLower(fullPath)

	return fullPath, nil
}

