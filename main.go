package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	NUM_WORKERS     = 2
	TIMEOUT_SECONDS = 5
)

var (
	urls = []string{
		"http://google.com",
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
	timeout := time.NewTimer(time.Second * TIMEOUT_SECONDS)
	result := make(chan string)
	go func(result chan string) {
		_, err := http.Get(url)
		if err != nil {
			result <- fmt.Sprintf("Failed to get %v", url)
		}
		result <- fmt.Sprintf("Sucessfully got %v", url)
	}(result)

	select {
	case val := <-result:
		return val
	case <-timeout.C:
		return fmt.Sprintf("Timed out trying to get %v", url)
	}

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
