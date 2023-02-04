package main

import (
	"fmt"
)

func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
		return false
	}

	m1 := make(map[byte]byte)
	m2 := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		if m1[s[i]] == 0 {
			m1[s[i]] = t[i]

		} else if m1[s[i]] != t[i] {
			return false
		}

		if m2[t[i]] == 0 {
			m2[t[i]] = s[i]
		} else if m2[t[i]] != s[i] {
			return false
		}
	}
    return true
}

func main() {
	s := "egg"
	t := "asd"
	fmt.Println(isIsomorphic(s, t))
}