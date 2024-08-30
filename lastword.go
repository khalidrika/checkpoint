package main

import (
	"fmt"
)

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
}

func LastWord(s string) string {
	var word string
	end := len(s) - 1

	// Skip any trailing spaces at the end of the string
	for end >= 0 && s[end] == ' ' {
		end--
	}

	// Iterate backwards to find the start of the last word
	for end >= 0 && s[end] != ' ' {
		word = string(s[end]) + word
		end--
	}

	return word + "\n"

}
