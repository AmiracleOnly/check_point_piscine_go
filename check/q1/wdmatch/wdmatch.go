package main

import "os"

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		return
	}

	a := args[0]
	b := args[1]

	lenght := len(args[1])
	var result string

	count := 0

	for _, char := range a {
		for i := count; i < lenght; i++ {
			if char == rune(b[i]) {
				result += string(char)
				count = i + 1
				break
			}
		}
	}
	if a == result {
		os.Stdout.WriteString(result)
	}
}
