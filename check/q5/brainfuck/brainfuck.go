package main

import (
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	var array [2048]byte
	ptr := 0
	args := os.Args[1]

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case '>':
			ptr++
		case '<':
			ptr--
		case '+':
			array[ptr]++
		case '-':
			array[ptr]--
		case '.':
			os.Stdout.Write(array[ptr : ptr+1])
		case '[':
			if array[ptr] == 0 {
				count := 0
				for j := i + 1; j < len(args); j++ {
					if args[j] == '[' {
						count++
					} else if args[j] == ']' {
						if count == 0 {
							i = j
							break
						}
						count--
					}
				}
			}
		case ']':
			if array[ptr] != 0 {
				count := 0
				for j := i - 1; j >= 0; j-- {
					if args[j] == ']' {
						count++
					} else if args[j] == '[' {
						if count == 0 {
							i = j
							break
						}
						count--
					}
				}
			}
		}
	}
}
