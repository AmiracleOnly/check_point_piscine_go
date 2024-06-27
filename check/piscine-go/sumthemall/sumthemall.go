package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(os.Args) < 2 {
		fmt.Println("0")
	}

	sum := 0

	for _, char := range args {
		num, err := strconv.Atoi(char)
		if err != nil {
			fmt.Println("0")
			return
		}

		if sum > 0 && num > 0 && sum > (1<<31-1-num) {
			fmt.Println("0")
			return
		}

		if sum < 0 && num < 0 && sum < (-1<<31-num) {
			fmt.Println("0")
			return
		}

		sum += num
	}
	fmt.Println(sum)
}

func Atoi(s string) int {
	res := 0
	sign := 0

	for i, char := range s {
		if i == 0 && char == '-' {
			sign += -1
			continue
		}

		if char >= '0' && char <= '9' {
			res = res*10 + int(char-'0')
		}
	}
	return sign * res
}
