package main

import "fmt"

func main() {

	valid := isValid("[({))}]")
	fmt.Printf("valid: %v\n", valid)
}

func isValid(s string) bool {
	stack := make([]rune, 0, len(s))

	mp := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, ch := range s {

		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if top != mp[ch] {
			return false
		}
	}

	return len(stack) == 0
}
