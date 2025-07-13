package main

import (
	"fmt"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {
	maxConcurrency := 5
	maxPages :=	100

	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(os.Args) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	rawBaseURL := strings.TrimSpace(os.Args[1])
	if len(os.Args) > 2 {
		concurrency, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Error converting %s to int\n", os.Args[2])
			os.Exit(1)
		}
		maxConcurrency = concurrency
	}
	if len(os.Args) > 3 {
		pages, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Printf("Error converting %s to int\n", os.Args[3])
			os.Exit(1)
		}
		maxPages = pages
	}
	

	fmt.Println("===========================================")
	fmt.Printf("starting crawl of: %s\n", rawBaseURL) 
	fmt.Println("===========================================")

	parsedBaseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("Error parsing base URL %s: %v\n", rawBaseURL, err)
		os.Exit(1)
	}

	fmt.Printf("maxConcurrency = %d\n", maxConcurrency)
	fmt.Printf("maxPages = %d\n", maxPages)
	cfg := &config{
		pages:							make(map[string]int),
		maxPages:						maxPages,
		baseURL:						parsedBaseURL,
		mu:									&sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:									&sync.WaitGroup{},
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)

	cfg.wg.Wait()

	fmt.Println("=================================================")
	fmt.Println("CRAWL COMPLETE")
	fmt.Println("=================================================")

	printReport(cfg.pages, rawBaseURL)
}

type PageLink struct {
	Page		string
	Links		int
}

func printReport(pages map[string]int, baseURL string) {
	fmt.Println("========================================")
	fmt.Printf("REPORT for %s\n", baseURL)
	fmt.Println("========================================")

	var pageLinks []PageLink
	for page, links := range pages {
		pageLinks = append(pageLinks,  PageLink{
			Page:		page,
			Links:	links,
		})
	}

	sort.Slice(pageLinks, func(i, j int) bool {
		// If number of links are not equal, sort in descending order by  number of links
		if pageLinks[i].Links != pageLinks[j].Links {
			return pageLinks[i].Links > pageLinks[j].Links
		}

		// if the number of links are equal, sort alphabetically by Page url
		return pageLinks[i].Page < pageLinks[j].Page
	})

	for _, pageLink := range pageLinks {
		fmt.Printf("Found %d internal links to %s\n", pageLink.Links, pageLink.Page)
	}
}
