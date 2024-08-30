package main

import "fmt"

func main() {
	fmt.Println(CheckNumber("Hello"))
	fmt.Println(CheckNumber("Hello1"))
}

func CheckNumber(arg string) bool {

	for _, i := range arg {
		if i >= '0' && i <= '9' {
			return true
		}
	}
	return false

}
