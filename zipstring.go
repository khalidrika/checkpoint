package main

import (
	"fmt"
)

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}

func ZipString(s string) string {
	a := 1
	res := ""

	for i := 0; i < len(s); i++ {
		if i != 0 && s[i] == s[i-1] {
			continue
		} else if len(s) > i+a {
			for s[i] == s[i+a] {
				a++
			}
		}
		res += string(48+a) + string(s[i])
		a =1

	}
	return res
}
