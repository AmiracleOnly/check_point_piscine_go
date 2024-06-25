package main

import "fmt"

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}

func Chunk(slice []int, size int) {
	if len(slice) == 0 {
		return
	}

	s := len(slice)

	if size == 0 {
		return
	}

	var result [][]int

	for i := 0; i < s; i += size {
		end := i + size
		if end > s {
			end = s
		}
		result = append(result, slice[i:end])
	}
	fmt.Println(result)

}
