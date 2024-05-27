package main

import (
	"fmt"
)

// Producer: Send-only channel
func producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}

// Consumer: Receive-only channel
func consumer(ch <-chan int) {
	for value := range ch {
		fmt.Println("Received:", value)
	}
}

func main() {
	ch := make(chan int)

	go producer(ch) // Start producer goroutine
	consumer(ch)    // Start consumer in main goroutine
}
