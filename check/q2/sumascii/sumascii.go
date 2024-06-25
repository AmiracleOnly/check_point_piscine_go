package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("0")
		return
	}

	text := args[0]

	var sum byte
	for _, char := range text {
		sum += byte(char)
	}
	fmt.Println(sum)

	sum = 0
	for i := 0; i < len(text); i++ {
		char := text[i]
		sum += byte(char)
	}
	fmt.Println(sum)
}
