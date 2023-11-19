/*

Before running the code, make sure to install the necessary package:

Terminal:
	go get -u golang.org/x/net/html

*/

package features

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"golang.org/x/net/html"
)

// Crawl function takes a URL and recursively crawls the pages
func Crawl(url string, depth int) {
	if depth <= 0 {
		return
	}

	// Make an HTTP request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Parse the HTML content
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return
	}

	// Process the links on the current page
	processLinks(doc)

	// Recursively crawl the linked pages
	links := extractLinks(doc)
	for _, link := range links {
		Crawl(link, depth-1)
	}
}

// processLinks extracts and prints the links on the current page
func processLinks(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				fmt.Println("Link:", a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		processLinks(c)
	}
}

// extractLinks returns a slice of unique links from the HTML document
func extractLinks(n *html.Node) []string {
	var links []string
	visited := make(map[string]bool)

	var visitNode func(*html.Node)
	visitNode = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					link := a.Val
					if !visited[link] {
						links = append(links, link)
						visited[link] = true
					}
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			visitNode(c)
		}
	}

	visitNode(n)
	return links
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run crawler.go <url> <depth>")
		return
	}

	url := os.Args[1]
	depth := os.Args[2]
    IntDepth, _ := strconv.Atoi(depth)

	Crawl(url, IntDepth)
}
