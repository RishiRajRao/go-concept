package main

import (
	"fmt"
	"sync"
	"time"
)

func ChannelsFunc() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(2)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 10)
		fmt.Println("go routine")
		ch <- 1
		ch <- 2
	}()

	go func() {
		defer wg.Done()
		fmt.Println("channel value:", <-ch, <-ch)
		fmt.Println("waiting for channels inside go routine")

	}()
	// fmt.Println("in main stack go func", <-ch)
	fmt.Println("waiting for channels in func stack")

	wg.Wait()

}
