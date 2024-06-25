package main

import (
	"fmt"
)

func main() {
	fmt.Println(HalfSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(HalfSlice([]int{1, 2, 3}))
}

func HalfSlice(slice []int) []int {
	lenght := len(slice) + 1
	half := (lenght) / 2
	return slice[:half]
}
