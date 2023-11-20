package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	// split the words in the string
	words := strings.Split(strings.TrimSpace(s), " ")

	if len(words) == 0 {
		return 0
	}

	return len(words[len(words)-1])
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
}
