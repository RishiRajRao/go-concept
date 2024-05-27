package main

func Fibonacci(num int) int {
	if num <= 1 {
		return num
	} else {
		return Fibonacci(num-1) + Fibonacci(num-2)
	}
}
