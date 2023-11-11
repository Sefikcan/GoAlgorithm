package main

import "fmt"

func main() {
	fmt.Println(isPalindromeNumber(43))
}

func isPalindromeNumber(num int) bool {
	var originalNumber int = num
	var reversedNumber int = 0

	for num > 0 {
		digit := num % 10
		reversedNumber = reversedNumber*10 + digit
		num /= 10
	}

	return originalNumber == reversedNumber
}
