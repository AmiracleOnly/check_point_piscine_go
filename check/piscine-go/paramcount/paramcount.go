package main

import (
	"os"
)

func main() {
	if len(os.Args) == 0 {
		return
	}
	args := os.Args[1:]

	count := len(args)
	os.Stdout.WriteString(itoa(count))

}

func itoa(num int) string {
	if num == 0 {
		return "0"
	}

	res := ""
	for num > 0 {
		digit := num % 10
		res = string(rune(digit+'0')) + res
		num /= 10
	}
	return res
}
