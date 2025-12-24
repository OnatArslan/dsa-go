package main

import "fmt"

func main() {

	valid := isValid("[({))}]")
	fmt.Printf("valid: %v\n", valid)
}

func isValid(s string) bool {
	return true
}
