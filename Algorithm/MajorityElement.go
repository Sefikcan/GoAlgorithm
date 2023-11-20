package main

import "fmt"

func majorityElement(nums []int) int {
	// Create a map to find majority elements
	dict := make(map[int]int)

	// iterate the nums and count the number of each element
	for _, num := range nums {
		dict[num]++

		// if the number of an elements exceeds n/2, that is the majority element
		if dict[num] > len(nums)/2 {
			return num
		}
	}

	return -1
}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
}
