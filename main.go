package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {

	urls := []string{"https://jsonplaceholder.typicode.com/todos/1", "https://jsonplaceholder.typicode.com/todos/2", "https://jsonplaceholder.typicode.com/todos/3"}
	var (
		wg     sync.WaitGroup
		result []interface{}
	)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	ch := make(chan interface{})
	wg.Add(len(urls))

	for _, url := range urls {
		go getTodos(ctx, url, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for val := range ch {
		var todo Todo
		if err := json.Unmarshal(val.([]byte), &todo); err != nil {
			fmt.Printf("Error unmarshalling JSON: %v\n", err)
			continue
		}
		result = append(result, todo)
	}

	fmt.Println("final result ", result)

	defer cancel()

}

func getTodos(ctx context.Context, url string, ch chan<- interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		ch <- err
		return
		// return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		ch <- err
		return
		// return nil, fmt.Errorf("failed to get response: %w", err)
	}
	defer resp.Body.Close()

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to read response body: %w", err)
	// }

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- err
		return
	}

	ch <- body

}
