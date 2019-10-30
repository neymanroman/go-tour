package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	result := make(map[string]int)
	fields := strings.Fields(s)

	for _, word := range fields {
		_, ok := result[word]
		if ok {
			result[word] = result[word] + 1
		} else {
			result[word] = 1
		}
	}

	return result
}

func main() {
	wc.Test(WordCount)
}
