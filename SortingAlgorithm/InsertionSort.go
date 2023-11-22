package main

import "fmt"

/*
* It performs the sorting process by check data starting from the last.
* Time complexity is O(n^2)
* Space complexity is O(1). Because we are operating on the existing array.
 */
func main() {
	numbers := []int{5, 2, 7, 9, 1, 0, 6, 3, 4}
	names := []string{"Ali", "Veli", "Ayşe", "Fatma"}

	insertionSort(numbers)
	insertionSort(names)

	fmt.Println(numbers)
	fmt.Println(names)
}

func insertionSort(arr interface{}) {
	switch array := arr.(type) {
	case []int:
		for i := 1; i < len(array); i++ {
			key := array[i]
			j := i - 1

			for j >= 0 && array[j] > key {
				array[j+1] = array[j]
				j--
			}

			array[j+1] = key
		}

	case []string:
		for i := 1; i < len(array); i++ {
			key := array[i]
			j := i - 1

			for j >= 0 && array[j] > key {
				array[j+1] = array[j]
				j--
			}

			array[j+1] = key
		}
	}
}
