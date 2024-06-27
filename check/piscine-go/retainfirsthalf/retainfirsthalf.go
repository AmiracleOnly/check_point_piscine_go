package main

import (
	"fmt"
)

func main() {
	fmt.Println(RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(RetainFirstHalf("write %d ==> 45m$"))
	fmt.Println(RetainFirstHalf("Kimetsu no Yaiba"))
	fmt.Println(RetainFirstHalf("-552"))
	fmt.Println(RetainFirstHalf("123@live.fr"))
}


func RetainFirstHalf(str string) string {
	word := ""

	for i, char := range str {
		if i == 0 {
			word += string(char)
		} else if Upper(char) {
			break
		} else {
			word += string(char)
		}
	}
	return word
}

func Upper(char rune) bool {
	return char >= 'A' && char <= 'Z'
}
