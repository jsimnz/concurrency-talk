package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	NUM_JOBS = 5
)

func workerFn(input string) string {
	time.Sleep(time.Millisecond * time.Duration(500+rand.Int31n(500))) // Simulate doing some work for some amount of time between 0.5 - 1 second
	return fmt.Sprint(input, " - FINISHED")
}

func main() {
	// Submit jobs
	for i := 0; i < NUM_JOBS; i++ {
		output := workerFn(fmt.Sprintf("Job #: %d", i))
		fmt.Println(output)
	}
}
