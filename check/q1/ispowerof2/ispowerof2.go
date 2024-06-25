package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		return
	}

	args := os.Args[1]

	power := 2
	num := Atoi(args)

	for power < num {
		power = power * 2
	}

	if power > num {
		fmt.Println("False")
	} else {
		fmt.Println("True")
	}

}

func Atoi(s string) int {
	result := 0

	for _, char := range s {
		if char >= '0' && char <= '9' {
			result = result*10 + int(char-'0')
		} else {
			return 0
		}
	}
	return result
}
