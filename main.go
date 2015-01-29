package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	NUM_WORKERS = 10
	NUM_JOBS    = 100
)

func workerFn(input string) string {
	time.Sleep(time.Millisecond * time.Duration(500+rand.Int31n(500))) // Simulate doing some work for some amount of time between 0.5 - 1 second
	return fmt.Sprint(input, " - FINISHED")
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
	results := make(chan string, NUM_JOBS)

	// Start your workers!
	for i := 0; i < NUM_WORKERS; i++ {
		go worker(jobs, results)
	}
	go handleResults(results)

	// Submit jobs
	for i := 0; i < NUM_JOBS; i++ {
		jobs <- fmt.Sprintf("Job #: %d", i)
	}

	// Wait for everything to finish
	time.Sleep(time.Second * ((NUM_JOBS / NUM_WORKERS) + 1))
}
