package main

import (
	"fmt"
)

func main() {
	fmt.Println(DigitLen(100, 10))
	fmt.Println(DigitLen(100, 2))
	fmt.Println(DigitLen(-100, 16))
	fmt.Println(DigitLen(100, -1))
}

func DigitLen(num, base int) int {

	count := 0
	if base < 2 || base > 36 {
		return -1
	}

	// count := 0

	if num < 0 {
		num *= -1
	}

	for num > 0 {
		num /= base
		count++
	}
	return count
}
