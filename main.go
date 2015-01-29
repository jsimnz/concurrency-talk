package main

import (
	"fmt"
)

func main() {
	in := make(chan string, 1)

	in <- "Hello World"
	fmt.Println(<-in)
}
