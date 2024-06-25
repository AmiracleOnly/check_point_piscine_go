package main

import "fmt"

func main() {
	table := []int{1, 2, 3}
	ac := 93
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
	fmt.Println()

	table = []int{0}
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
}

func FoldInt(f func(int, int) int, a []int, n int) {
	res := n

	for _, v := range a {
		res = f(res, v)
	}
	fmt.Println(res)
}

func Mul(a, b int) int {
	return a * b
}

func Sub(a, b int) int {
	return a - b
}

func Add(a, b int) int {
	return a + b
}
