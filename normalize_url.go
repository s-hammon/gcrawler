package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	myURL, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("could not parse URL: %v", err)
	}

	path := myURL.Host + myURL.Path
	path = strings.ToLower(path)
	path = strings.TrimSuffix(path, "/")

	return path, nil
}
