package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 || len(os.Args) == 0 {
		return
	}

	a, err1 := strconv.Atoi(os.Args[1])
	if err1 != nil {
		os.Stdout.WriteString("strconv.Atoi: parsing " + "\"" + os.Args[1] + "\"invalid syntax")
	}
	b, err2 := strconv.Atoi(os.Args[2])
	if err2 != nil {
		os.Stdout.WriteString("strconv.Atoi: parsing " + "\"" + os.Args[2] + "\"invalid syntax")
	}

	fmt.Print(b)

	for a != b {
		if b < a {
			b++
		} else {
			b--
		}
		fmt.Print(" ", b)
	}
	fmt.Print()
}
