// Permutations of a given string

// abc
package main

import "fmt"

func permutationString(str string) []string {
	res := []string{}
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		currStr := string(runes[i])
		for j := (i + 1) % len(runes); j != i; {
			currStr += string(runes[j])
			j++
			j = j % len(runes)
		}
		fmt.Println("Character at index", currStr)
		res = append(res, currStr)
	}
	return res
}
