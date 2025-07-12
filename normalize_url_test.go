package main

import (
	"strings"
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name					string
		inputURL			string
		expected			string
		errorContains	string
	}{
		{
			name:			"remove scheme",
			inputURL:	"https://blog.boot.dev/path",
			expected:	"blog.boot.dev/path",
		},
		{
			name:			"remove trailing slash",
			inputURL:	"http://blog.boot.dev/path/",
			expected:	"blog.boot.dev/path",
		},
		{
			name:			"url set to lowercase",
			inputURL:	"BLOG.boot.Dev/path",
			expected:	"blog.boot.dev/path",
		},
		{
			name:			"error on invalid url",
			inputURL:	":\\blog.",
			expected: "",
			errorContains: "error: could not parse URL",	
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil && !strings.Contains(err.Error(), tc.errorContains) {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err != nil && tc.errorContains == "" {
				t.Errorf("Test %v = '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err == nil && tc.errorContains != "" {
				t.Errorf("Test %v = '%s' FAIL: expected error containing '%v', got none.", i, tc.name, tc.errorContains)
				return
			}

			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

