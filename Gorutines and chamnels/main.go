// You have a list of URLs (like https://golang.org, https://example.com).

// You create a channel for jobs (URLs) and a channel for results (page titles or errors).

// You create N workers (goroutines).

// Each worker takes one URL from the jobs channel.

// It tries to fetch the page.

// If successful → extracts the <title> tag.

// If too slow (>2 seconds) → sends an error.

// Sends result/error to the results channel.

// The main function:

// Puts all URLs into the jobs channel.

// Collects results from the results channel.

// Stops everything if we already got enough results (using context.Cancel).
package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"golang.org/x/net/html"
)

func main() {
	var wg sync.WaitGroup
	UrlList := []string{"https://golang.org", "https://example.com", "https://httpbin.org/delay/10"}
	Jobchan := make(chan string)

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			RecieveChannel(Jobchan, &wg)

		}()
	}
	for _, url := range UrlList {
		Jobchan <- url
	}

	close(Jobchan)
	wg.Wait()
}

func RecieveChannel(Jobchan chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	resultChan := make(chan string, 1)

	url := <-Jobchan

	go func() {
		fmt.Println(url)
		response, err := http.Get(url)
		defer response.Body.Close()
		if err != nil {
			fmt.Println("Error fetching:", url, "-", err)
			return
		}
		doc, err := html.Parse(response.Body)
		if err != nil {

			fmt.Print(err)
		}
		title := getTitle(doc)
		resultChan <- title
		close(resultChan)

	}()
	select {
	case data := <-resultChan:
		fmt.Println("The title is", data)
	case <-time.After(2 * time.Second):
		fmt.Println("Cannot complete this operation taking more than 2 seconds")

	}

}

func getTitle(n *html.Node) string {

	//html
	//  ├── head
	//  │     └── title
	//  │           └── "Example Domain"   (this is a text node)
	//  └── body
	//        └── h1
	//              └── "Hello"
	// Check if this node is a <title>
	if n.Type == html.ElementNode && n.Data == "title" {
		if n.FirstChild != nil {
			return n.FirstChild.Data // return the text inside <title>
		}
	}

	// Recursively search children
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result := getTitle(c)
		if result != "" {
			return result
		}
	}

	// If not found, return empty string
	return ""
}

//

// func main() {
// 	ch := make(chan string)

// 	go func() {
// 		var1 := <-ch
// 		fmt.Println(var1)
// 	}()
// 	ch <- "ball"

// 	var1 := <-ch
// 		fmt.Println(var1)
// }
