package main

import (
	"fmt"
	"strings"
)

func reverse(words []string) {
	start, end := 0, len(words)-1

	for start < end {
		words[start], words[end] = words[end], words[start]
		start++
		end--
	}
}

func reverseWords(s string) string {
	// remove the spaces in the input string
	// split string into words
	words := strings.Fields(strings.TrimSpace(s))

	// reverse the words
	reverse(words)

	// join the reversed words
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(reverseWords("the sky is blue"))
}
