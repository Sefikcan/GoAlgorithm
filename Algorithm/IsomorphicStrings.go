package main

import "fmt"

func isIsomorphic(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	result := make(map[rune]rune)

	// iterate
	for i := 0; i < len(s); i++ {
		sItem := rune(s[i])
		tItem := rune(t[i])

		if val, ok := result[sItem]; ok {
			if tItem != val {
				return false
			}
		} else {
			for _, isExists := range result {
				if isExists == tItem {
					return false
				}
			}

			result[sItem] = tItem
		}
	}

	return true
}

func main() {
	fmt.Println(isIsomorphic("egg", "add"))
}
