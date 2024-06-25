package main

import (
	"fmt"
)

func main() {
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
	fmt.Println(AtoiBase("125", "0123456789"))
}

func AtoiBase(s string, base string) int {
	res := 0
	sign := 1

	if len(base) < 2 {
		return 0
	}
	// if base is wrong return 0

	d := make(map[rune]int)
	for i, c := range base {
		if c == '-' || c == '+' {
			return 0
		}
		if d[c] == 0 {
			d[c] = i + 1
		} else {
			return 0
		}
	}

	for i, char := range s {
		if i == 0 && char == '-' {
			sign *= -1
			continue
		}

		index := d[char]
		if index != 0 {
			res = res*len(base) + index - 1
		} else {
			return 0
		}
	}
	return sign * res
}
