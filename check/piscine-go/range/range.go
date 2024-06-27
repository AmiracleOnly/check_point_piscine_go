package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// args := os.Args[1:]

	if len(os.Args) != 3 || len(os.Args) == 0 {
		return
	}

	a, err1 := strconv.Atoi(os.Args[1])
	b, err2 := strconv.Atoi(os.Args[2])

	if err1 != nil {
		os.Stdout.WriteString("strconv.Atoi: parcinv: " + "\"" + os.Args[1] + "\": invalid syntax")
	}
	if err2 != nil {
		os.Stdout.WriteString("strconv.Atoi: parcinv: " + "\"" + os.Args[2] + "\": invalid syntax")
	}

	fmt.Print(a)
	// result := ""
	for i := a + 1; i <= b; i++ {
		fmt.Print(" ", i)
	}
	fmt.Print()
}
