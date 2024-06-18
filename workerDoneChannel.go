package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	ID int
	// Add more fields as necessary
}

// Worker function to process jobs
func worker(id int, jobs <-chan Job, done <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-done:
			fmt.Printf("Worker %d received done signal, stopping...\n", id)
			return
		case job := <-jobs:
			fmt.Printf("Worker %d processing job %d\n", id, job.ID)
			// Simulate job processing time
			time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)
			fmt.Printf("Worker %d finished job %d\n", id, job.ID)
		}
	}
}

func maint() {
	jobs := make(chan Job)
	done := make(chan struct{})
	var wg sync.WaitGroup

	// Launch worker goroutines
	numWorkers := 3
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, done, &wg)
	}

	// Simulate sending jobs to workers
	go func() {
		jobID := 1
		for {
			select {
			case <-done:
				close(jobs)
				return
			default:
				jobs <- Job{ID: jobID}
				fmt.Printf("Sent job %d to the queue\n", jobID)
				jobID++
				time.Sleep(500 * time.Millisecond) // Simulate job arrival rate
			}
		}
	}()

	// Let the workers process jobs for a while
	time.Sleep(10 * time.Second)

	// Signal the workers to stop
	close(done)

	// Wait for all workers to finish
	wg.Wait()

	fmt.Println("All workers stopped")
}
