package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		return
	}

	a, err := strconv.Atoi(args[0])
	if err != nil || a < 0 {
		fmt.Println("ERROR")
		return
	}
	fmt.Println(itoaHex(a))

}

func itoaHex(num int) string {
	a := "0123456789abcdef"

	res := ""

	for num > 0 {
		digit := num % 16
		res = string(a[digit]) + res
		num /= 16
	}

	if res == "" {
		res = "0"
	}
	return res

}
