package main

import (
	"fmt"
)

func main() {
	fmt.Print(PrintIf("abcdefz"))
	fmt.Print(PrintIf("abc"))
	fmt.Print(PrintIf(""))
	fmt.Print(PrintIf("14"))
}

func PrintIf(str string) string {
	if len(str) > 3 || str == "" {
		return "G\n"
	} else {
		return ("Invalid Output\n")
	}

}
