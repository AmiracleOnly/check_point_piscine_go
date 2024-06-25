package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(os.Args) != 2 {
		return
	}

	num, err := strconv.Atoi(args[0])
	if err != nil || num <= 0 {
		return
	}

	f := fprime(num)
	if len(f) == 0 {
		return
	}

	for i, v := range f {
		if i != 0 {
			fmt.Print("*")
		}
		fmt.Print(v)
	}
	fmt.Println()
}

func fprime(num int) []int {
	f := []int{}

	for i := 2; i*i <= num; i++ {
		for num%i == 0 {
			f = append(f, i)
			num /= i
		}
	}

	if num > 1 {
		f = append(f, num)
	}
	return f
}
