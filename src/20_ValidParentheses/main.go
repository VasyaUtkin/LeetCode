package main

import (
	"LeetCodeSolution/src/20_ValidParentheses/solution"
	"fmt"
)

func main() {

	fmt.Println(solution.IsValid("()"))
	fmt.Println(solution.IsValid("()[]{}"))
	fmt.Println(solution.IsValid("(]"))

}
