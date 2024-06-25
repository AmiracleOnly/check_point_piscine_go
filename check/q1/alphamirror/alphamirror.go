package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}

	args := []rune(os.Args[1])

	for i, char := range args {
		if char >= 'a' && char <= 'z' {
			args[i] = 'z' - char + 'a'
		} else if char >= 'A' && char <= 'Z' {
			args[i] = 'Z' - char + 'A'
		}
	}
	fmt.Println(string(args))

}
