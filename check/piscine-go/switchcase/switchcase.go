package main

import "os"

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		return
	}

	word := ""
	for _, char := range args[0] {
		if char >= 'A' && char <= 'Z' {
			word += string(char + 32)
		} else if char >= 'a' && char <= 'z' {
			word += string(char - 32)
		} else {
			word += string(char)
		}
	}
	os.Stdout.WriteString(word)
}
