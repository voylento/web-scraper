package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("Error: GET operation on %s failed: %v\n", rawURL, err)
	}
	defer res.Body.Close()

	if res.StatusCode >= 399 {
		return "", fmt.Errorf("Error: GET operation returned status code %d", res.StatusCode)
	}

	content_type := res.Header[http.CanonicalHeaderKey("content-type")]
	if !strings.Contains(content_type[0], "text/html") {
		return "", fmt.Errorf("Error: invalid content type at %s, expect 'text/html'",  rawURL)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("Error: error reading body: %v", err)
	}

	return string(body), nil 
}
