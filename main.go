package main

import (
	"fmt"
)

func main() {
	rootURL := "https://www.boot.dev"
	rawHTML := `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>`

	readHTMLandPrintURLs(rawHTML, rootURL)

rawHTML = `
<html>
	<body>
		<h1>This is a test</h1>
		<a href="/testity/test/test">Test</a>
	</html>
</body>`

	readHTMLandPrintURLs(rawHTML, rootURL)
}

func readHTMLandPrintURLs(html, rootURL string) {
	anchorTags, err := getURLsFromHTML(html, rootURL) 
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		fmt.Printf("Found %d urls in html\n", len(anchorTags))

		for i, tag := range anchorTags {
			fmt.Printf("anchorTag[%d] = %s\n", i, tag)
		}
}
