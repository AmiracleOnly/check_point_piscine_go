package main

import (
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}

	text := args[0]
	a := []byte(args[1])
	b := []byte(args[2])

	result := ""

	for _, char := range text {
		if char == rune(a[0]) {
			char = rune(b[0])
		}
		result += string(char)
	}
	os.Stdout.WriteString(result)

}
