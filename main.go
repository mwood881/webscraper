// import package
package main

// import packages needed for assingment
import (
	"context"
	"encoding/json"
	"log"
	"os"
	"sync"
	"time"

	"github.com/gocolly/colly"
)

// Use type to structure stored data
// It holds the webpage URL and its extracted text
type ScrapedData struct {
	URL string `json:"url"`

	Text string `json:"text"`
}

// Wikipedia URLs to scrape about robotics and AI
// List was given to us in the Canvas instructions
// Make them into a string to store the urls
var urls = []string{
	"https://en.wikipedia.org/wiki/Robotics",
	"https://en.wikipedia.org/wiki/Robot",
	"https://en.wikipedia.org/wiki/Reinforcement_learning",
	"https://en.wikipedia.org/wiki/Robot_Operating_System",
	"https://en.wikipedia.org/wiki/Intelligent_agent",
	"https://en.wikipedia.org/wiki/Software_agent",
	"https://en.wikipedia.org/wiki/Robotic_process_automation",
	"https://en.wikipedia.org/wiki/Chatbot",
	"https://en.wikipedia.org/wiki/Applications_of_artificial_intelligence",
	"https://en.wikipedia.org/wiki/Android_(robot)",
}

// Function to scrape a given URL
// Sends the data through a channel to be written to a file
func scrape(ctx context.Context, url string, wg *sync.WaitGroup, dataChan chan<- ScrapedData) {
	defer wg.Done()

	// Create a new collector
	c := colly.NewCollector()

	// Variable to store scraped text
	var textContent string

	// extracts text from the content area of the Wikipedia page
	// used chatgpt to help learn how to scrape data
	c.OnHTML("#mw-content-text", func(e *colly.HTMLElement) {
		textContent = e.Text
	})

	// Handle errors
	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Error scraping %s: %v", url, err)
	})

	// Start scraping
	err := c.Visit(url)
	if err != nil {
		log.Printf("Error visiting %s: %v", url, err)
		return
	}

	select {
	case dataChan <- ScrapedData{URL: url, Text: textContent}:
	case <-ctx.Done():
		log.Printf("Scraping %s cancelled", url)
	}
}

func main() {
	// Create output file
	// Start timing the execution
	start := time.Now()
	outputFileName := "scraped_data.jl"

	file, err := os.Create(outputFileName)
	if err != nil {
		log.Fatalf("Could not create file: %v", err)
	}
	defer file.Close()

	// Channel for collecting scraped data
	dataChan := make(chan ScrapedData, len(urls))
	var wg sync.WaitGroup

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Launch concurrent scrapers
	// Loop through all URLs and start scraping each one
	for _, url := range urls {
		wg.Add(1)
		go scrape(ctx, url, &wg, dataChan)
	}

	// Close the channel after all goroutines finish
	go func() {
		wg.Wait()
		close(dataChan)
	}()

	// Write JSON lines to file
	for data := range dataChan {
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Printf("Error marshalling JSON for %s: %v", data.URL, err)
			continue
		}
		file.WriteString(string(jsonData) + "\n")
	}

	log.Printf("Data written to %s", outputFileName)

	// Calculate and print execution time
	elapsed := time.Since(start)

	log.Printf("Execution time: %s", elapsed)
}
