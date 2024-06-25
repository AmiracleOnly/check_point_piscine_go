package main

import (
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		return
	}

	str := args[0]
	s := []string{}
	prev := 0
	word := false

	for i, char := range str {
		if char == ' ' {
			if word {
				s = append(s, str[prev:i])
				word = false
			}
			prev = i + 1
		} else {
			word = true
		}
	}

	if word {
		s = append(s, str[prev:])
	}

	// res := ""
	for i, char := range s {
		if i > 0 {
			os.Stdout.WriteString("   ")
		}
		os.Stdout.WriteString(char)
	}
	os.Stdout.WriteString("\n")

}
