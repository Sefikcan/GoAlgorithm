package main

import "fmt"

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// build hashmap for "s"
	anagram := make(map[rune]int) // char::freq key::value

	for _, char := range s {
		anagram[char]++
	}

	// verify / compare freq across "s" and "t"
	for _, char := range t {
		anagram[char]--

		if anagram[char] < 0 {
			return false
		}
	}

	// "s" and "t" are anagrams
	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}
