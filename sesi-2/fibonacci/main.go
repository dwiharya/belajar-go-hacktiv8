package main

import (
	"fmt"
)

func main() {
	fmt.Println(Fib(0))
	fmt.Println(Fib(1))
	fmt.Println(Fib(2))
	fmt.Println(Fib(3))
	fmt.Println(Fib(4))

	// fmt.Println(Fib(0)) --> expected cetak angka 0
	// fmt.Println(Fib(1)) --> expected cetak angka 1
	// fmt.Println(Fib(2)) --> expected cetak angka 1
	// fmt.Println(Fib(3)) --> expected cetak angka 2
	// fmt.Println(Fib(4))

	// --> expected cetak angka 3

}

//

// Fib adalah function untuk menghitung angka Fibonacci ke n
// 1, 1, 2, 3, 5, 8
// Fib(0) = 0
// Fib(1) = 1
// Fib(2) = Fib(1) + Fib(0) = 1 + 0 = 1
// Fib(3) = Fib(2) + Fib(1) = 1 + 1 = 2
// Fib(4) = Fib(3) + Fib(2) = 2 + 1 = 3
func Fib(n int) int {
	if n < 2 {
		return n
	}

	return Fib(n-1) + Fib(n-2)
}
