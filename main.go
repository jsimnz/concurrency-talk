package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	NUM_WORKERS = 1
	NUM_JOBS    = 5
)

func workerFn(input string) string {
	time.Sleep(time.Millisecond * time.Duration(500+rand.Int31n(500))) // Simulate doing some work for some amount of time between 0.5 - 1 second
	return fmt.Sprint(input, " - FINISHED")
}

func worker(jobs chan string) {
	for job := range jobs {
		output := workerFn(job)
		fmt.Println(output)
	}
}

func main() {
	jobs := make(chan string, NUM_JOBS)

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
