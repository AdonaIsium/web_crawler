package main

import (
	"fmt"
	"net/url"
)

func normalizeURL(s string) (string, error) {

	parsedURL, err := url.Parse(s)
	if err != nil {
		return "", err
	}
	if parsedURL.Scheme != "https" && parsedURL.Scheme != "" {
		return "", fmt.Errorf("expected 'https' host, got %s", parsedURL.Scheme)
	}

	normalizedURL := fmt.Sprintf("%s%s", parsedURL.Host, parsedURL.Path)

	return normalizedURL, nil
}
