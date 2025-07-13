package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages								map[string]int
	maxPages						int
	baseURL							*url.URL
	mu									*sync.Mutex
	concurrencyControl	chan struct{}
	wg									*sync.WaitGroup
}


func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	fmt.Println("----------------------------------------")
	fmt.Printf("CRAWLING %s\n", rawCurrentURL)
	fmt.Println("----------------------------------------")

	cfg.mu.Lock()
	if len(cfg.pages) >= cfg.maxPages {
		cfg.mu.Unlock()
		return
	}
	cfg.mu.Unlock()

	parsedCurrentUrl, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error parsing current url %s: %v\n", rawCurrentURL, err)
		return
	}

	if cfg.baseURL.Hostname() != parsedCurrentUrl.Hostname() {
		return
	}

	currentUrl, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error getting normalized url for %s\n", rawCurrentURL)
		return
	}

	// thread-safe check if url already visited
	cfg.mu.Lock()
	if count, visited := cfg.pages[currentUrl]; visited {
		cfg.pages[currentUrl] = count + 1
		cfg.mu.Unlock()
		return
	}
	cfg.pages[currentUrl] = 1
	cfg.mu.Unlock()

	fullCurrentUrl := parsedCurrentUrl.Scheme + "://" + currentUrl
	htmlText, err := getHTML(fullCurrentUrl)
	if err != nil {
		fmt.Printf("Error getting html for page %s: %v\n", currentUrl, err)
		return
	}

	urls, err := getURLsFromHTML(htmlText, cfg.baseURL.String())
	if err != nil {
		fmt.Printf("Error getting urls from html for page %s: %v\n", currentUrl, err)
		return
	}

	for _, discoveredURL := range urls {
		cfg.wg.Add(1)
		go cfg.crawlPage(discoveredURL)
	}
}
