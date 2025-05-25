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
		return "", fmt.Errorf("error reading the url: %s", err)
	}

	defer res.Body.Close()

	if res.StatusCode > 399 {
		return "", fmt.Errorf("got HTTP error: %s", err)
	}

	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("got non-HTML response: %s", err)
	}

	htmlBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("couldn't read respond body: %s", err)
	}

	htmlBody := string(htmlBytes)

	return htmlBody, nil 
}

