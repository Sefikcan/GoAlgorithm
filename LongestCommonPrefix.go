package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// set the prefix
	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return prefix
			}
		}
	}

	return prefix
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}
