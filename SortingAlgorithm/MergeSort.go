package main

import "fmt"

/*
* Time complexity O(n* log(n))
* Space complexity is O(n)
 */
func main() {
	intArr := []int{2, 3, 4, 1, 5, 6}
	mergeSort(intArr)

	fmt.Println(intArr)
}

func mergeSort(arr []int) {
	arrLength := len(arr)
	if arrLength <= 1 {
		return
	}

	// get middle of array
	middle := arrLength / 2

	// create left and right arrays
	leftArr := make([]int, middle)
	rightArr := make([]int, arrLength-middle)

	// Copy array to left and right subarrays
	copy(leftArr, arr[:middle])
	copy(rightArr, arr[middle:])

	// Sort left and right arrays
	mergeSort(leftArr)
	mergeSort(rightArr)

	// Merge sorted left and right subarrays
	merge(arr, leftArr, rightArr)
}

func merge(arr, leftArr, rightArr []int) {
	leftIdx, rightIdx, mergedIdx := 0, 0, 0

	// Compare left and right array item and combine
	for leftIdx < len(leftArr) && rightIdx < len(rightArr) {
		if leftArr[leftIdx] <= rightArr[rightIdx] {
			arr[mergedIdx] = leftArr[leftIdx]
			leftIdx++
		} else {
			arr[mergedIdx] = rightArr[rightIdx]
			rightIdx++
		}
		mergedIdx++
	}

	// add remaining elements in the left subarray to result
	for leftIdx < len(leftArr) {
		arr[mergedIdx] = leftArr[leftIdx]
		leftIdx++
		mergedIdx++
	}

	// add remaining elements in the right subarray to result
	for rightIdx < len(rightArr) {
		arr[mergedIdx] = rightArr[rightIdx]
		rightIdx++
		mergedIdx++
	}
}
