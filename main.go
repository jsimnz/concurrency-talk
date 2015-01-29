package main

import (
	"fmt"
	"net/http"
	//"time"
)

const (
	NUM_WORKERS = 2
)

var (
	urls = []string{"http://google.com",
		"http://twitter.com",
		"http://facebook.com",
		"http://conferencecloud.co",
		"http://google.com",
		"http://duckduckgo.com",
		"http://golang.com",
	}

	NUM_JOBS = len(urls)
)

func workerFn(url string) string {
	_, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("Failed to get %v", url)
	}
	return fmt.Sprintf("Sucessfully got %v", url)
}

func worker(jobs chan string, results chan string) {
	for job := range jobs {
		output := workerFn(job)
		results <- output
	}
}

func handleResults(results chan string) {
	for result := range results {
		fmt.Println(result)
	}
}

func main() {
	jobs := make(chan string, NUM_JOBS)
	results := make(chan string)

	// Start your workers!
	for i := 0; i < NUM_WORKERS; i++ {
		go worker(jobs, results)
	}
	go handleResults(results)

	// Submit jobs
	for i := 0; i < NUM_JOBS; i++ {
		jobs <- urls[i]
	}

	// Block indefinately
	<-make(chan bool)
}
