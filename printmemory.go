package main

import "fmt"

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

func PrintMemory(arr [10]byte) {
	for i := 0; i < 10; i += 4 {
		for j := i; j < i+4 && j < 10; j++ {
			fmt.Printf("%02x ", arr[i])
		}
		fmt.Println()
	}
	for i := 0; i < 10; i++ {
		if arr[i] >= 32 && arr[i] <= 126 {
			fmt.Printf("%c", arr[i])
		} else {
			fmt.Printf(".")
		}
	}
	fmt.Println()
}
