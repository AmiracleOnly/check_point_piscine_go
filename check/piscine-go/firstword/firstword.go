package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	text := args[0]

	word := false
	prev := 0
	ss := []string{}

	for i, char := range text {
		if char == ' ' {
			if word {
				ss = append(ss, text[prev:i])
				// fmt.Println(text[prev:i])
				word = false
			}
			prev = i + 1
		} else {
			word = true
		}
	}

	if word {
		ss = append(ss, text[prev:])
	}
	fmt.Println(ss[len(ss)-1])
}
