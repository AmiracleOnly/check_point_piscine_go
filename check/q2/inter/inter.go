package main

import "os"

func main() {
	args := os.Args

	if len(args) != 3 || len(args) == 0 {
		return
	}

	text := args[1]
	word := args[2]

	var res string

	s := make(map[rune]bool)

	for _, char := range word {
		s[char] = true
	}

	for _, i := range text {
		if s[i] {
			res += string(i)
			delete(s, i)
		}
	}
	os.Stdout.WriteString(res)
	os.Stdout.WriteString("\n")
}
