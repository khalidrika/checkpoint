package main

import (
	"fmt"
)

func main() {
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("hello   .........  bye"))
}

func FirstWord(s string) string {
	var a string
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			a += string(s[i])
		} else {
			break
		}
	}
	return a + "\n"
}
