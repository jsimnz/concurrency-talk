package main

import (
	"fmt"
	"time"
)

func lameFn(loops int) {
	for i := 0; i < loops; i++ {
		fmt.Println("Loop #:", i)
	}
}

func main() {
	go lameFn(10) // <-- We just went full async
	time.Sleep(time.Second * 1)
}
