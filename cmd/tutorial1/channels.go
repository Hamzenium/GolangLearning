package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

// Define a struct to hold the result from scraping each page along with its index
type PageResult struct {
	Index int
	URL   string
	Title string
}

var wg sync.WaitGroup

// Scrape function that fetches the content of a URL
func scrapePage(url string, index int, ch chan<- PageResult) {
	defer wg.Done() // Decrement the WaitGroup counter when done

	// Send GET request to the URL
	resp, _ := http.Get(url) // Ignore error for now
	defer resp.Body.Close()

	// Read the response body
	body, _ := ioutil.ReadAll(resp.Body) // Ignore error for now

	// Simple logic to extract the title from the HTML (for demo purposes)
	title := extractTitle(body)

	// Send the result to the channel along with the index to preserve order
	ch <- PageResult{Index: index, URL: url, Title: title}
}

// A simple function to extract the title of the page (very basic example)
func extractTitle(body []byte) string {
	// This is a very basic way to extract the title, it doesn't handle all edge cases
	// In a real-world scenario, you'd likely use a proper HTML parsing library
	start := "<title>"
	end := "</title>"
	startIndex := stringIndex(body, start)
	endIndex := stringIndex(body, end)

	if startIndex == -1 || endIndex == -1 {
		return "No Title"
	}

	// Extract the title string
	return string(body[startIndex+len(start) : endIndex])
}

// Helper function to find the index of a substring within a byte slice
func stringIndex(body []byte, substr string) int {
	return stringIndexFrom(body, substr, 0)
}

func stringIndexFrom(body []byte, substr string, start int) int {
	for i := start; i < len(body)-len(substr); i++ {
		if string(body[i:i+len(substr)]) == substr {
			return i
		}
	}
	return -1
}

func main() {
	// List of URLs to scrape
	urls := []string{
		"https://golang.org",
		"https://github.com",
	}

	// Channel to receive scrape results
	ch := make(chan PageResult, len(urls))

	// Launch goroutines to scrape each page, pass the index to preserve the order
	for i, url := range urls {
		wg.Add(1)
		go scrapePage(url, i, ch)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Close the channel when done
	close(ch)

	// Create a slice to store the results in the correct order
	results := make([]PageResult, len(urls))

	// Collect results and place them in the correct index
	for result := range ch {
		results[result.Index] = result
	}

	// Process the results in the correct order
	for _, result := range results {
		fmt.Printf("Scraped URL: %s, Title: %s\n", result.URL, result.Title)
	}
}
