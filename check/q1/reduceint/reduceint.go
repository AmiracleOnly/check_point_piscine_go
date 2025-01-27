package main

import "fmt"

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{500, 2}
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)
}

func ReduceInt(a []int, f func(int, int) int) {
	slice := a[0]

	s := len(a)

	if s == 0 {
		return
	}

	if s >= 2 {
		for i := 1; i < s; i++ {
			slice = f(slice, a[i])
		}
	}
	fmt.Println(slice)

}
