/*

Before running the code, make sure to install the necessary package:

Terminal:
	go get -u golang.org/x/net/html

*/

package crawl
import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"time"
)

// Crawl function takes a URL and recursively crawls the pages
func Crawl(url string, depth int, searchDir string) {
	if depth <= 0 {
		return
	}

	// Create a directory for the search
	today := time.Now().Format("2006-01-02")
	searchDir = fmt.Sprintf("%s-%s", today, searchDir)

	err := os.Mkdir(searchDir, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	// Create a file to store search results
	resultFile, err := os.Create(fmt.Sprintf("%s/searchResult.txt", searchDir))
	if err != nil {
		fmt.Println("Error creating result file:", err)
		return
	}
	defer resultFile.Close()

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

	// Process the links on the current page and write them to the result file
	ProcessLinks(doc, resultFile)

	// Recursively crawl the linked pages
	links := ExtractLinks(doc)
	for _, link := range links {
		Crawl(link, depth-1, searchDir)
	}
}

// ProcessLinks extracts and prints the links on the current page
func ProcessLinks(n *html.Node, resultFile *os.File) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				link := fmt.Sprintf("Link: %s\n", a.Val)
				resultFile.WriteString(link)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ProcessLinks(c, resultFile)
	}
}

// ExtractLinks returns a slice of unique links from the HTML document
func ExtractLinks(n *html.Node) []string {
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