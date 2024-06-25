package main

import (
	"fmt"
)

func main() {
	fmt.Println(Itoa(0))
	fmt.Println(Itoa(-1234))
	fmt.Println(Itoa(987654321))
	fmt.Println(Itoa(12345))
}

func Itoa(num int) string {
	res := ""
	sign := ""

	if num == 0 {
		return "0"
	}

	if num < 0 {
		sign += "-"
		num *= -1
	}

	for num > 0 {
		digit := num % 10
		res = string(rune(digit+'0')) + res
		num /= 10
	}
	return sign + res
}
