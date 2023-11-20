package main

import "fmt"

func numJewelsInStones(jewels, stones string) int {
	jewelsHashMap := make(map[rune]bool)
	// freq of occurrance of jewels in stones
	count := 0

	// update the jewelHashMap
	// lets image that
	// jewel = "aA"
	// jewelsHashMap[a] = true
	// jewelsHashMap[A] = true
	for _, jewel := range jewels {
		jewelsHashMap[jewel] = true
	}

	// iterate over the stones
	// lets image that
	// stones = "aAAbbbb"
	// jewelHashMap[a] -> count = 1
	// jewelHashMap[A] -> count = 2
	// jewelHashMap[A] -> count = 3
	// jewelHashMap[b] -> count = 3
	// jewelHashMap[b] -> count = 3
	// jewelHashMap[b] -> count = 3
	// jewelHashMap[b] -> count = 3
	for _, stone := range stones {
		if jewelsHashMap[stone] == true {
			count++
		}
	}

	// count of jewels in stones
	return count
}

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb"))
}
