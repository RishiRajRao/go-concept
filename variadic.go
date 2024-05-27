package main

import "fmt"

func Variadic(first string, second string, all ...string) string {
	sum := ""

	for _, val := range all {
		sum += val
	}

	println(first, second, sum)
	final := fmt.Sprintf(first + second + sum)
	return final
}
