package main

import "fmt"

func main() {

	var arr = []int{7, 2, 1}
	var sumArr = 2

	var start, end = maxSubArrOpt(arr, sumArr)

	fmt.Println("print value ", start, end)

}
