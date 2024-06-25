package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 0 {
		return
	}

	args := os.Args[1:]
	count := 0

	for _, char := range args {
		for range char {
			count++
		}
	}
	fmt.Println(count)

}
