package main

import (
	"fmt"
)

func main() {
	fmt.Println(BeZero([]int{1, 2, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31}))
	fmt.Println(BeZero([]int{}))
}

func BeZero(slice []int) []int {
	res := 0

	for i := range slice {
		slice[i] = res
	}
	return slice
}
