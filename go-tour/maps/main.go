package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	count := make(map[string]int)
	for _, word := range strings.Fields(s) {
		count[word] += 1
	}
	return count
}

func main() {
	wc.Test(WordCount)
}
