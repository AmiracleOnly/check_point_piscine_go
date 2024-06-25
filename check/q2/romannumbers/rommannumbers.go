package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args[1:]) != 1 {
		Invalid()
		return
	}

	args, err := Atoi(os.Args[1])
	if !err || args == 0 || args >= 4000 {
		Invalid()
		return
	}

	arab := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	roman := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	calc := []string{"M+", "(M-C)+", "D+", "(D-C)+", "C+", "(C-X)+", "L+", "(L-X)+", "X+", "(X-I)+", "V+", "(V-I)+", "I+"}

	var calculate, rn string

	for i, char := range arab {
		for args >= arab[i] {
			rn += roman[i]
			calculate += calc[i]
			args -= char
		}
	}

	calculate = calculate[:len(calculate)-1]
	fmt.Printf("%s\n%s\n", calculate, rn)

}

func Atoi(s string) (int, bool) {
	res := 0
	sign := 1

	for i, char := range s {
		if i == 0 && char == '-' {
			sign *= -1
		} else if char >= '0' && char <= '9' {
			res = res*10 + int(char-'0')
		}
	}
	return res * sign, true
}

func Invalid() {
	fmt.Println("ERROR: cannot convert to roman digit")
}
