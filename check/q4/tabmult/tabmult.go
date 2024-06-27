package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 2 {
		number, _ := strconv.Atoi(os.Args[1])
		tabmult(number)
	} else {
		fmt.Println()
	}

}

func tabmult(num int) {
	i := 1

	for i < 10 {
		result := num * i
		fmt.Println(i, "x", num, "=", result)
		i++
	}
}
