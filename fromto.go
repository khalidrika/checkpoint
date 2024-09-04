 package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func FromTo(from int, to int) string {
// 	if from < 0 || from > 99 || to < 0 || to > 99 {
// 		return "Invalid\n"
// 	}
// 	if from == to {
// 		return strconv.Itoa(from) + "\n"
// 	}
// 	result := ""
// 	if from > to {
// 		for i := from; i >= to; i-- {
// 			if i < 10 {
// 				result += "0" + strconv.Itoa(i) + ", "
// 			} else {
// 				result += strconv.Itoa(i) + ", "
// 			}
// 		}
// 	} else {
// 		for i := from; i <= to; i++ {
// 			if i < 10 {
// 				result += "0" + strconv.Itoa(i) + ", "
// 			} else {
// 				result += strconv.Itoa(i) + ", "
// 			}
// 		}
// 	}

// 	return result + "\n"
// }

// func main() {
// 	fmt.Print(FromTo(1, 10))
// 	fmt.Print(FromTo(10, 1))
// 	fmt.Print(FromTo(10, 10))
// 	fmt.Print(FromTo(100, 10))
// }