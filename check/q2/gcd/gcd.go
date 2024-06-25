package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) != 3 || len(args) == 0 {
		return
	}

	num1 := atoi(args[1])
	num2 := atoi(args[2])

	num := gcd(num1, num2)
	fmt.Println(num)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func atoi(s string) int {
	res := 0
	sign := 0

	for i, char := range s {
		if i == 0 && char == '-' {
			sign *= -1
		} else if char >= '0' && char <= '9' {
			res = res*10 + int(char-'0')
		}
	}
	return res
}
