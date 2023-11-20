package main

import "fmt"

func finalValueAfterOperations(operations []string) int {
	value := 0
	for _, operation := range operations {
		if operation == "X++" || operation == "++X" {
			value++
		} else {
			value--
		}
	}

	return value
}

func main() {
	fmt.Println(finalValueAfterOperations([]string{"--X", "X++", "X++"}))
}
