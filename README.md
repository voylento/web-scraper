# web-scraper
This is my code for the guided project "Build a Web Scraper in Go" on [Boot.dev](https://www.boot.dev)

This is a command line program that takes a web url and traveres all the links on the webpage that are
part of the same root url. It stores the number of times each link (relative or absolute) is linked via an
anchor tag with href. It prints a report at the end.

To run, after cloning:

go run . base-url max-concurrent-goroutines max-pages-to-traverse

example:
go run . "https://www.wagslane.dev" 5 250

produces the report below:

===================================<br>
REPORT for https://www.wagslane.dev
===================================<br>
Found 63 internal links to www.wagslane.dev
Found 62 internal links to www.wagslane.dev/about
Found 62 internal links to www.wagslane.dev/index.xml
Found 62 internal links to www.wagslane.dev/tags
Found 5 internal links to www.wagslane.dev/posts/leave-scrum-to-rugby
...




