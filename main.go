package main

import (
	"fmt"
)

func connect(out chan string, in chan string) {
	for str := range in {
		out <- str
	}
}

func main() {
	in := make(chan string)
	out := make(chan string)
	go connect(out, in)

	in <- "Hello World"
	fmt.Println(<-out)
}
