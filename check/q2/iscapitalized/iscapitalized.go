package main

import (
	"fmt"
)

func main() {
	fmt.Println(IsCapitalized("Hello How Are 4you"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized(" "))
	fmt.Println(IsCapitalized(""))
	fmt.Println(IsCapitalized("1"))
	fmt.Println(IsCapitalized("!"))
	fmt.Println(IsCapitalized("A"))
	fmt.Println(IsCapitalized("131!"))
	fmt.Println(IsCapitalized("H3110 W0r1d!"))
	fmt.Println(IsCapitalized("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"))
}

func IsCapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}

	if s == "" {
		return false
	}

	word := false

	for i, char := range s {
		if i == 0 && char >= 'a' && char <= 'z' {
			return false
		} else if char == ' ' {
			word = true
		} else if word {
			if char >= 'a' && char <= 'z' {
				return false
			} else {
				word = false
			}
		}
	}
	return true

}
