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
	lameFn(10)
}
