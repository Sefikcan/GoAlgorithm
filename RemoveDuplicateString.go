package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(removeDuplicate("abaccc"))
}

func removeDuplicate(str string) string {
	var removeDuplicate string = ""

	for _, char := range str {
		if !strings.ContainsRune(removeDuplicate, char) {
			removeDuplicate += string(char)
		}
	}

	return removeDuplicate
}
