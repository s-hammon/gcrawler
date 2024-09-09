package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages    map[string]int
	baseURL  *url.URL
	mu       *sync.Mutex
	chConn   chan struct{}
	wg       *sync.WaitGroup
	maxPages int
}

func (c *config) addPageVisit(normalizedURL string) (isFirst bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.pages[normalizedURL]; ok {
		c.pages[normalizedURL]++
		return false
	}

	c.pages[normalizedURL] = 1
	return true
}

func (c *config) pagesLen() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return len(c.pages)
}

func configure(rawBaseURL string, maxConcurrency, maxPages int) (*config, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse base URL: %v", err)
	}

	return &config{
		pages:    make(map[string]int),
		baseURL:  baseURL,
		mu:       &sync.Mutex{},
		chConn:   make(chan struct{}, maxConcurrency),
		wg:       &sync.WaitGroup{},
		maxPages: maxPages,
	}, nil
}
