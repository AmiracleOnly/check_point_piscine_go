package main

import (
	"fmt"
)

func main() {
	fmt.Println(SetSpace("loWerCaseInBeginning"))
	fmt.Println(SetSpace("RightWordLargeWorldALotOfWords"))
	fmt.Println(SetSpace(" "))
	fmt.Println(SetSpace("Wron2Word"))
	fmt.Println(SetSpace("#@13word"))
}

func SetSpace(s string) string {
	var result string

	a := len(s)
	if s == "" {
		return ""
	}

	for i := 0; i < a; i++ {
		if i != 0 && (s[i] >= 'A' && s[i] <= 'Z') {
			result += " " + string(s[i])
		} else if (s[i] >= 'a' && s[i] <= 'z') && i != 0 || i == 0 && (s[i] >= 'A' && s[i] <= 'Z') {
			result += string(s[i])
		} else {
			return "Error"
		}
	}
	return result

}
