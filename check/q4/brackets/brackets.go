package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) == 0 {
		return
	} else if len(args) == 1 {
		fmt.Println()
	} else {
		for _, v := range os.Args[1:] {
			if Matched(v) {
				fmt.Println("OK")
			} else {
				fmt.Println("ERROR")
			}
		}
	}
}

func Matched(s string) bool {
	stack := []rune{}

	brackets := map[rune]rune{'(': ')', '[': ']', '{': '}'}

	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else if char == ')' || char == ']' || char == '}' {
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if brackets[last] != char {
				return false
			}
		}
	}
	return len(stack) == 0

}
