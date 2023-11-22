/*

Before running the code, make sure to install the necessary package:

Terminal:
	go get -u golang.org/x/net/html

*/

package crawl


import (
	"fmt"
	"net/http"
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

	// Process the Links on the current page
	ProcessLinks(doc)

	// Recursively crawl the linked pages
	Links := ExtractLinks(doc)
	for _, link := range Links {
		Crawl(link, depth-1)
	}
}

// ProcessLinks extracts and prints the Links on the current page
func ProcessLinks(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				fmt.Println("Link:", a.Val)
			}
		}
	}

	for C := n.FirstChild; C != nil; C = C.NextSibling {
		ProcessLinks(C)
	}
}

// ExtractLinks returns a slice of unique Links from the HTML document
func ExtractLinks(n *html.Node) []string {
	var Links []string
	visited := make(map[string]bool)

	var visitNode func(*html.Node)
	visitNode = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					link := a.Val
					if !visited[link] {
						Links = append(Links, link)
						visited[link] = true
					}
				}
			}
		}

		for C := n.FirstChild; C != nil; C = C.NextSibling {
			visitNode(C)
		}
	}

	visitNode(n)
	return Links
}
