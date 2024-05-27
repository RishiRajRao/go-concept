package main

import "fmt"

func ArrayAndSlice() {
	raw := make([]int, 5)
	arr := [5]int{1, 2, 3, 4, 5}
	slice1 := arr[2:4]
	slice2 := arr[3:4]
	slice2[0] = 9
	raw = append(raw, 1, 2)
	fmt.Println("Array is", arr, slice1, slice2, raw, len(raw), cap(raw))
}
