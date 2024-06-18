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
