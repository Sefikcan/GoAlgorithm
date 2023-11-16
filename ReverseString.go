package main

import "fmt"

func reverseString(s []string) {
	start, end := 0, len(s)-1

	// walk over the string "s"
	for ; start < end; start, end = start+1, end-1 {
		// swap the characters at start and end
		s[start], s[end] = s[end], s[start]
	}
}

func main() {
	s1 := []string{"h", "e", "l", "l", "o"}
	reverseString(s1)
	fmt.Println(s1)
}
