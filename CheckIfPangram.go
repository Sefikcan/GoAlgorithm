package main

import "fmt"

// Number of letters in english language
const expectedLetterCount = 26

func checkIfPangram(sentence string) bool {
	// create a hash table
	sentenceHash := make(map[rune]int)

	// iterate over the string
	for _, char := range sentence {
		sentenceHash[char]++
	}

	// check len of sentence equals number of letters in english language
	if len(sentenceHash) == expectedLetterCount {
		return true
	}

	return false
}

func main() {
	fmt.Println(checkIfPangram("thequickbrownfoxjumpsoverthelazydog"))
}
