package main

import (
	"strings"
)

func checkPalindrome(str string) bool {
	str = strings.ToLower(str)
	chars := []rune(str)
	// isPalin := true

	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		if chars[i] != chars[j] {
			return false
		}
	}

	return true

}
