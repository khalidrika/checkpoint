package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(IsCapitalized("Hello! How are you?"))
// 	fmt.Println(IsCapitalized("Hello How Are You"))
// 	fmt.Println(IsCapitalized("Whats 4this 100K?"))
// 	fmt.Println(IsCapitalized("Whatsthis4"))
// 	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
// 	fmt.Println(IsCapitalized(""))
// }

// func IsCapitalized(s string) bool {
// 	if s == "" {
// 		return false
// 	}
// 	for i := 0; i < len(s); i++ {
// 		if s[i] == ' ' || s[i] == '\t' {
// 			fmt.Println("1")
// 			if s[i]+1 >= 'a' || s[i]+1 <= 'z' {
// 				fmt.Println(s[i])
// 				return false
// 			} else {
// 				return true
// 			}
// 		}
// 	}
// 	if s[0] >= 'a' && s[0] <= 'z' {
// 		return false
// 	} else {
// 		return true
// 	}
// 	return false
// }
