package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(Itoa(12345))
// 	fmt.Println(Itoa(0))
// 	fmt.Println(Itoa(-1234))
// 	fmt.Println(Itoa(987654321))
// }

// func Itoa(n int) string {
// 	base := "0123456789"
// 	res := ""
// 	isnegative := false
// 	if n == 0 {
// 		return "0"
// 	}

// 	if n < 0 {
// 		n *= -1
// 		isnegative = true
// 	}
// 	for n > 0 {
// 		res = string(base[n%10]) + res
// 		n /= 10
// 	}
// 	if isnegative {
// 		res = "-" + res
// 	}
// 	return res
// }
