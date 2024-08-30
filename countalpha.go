package main

import (
	"fmt"
)

func main() {
	fmt.Println(CountAlpha("Hello world"))
	fmt.Println(CountAlpha("H e l l o"))
	fmt.Println(CountAlpha("H1e2l3l4o"))
}

func CountAlpha(s string) int {
	z := 0
	for _, i := range s {
		if i >= 65 && i <= 90 || i >= 97 && i <= 122 {
			z++
		}
	}
	return z
}
