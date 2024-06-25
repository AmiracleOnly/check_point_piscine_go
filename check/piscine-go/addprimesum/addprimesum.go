package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println(0)
		return
	}
	res := 0
	num := Atoi(args[0])
	if num <= 1 {
		fmt.Println(0)
		return
	}

	for i := num; i > 1; i-- {
		// fmt.Println(i)
		if IsPrime(i) {
			res += i
		}
	}
	fmt.Println(res)
}

func IsPrime(num int) bool {

	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func Atoi(s string) int {
	res := 0
	for _, char := range s {
		res = res*10 + int(char-'0')
	}
	return res
}
