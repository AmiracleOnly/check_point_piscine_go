package main

import "os"

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
			}
			word = false
			prev = i + 1
		} else {
			word = true
		}
	}

	if word {
		s = append(s, str[prev:])
	}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	result := ""
	for i, c := range s {
		if i > 0 {
			result += " "
		}
		result += c
	}
	os.Stdout.WriteString(result)
	os.Stdout.WriteString("\n")

}
