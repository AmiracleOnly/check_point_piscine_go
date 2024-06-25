package main

import "os"

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		os.Stdout.WriteString("\n")
		return
	}

	text1 := args[0]
	text2 := args[1]

	s := make(map[rune]bool)

	for _, c := range text1 + text2 {
		if s[c] {
			continue
		} else {
			s[c] = true
			os.Stdout.WriteString(string(c))

		}
	}
	os.Stdout.WriteString("\n")

}
