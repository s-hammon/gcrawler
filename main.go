package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args[1:]) < 3 {
		fmt.Println("too few arguments provided")
		os.Exit(1)
	}
	if len(os.Args[1:]) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURL := os.Args[1]
	maxConcurrency, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("error parsing max concurrency: %v", err)
		return
	}

	maxPages, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Printf("error parsing max pages: %v", err)
		return
	}

	cfg, err := configure(baseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("error configuring: %v", err)
		return
	}
	fmt.Printf("starting to crawl %s\n", baseURL)

	cfg.wg.Add(1)
	cfg.crawlPage(baseURL)
	cfg.wg.Wait()

	printReport(cfg.pages, baseURL)
}
