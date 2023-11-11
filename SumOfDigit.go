package main

import "fmt"

func main() {
	fmt.Println(sumOfDigit(11))
}

func sumOfDigit(num int) int {
	var sum int = 0

	for num > 0 {
		sum += num % 10
		num /= 10
	}

	return sum
}
