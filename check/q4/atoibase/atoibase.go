package main

import (
	"fmt"
)

func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}

func AtoiBase(s string, base string) int {
	if len(base) == 0 || len(base) < 2 || len(s) == 0 {
		return 0
	}

	for _, char := range base {
		if char == '-' || char == '+' {
			return 0
		}
	}

	basemap := make(map[rune]int)

	for i, char := range base {
		if _, exist := basemap[char]; exist {
			return 0
		}
		basemap[char] = i
	}

	res := 0
	baselen := len(base)

	for _, c := range s {
		if _, exist := basemap[c]; !exist {
			return 0
		}
		res = res*baselen + basemap[c]
	}
	return res
}
