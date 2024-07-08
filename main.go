package main

import "fmt"

func main() {

	// var val = "ABC"
	var val = []int{-2, -3, 4, -1, -2, 1, 5, -3}

	var start = kadanesAlgo(val)

	fmt.Printf("print value %v\n", start)

}
