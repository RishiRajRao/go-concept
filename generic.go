package main

import "fmt"

type Animal[T any] struct {
	species T
}

func (animal *Animal[T]) SetAnimal(val T) {
	animal.species = val
}

func PrintAll[T any](list []T) {
	for _, val := range list {
		fmt.Println(val)
	}
}
