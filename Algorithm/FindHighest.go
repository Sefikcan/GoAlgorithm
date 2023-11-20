package main

import "fmt"

func main() {
	fmt.Println(findHighest([]int{1, 3, 12, 5}))
}

func findHighest(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var maxNums int = nums[0]

	for i := 1; i < len(nums); i++ {
		if maxNums < nums[i] {
			maxNums = nums[i]
		}
	}

	return maxNums
}
