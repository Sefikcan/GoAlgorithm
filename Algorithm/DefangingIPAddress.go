package main

import (
	"fmt"
)

func defangIPaddr(address string) string {
	defang := ""

	// walk over the range loop for input address
	// build the defanged address
	for _, ch := range address {
		if ch == '.' {
			defang += string("[.]")
		} else {
			defang += string(ch)
		}
	}

	return defang
}

func main() {
	fmt.Println(defangIPaddr("1.1.1.1"))
}
