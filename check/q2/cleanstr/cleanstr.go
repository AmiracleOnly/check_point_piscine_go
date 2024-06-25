package main

import (
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	s := os.Args[1]

	str := []string{}
	prev := 0
	word := false

	for i, c := range s {
		if c == ' ' {
			if word {
				str = append(str, s[prev:i])
			}
			prev = i + 1
			word = false
		} else {
			word = true
		}

	}
	if word {
		str = append(str, s[prev:])
	}

	res := ""
	for i, char := range str {
		if i > 0 {
			res += " "
		}
		res += char
	}
	os.Stdout.WriteString(res)
}
