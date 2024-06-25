package main

import (
	"fmt"
)

func main() {
	fmt.Println(ByeByeFirst([]string{}))
	fmt.Println(ByeByeFirst([]string{"one arg"}))
	fmt.Println(ByeByeFirst([]string{"first", "second"}))
	fmt.Println(ByeByeFirst([]string{"", "abcd", "efg"}))
}

func ByeByeFirst(strings []string) []string {
	var str []string

	for i, char := range strings {
		if i > 0 {
			str = append(str, char)
		}
	}
	return str
}