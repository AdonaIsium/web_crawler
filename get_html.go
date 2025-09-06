package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func getHTML(rawURL string) (string, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest(http.MethodGet, rawURL, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Accept", "text/html")
	req.Header.Set("User-Agent", "GosuCrawler/1.0")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return "", fmt.Errorf("unexpected status code received: %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}
