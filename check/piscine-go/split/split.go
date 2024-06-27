package main

import (
	"fmt"
)

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}

func Split(s, sep string) []string {
	if s == "" || sep == "" {
		return []string{}
	}

	ss := []string{}
	prev := 0
	l := len(sep)

	for i := 0; i < len(s)-l; i++ {
		if s[i:i+l] == sep {
			ss = append(ss, s[prev:i])
			prev = i + l
		}
	}
	ss = append(ss, s[prev:])
	return ss
}
