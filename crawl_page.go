package main

import (
	"fmt"
	"net/url"
)

func (c *config) crawlPage(rawCurrentURL string) {
	c.chConn <- struct{}{}
	defer func() {
		<-c.chConn
		c.wg.Done()
	}()

	if c.pagesLen() >= c.maxPages {
		return
	}

	currentUrl, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("couldn't parse URL: %v", err)
		return
	}

	if currentUrl.Hostname() != c.baseURL.Hostname() {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("error normalizing URL: %v", err)
	}

	isFirst := c.addPageVisit(normalizedURL)
	if !isFirst {
		return
	}

	fmt.Printf("crawling %s\n", rawCurrentURL)
	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error fetching %s: %v", rawCurrentURL, err)
		return
	}

	links, err := getURLsFromHTML(html, c.baseURL)
	if err != nil {
		fmt.Printf("error getting URLs from HTML: %v", err)
		return
	}

	for _, link := range links {
		c.wg.Add(1)
		go c.crawlPage(link)
	}
}
