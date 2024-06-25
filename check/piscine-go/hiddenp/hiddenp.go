package main

import "os"

func main() {
	args := os.Args
	if len(os.Args) != 3 {
		return
	}

	word1 := args[1]
	word2 := args[2]

	if len(word1) == 0 {
		os.Stdout.WriteString("1\n")
		return
	}

	var res string
	count := 0

	for _, char := range word1 {
		for i := count; i < len(word2); i++ {
			if char == rune(word2[i]) {
				res += string(char)
				count = i + 1
				break
			}
		}
	}
	if word1 == res {
		os.Stdout.WriteString("1\n")
	} else {
		os.Stdout.WriteString("0\n")
	}

	// os.Stdout.WriteString("\n")

}
