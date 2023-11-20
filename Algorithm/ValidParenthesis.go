package main

import "fmt"

func isValid(s string) bool {
	// first way
	/*// create a stack
	runeStack := []rune{}

	// iterate over the string
	for _, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			runeStack = append(runeStack, ch)
		} else if ch == ')' && len(runeStack) > 0 && runeStack[len(runeStack)-1] == '(' {
			// pop the stack
			runeStack = runeStack[:len(runeStack)-1]
		} else if ch == ']' && len(runeStack) > 0 && runeStack[len(runeStack)-1] == '[' {
			// pop the stack
			runeStack = runeStack[:len(runeStack)-1]
		} else if ch == '}' && len(runeStack) > 0 && runeStack[len(runeStack)-1] == '{' {
			// pop the stack
			runeStack = runeStack[:len(runeStack)-1]
		} else {
			return false
		}
	}

	// verify if stack is empty
	return len(runeStack) == 0*/

	// second way
	// create a stack
	runeStack := []rune{}
	mapping := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	// iterate over the string
	for _, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			runeStack = append(runeStack, ch)
		} else if ch == ')' || ch == ']' || ch == '}' {
			if len(runeStack) == 0 || runeStack[len(runeStack)-1] != mapping[ch] {
				return false
			}
			runeStack = runeStack[:len(runeStack)-1]
		}
	}

	// verify if stack is empty
	return len(runeStack) == 0
}

func main() {
	fmt.Println(isValid("()"))
}
