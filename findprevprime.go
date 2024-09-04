package main

// import (
// 	"github.com/01-edu/go-tests/lib/challenge"
// 	"github.com/01-edu/go-tests/lib/random"
// 	"github.com/01-edu/go-tests/solutions"
// )

// func main() {
// 	// fmt.Println(isprime(22216))
// 	args := append(random.IntSliceBetween(0, 99999), 5, 4, 1, 0)
// 	for _, arg := range args {
// 		challenge.Function("FindPrevPrime", FindPrevPrime, solutions.FindPrevPrime, arg)
// 	}
// }

// func isprime(n int) bool {
// 	if n <= 1 {
// 		return false
// 	}
// 	if n == 2 {
// 		return true
// 	}
// 	if n == 3 {
// 		return true
// 	}
// 	if n%2 == 0 {
// 		return false
// 	}
// 	for i := 3; i*i <= n; i += 2 {
// 		if n%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func FindPrevPrime(nb int) int {
// 	for i := nb; i >= 2; i-- {
// 		if isprime(i) {
// 			return i
// 		}
// 	}
// 	return 0
// }

// func FindPrevPrime(n int) int {
// 	if n <2 {
// 		return 0
// 	}
// 	if isprime(n) {
// 		return n
// 	}
	
// 	return FindPrevPrime(n-1)
// }
