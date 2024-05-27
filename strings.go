package main

import (
	"fmt"
)

func StringFunc() {
	board := []string{"rishi", "raj", "rao"}
	board = append(board, "nup")

	fmt.Println("value is", board)
}
