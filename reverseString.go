package main

import (
	"slices"
	"strings"
)

func reverseString(s string) string {
	stringsArr := strings.Split(s, ".")
	slices.Reverse(stringsArr)
	result := strings.Join(stringsArr, ".")
	return result
}
