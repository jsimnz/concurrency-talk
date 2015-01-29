package main

import (
	"fmt"
	"time"
)

const (
	NUM_WORKERS = 1
	NUM_JOBS    = 5
)

func workerFn(input string) string {
	time.Sleep(time.Second) // Simulate doing some work
	return fmt.Sprint(input, " - FINISHED")
}

func worker(jobs chan string) {
	for job := range jobs {
		output := workerFn(job)
		fmt.Println(output)
	}
}

func main() {
	jobs := make(chan string, 1000)

	// Start your workers!
	for i := 0; i < NUM_WORKERS; i++ {
		go worker(jobs)
	}

	// Submit jobs
	for i := 0; i < NUM_JOBS; i++ {
		jobs <- fmt.Sprintf("Job #: %d", i)
	}

	// Wait for everything to finish
	time.Sleep(time.Second * ((NUM_JOBS / NUM_WORKERS) + 1))
}
