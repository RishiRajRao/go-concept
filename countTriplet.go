// count the triplets

// arr = [2,4,5,6,10]
package main

import (
	"fmt"
	"sort"
)

func countTriplet(arr []int) int {
	sort.Ints(arr)
	var count, sum = 0, 0
	fmt.Printf("value of arr %v\n", arr)
	for i := 0; i < len(arr); i++ {
		sum = arr[i]
		for j := i + 1; j < len(arr); j++ {
			sum += arr[j]
			for l := j + 1; l < len(arr); l++ {
				if sum == arr[l] {
					println("true")
					count++
				}
			}
			sum = arr[i]
		}
	}
	return count
}
