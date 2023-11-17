package main

import "fmt"

func reverseStr(s string, k int) string {
	// Convert string to slice of runes for easy modification
	str := []rune(s)
	for i := 0; i < len(str); i += 2 * k {
		left := i
		right := i + k - 1
		if right >= len(str) {
			right = len(str) - 1
		}

		// Reverse the first k characters
		for left < right {
			str[left], str[right] = str[right], str[left]
			left++
			right--
		}
	}

	return string(str)
}

func main() {
	s1 := "abcdefg"
	k1 := 2
	result1 := reverseStr(s1, k1)
	fmt.Println(result1)
}
