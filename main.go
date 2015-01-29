package main

import (
	"fmt"
)

func lameFn(loops int) {
	for i := 0; i < loops; i++ {
		fmt.Println("Loop #:", i)
	}
}

func main() {
	go lameFn(10) // <-- We just went full async
}
