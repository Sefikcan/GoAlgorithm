package main

import "fmt"

/*
* We take an element from array and perform sorting operations by going through the entire list
* As the size of array increases, it causes performance problem.
* Time complexity is O(n^2)
* Space complexity is O(1). Because we are operating on the existing array.
 */
func main() {
	numbers := []int{5, 2, 7, 9, 1, 0, 6, 3, 4}
	names := []string{"Ali", "Veli", "Ayşe", "Fatma"}

	selectionSort(numbers)
	selectionSort(names)

	fmt.Println(numbers)
	fmt.Println(names)
}

func selectionSort(arr interface{}) {
	switch array := arr.(type) {
	case []int:
		for i := 0; i < len(array); i++ {
			minIndex := i
			for j := i + 1; j < len(array); j++ {
				if array[j] < array[minIndex] {
					minIndex = j
				}
			}

			array[minIndex], array[i] = array[i], array[minIndex]
		}
	case []string:
		for i := 0; i < len(array); i++ {
			minIndex := i
			for j := i + 1; j < len(array); j++ {
				if array[j] < array[minIndex] {
					minIndex = j
				}
			}

			array[minIndex], array[i] = array[i], array[minIndex]
		}
	}
}
