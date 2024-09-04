 package main

// import (
// 	"os"

// 	"github.com/01-edu/z01"
// )

// func main() {
// 	if len(os.Args) != 2 {
// 		z01.PrintRune('\n')
// 		return
// 	}
// 	input := os.Args[1]
// 	if len(input) == 0 {
// 		z01.PrintRune('\n')
// 		return
// 	}

// 	start := 0
// 	end := len(input) - 1
// 	for start < end && input[start] == ' ' || input[start] == '\t' {
// 		start++
// 	}
// 	for start < end && input[end] == ' ' || input[end] == '\t' {
// 		end--
// 	}
// 	if start > end {
// 		z01.PrintRune('\n')
// 		return
// 	}

// 	var out []byte
// 	isword := false

// 	for i := start; i < end; i++ {
// 		if input[i] == ' ' || input[i] == '\t' {
// 			if isword {
// 				out = append(out, ' ', ' ', ' ')
// 				isword = false
// 			}
// 		} else {
// 			out = append(out, input[i])
// 			isword = true
// 		}
// 	}
// 	for i := 0; i < len(out); i++ {
// 		z01.PrintRune(rune(out[i]))
// 	}
// 	z01.PrintRune('\n')
// }