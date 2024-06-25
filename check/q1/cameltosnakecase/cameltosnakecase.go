package main

import (
	"fmt"
)

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}

func CamelToSnakeCase(s string) string {
	var result []byte

	if prefUpper(s) {
		for i, char := range s {
			if IsUpper(char) && i != 0 {
				result = append(result, '_')
				result = append(result, byte(char))
			} else {
				result = append(result, byte(char))
			}
		}
	} else {
		result = []byte(s)
	}
	return string(result)
}

func prefUpper(str string) bool {

	var flag bool
	for _, char := range str {
		if IsUpper(char) {
			if flag {
				return false
			}
			flag = true
		} else {
			flag = false
		}
	}
	return true
}

func IsUpper(c rune) bool {
	return c >= 'A' && c <= 'Z'
}
