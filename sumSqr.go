package main

import "time"

func sumSqr(value int) int {
	ch, done, sum := make(chan int, 1), make(chan bool), 0

	go func() {
		for {
			select {
			case <-done:
				return
			default:

				for i := 0; i < value; i++ {
					ch <- (i + 1) * (i + 1)
					println("value=", i)
				}
				time.Sleep(time.Second * 1)
				close(done)
				close(ch)

			}

		}
	}()

	for val := range ch {
		sum += val
		println("in channel", sum)
	}
	return sum
}
