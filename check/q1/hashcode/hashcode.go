package main

import "fmt"

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}

func HashCode(dec string) string {
	var result string
	size := len(dec)

	for _, char := range dec {
		hash := char + rune(size)%127
		if hash < 32 || hash > 126 {
			hash = hash + 33
		}
		result = result + string(hash)
	}
	return result
}
