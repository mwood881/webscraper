package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/gocolly/colly"
)

// Struct for storing scraped data
type ScrapedData struct {
	URL  string `json:"url"`
	Text string `json:"text"`
}

// Wikipedia URLs to scrape
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

func scrape(url string, wg *sync.WaitGroup, dataChan chan<- ScrapedData) {
	defer wg.Done()

	// Create a new collector
	c := colly.NewCollector()

	// Variable to store scraped text
	var textContent string

	// OnHTML extracts text from the main content section
	c.OnHTML("#mw-content-text", func(e *colly.HTMLElement) {
		textContent = e.Text
	})

	// Handle errors
	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Error scraping %s: %s\n", r.Request.URL, err)
	})

	// Start scraping
	c.Visit(url)

	// Send data to the channel
	dataChan <- ScrapedData{URL: url, Text: textContent}
}

func main() {
	// Create output file
	start := time.Now()
	file, err := os.Create("scraped_data.jl")
	if err != nil {
		log.Fatal("Could not create file:", err)
	}
	defer file.Close()

	// Channel for collecting scraped data
	dataChan := make(chan ScrapedData, len(urls))
	var wg sync.WaitGroup

	// Launch concurrent scrapers
	for _, url := range urls {
		wg.Add(1)
		go scrape(url, &wg, dataChan)
	}

	// Close the channel after all goroutines finish
	go func() {
		wg.Wait()
		close(dataChan)
	}()

	// Write JSON lines to file
	for data := range dataChan {
		jsonData, _ := json.Marshal(data)
		file.WriteString(string(jsonData) + "\n")
	}

	fmt.Println("Scraping completed. Data saved to scraped_data.jl")

	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)
}
