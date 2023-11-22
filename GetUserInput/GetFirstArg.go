package getuserinput

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/sinalalebakhsh/Gocron/features"

	"net/http"
	"strconv"

	"golang.org/x/net/html"
)


func GetFirstArg() bool {

	if len(os.Args) == 1 {
		color.HiBlue(features.LOGO)
		return true
	} else if len(os.Args) > 1 {
		AllArgs := os.Args

		for _, value := range features.TitleOfHelp {
			FirstArg := strings.ToLower(AllArgs[1])
			if value == FirstArg {
				features.HelpMessage()
				return false
			}
		}

		for _, value := range features.TitleOfEveryWords {
			FirstArg := strings.ToLower(AllArgs[1])
			if value == FirstArg {
				features.HelpMessage()
				color.HiBlue(fmt.Sprintln(features.OriginalFeatures))
			}
		}
		if len(os.Args) >= 2 {			
			url := os.Args[1]
			depthstr := os.Args[2]
			
			depth, err := strconv.Atoi(depthstr)
			if err != nil {
				fmt.Println("Error converting depth to integer:", err)
				return false
			}
		
			CrawlArgs(url, depth)
		}
	

	}

	return true
}

var SliceOfSearching []string

// CrawlArgs function takes a URL and recursively crawls the pages
func CrawlArgs(url string, depth int) {
	if depth <= 0 {
		return
	}

	// Make an HTTP request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making request:", err)
		// storeObject := fmt.Sprint("Error making request:", err)
		// SliceOfSearching = append(SliceOfSearching, storeObject)
		return
	}
	defer resp.Body.Close()

	// Parse the HTML content
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		// storeObject := fmt.Sprint("Error parsing HTML:", err)
		// SliceOfSearching = append(SliceOfSearching, storeObject)
		return
	}

	// Process the links on the current page
	ProcessLinksArgs(doc)

	// Recursively crawl the linked pages
	links := ExtractLinksArgs(doc)
	for _, link := range links {
		CrawlArgs(link, depth-1)
	}
}

// ProcessLinksArgs extracts and prints the links on the current page
func ProcessLinksArgs(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
			fmt.Println("Link:", a.Val)
			// storeObject := fmt.Sprint("Link:", a.Val)
			// SliceOfSearching = append(SliceOfSearching, storeObject)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ProcessLinksArgs(c)
	}
}

// ExtractLinksArgs returns a slice of unique links from the HTML document
func ExtractLinksArgs(n *html.Node) []string {
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



/*
func RunStore(SliceOfSearching []string){


	// Step 1: Read current count from counter.txt or initialize to 0
	count, err := readCounter()
	if err != nil {
		fmt.Println("Error reading counter:", err)
		return
	}

	// Step 2: Create A{count}.txt
	filename := fmt.Sprintf("Crawl/A%d.txt", count)
	err = createFile(filename, SliceOfSearching)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	// Step 3: Increment count and write it back to counter.txt
	err = incrementAndWriteCounter(count + 1)
	if err != nil {
		fmt.Println("Error updating counter:", err)
		return
	}

	fmt.Printf("Program completed successfully. Created %s\n", filename)
}

func readCounter() (int, error) {
	// Read current count from counter.txt or initialize to 0
	content, err := os.ReadFile("counter.txt")
	if err != nil {
		if os.IsNotExist(err) {
			return 0, nil // File does not exist, starting from 0
		}
		return 0, err
	}

	// Parse the last line as an integer
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	lastLine := lines[len(lines)-1]
	count, err := strconv.Atoi(lastLine)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func createFile(filename string, content []string) error {
	// Create a new file with the given filename
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close() // Close the file when the function completes

	// Write the contents of the slice with line numbers to the file
	for i, line := range content {
		_, err := file.WriteString(fmt.Sprintf("%d: %s\n", i+1, line))
		if err != nil {
			return err
		}
	}

	return nil
}


func incrementAndWriteCounter(count int) error {
	// Append the updated count to counter.txt
	file, err := os.OpenFile("counter.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close() // Close the file when the function completes

	// Write the incremented count to the file
	_, err = file.WriteString(fmt.Sprintf("%d\n", count))
	if err != nil {
		return err
	}

	return nil
}
*/