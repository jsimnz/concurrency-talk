package main

import (
	"fmt"
)

const (
	NUM_CHANNELS = 2
)

func connectN(chans ...chan string) {
	for i, ch := range chans {
		if i != len(chans)-1 {
			go connect(chans[i+1], ch)
		}
	}
}

func connect(out chan string, in chan string) {
	for str := range in {
		out <- str
	}
}

func main() {
	chans := make([]chan string, NUM_CHANNELS)
	for i := 0; i < NUM_CHANNELS; i++ {
		chans[i] = make(chan string)
	}
	connectN(chans...)

	chans[0] <- "Hello World"            // in at the beggining of the chain
	fmt.Println(<-chans[NUM_CHANNELS-1]) // out at the end
}
