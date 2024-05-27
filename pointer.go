package main

import "fmt"

func PointerFunc() {

	var ptr *int
	a := 20
	ptr = &a

	fmt.Printf("value is %d and type is %d \n", a, *ptr)
}
