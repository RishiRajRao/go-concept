package main

import (
	"fmt"
	"sync"
	"time"
)

func maindy() {
	n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	ch := make(chan struct{})
	var repoMutex sync.Mutex // To protect the repo slice
	repo := []int{}

	var wg sync.WaitGroup
	wg.Add(2)

	// Worker 1
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ch:
				return
			default:
				repoMutex.Lock()
				repo = []int{} // Clear repo to observe the effect
				for _, val := range n {
					repo = append(repo, val*0)
				}
				fmt.Println("Worker 1 updated repo:", repo)
				repoMutex.Unlock()
				time.Sleep(100 * time.Millisecond) // Prevent tight loop
			}
		}
	}()

	// Worker 2
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ch:
				return
			default:
				repoMutex.Lock()
				repo = []int{} // Clear repo to observe the effect
				for _, val := range n {
					repo = append(repo, val*2)
				}
				fmt.Println("Worker 2 updated repo:", repo)
				repoMutex.Unlock()
				time.Sleep(100 * time.Millisecond) // Prevent tight loop
			}
		}
	}()

	// Allow workers to run for a while
	time.Sleep(2 * time.Second)
	close(ch)

	wg.Wait()
	fmt.Println("Final repo:", repo)
}
