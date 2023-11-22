package main

import "fmt"

func main() {
	numbers := []int{5, 2, 7, 9, 1, 0, 6, 3, 4}
	names := []string{"Ali", "Veli", "Ayşe", "Fatma"}

	bubbleSort(numbers)
	bubbleSort(names)

	fmt.Println(numbers)
	fmt.Println(names)
}

// this algorithm compare one by one and swap items
func bubbleSort(arr interface{}) {
	switch array := arr.(type) {
	case []int:
		for i := 0; i < len(array); i++ {
			for j := 0; j < len(array)-i-1; j++ {
				if array[j] > array[j+1] {
					// swap elements
					array[j], array[j+1] = array[j+1], array[j]
				}
			}
		}
	case []string:
		for i := 0; i < len(array); i++ {
			for j := 0; j < len(array)-i-1; j++ {
				if array[j] > array[j+1] {
					array[j], array[j+1] = array[j+1], array[j]
				}
			}
		}
	}
}
