package main

import (
	"fmt"
)

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}

func RepeatAlpha(s string) string {
	res := []byte{}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 'a' && c <= 'z' {
			v := int(c - 'a' + 1)
			for j := 0; j < v; j++ {
				res = append(res, c)
			}
		} else if c >= 'A' && c <= 'Z' {
			v := int(c - 'A' + 1)
			for j := 0; j < v; j++ {
				res = append(res, c)
			}
		} else {
			res = append(res, c)
		}
	}
	return string(res)
}
