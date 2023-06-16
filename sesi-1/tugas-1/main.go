package main

import "fmt"

func main() {
	for i := 0; i < 16; i++ {
		var f = "Fizz"
		var b = "Buzz"
		if i%3 == 0 {
			fmt.Println(f)
		}
		if i%5 == 0 {
			fmt.Println(b)
		}

		if i%3 == 0 && i%5 == 0 {
			fmt.Println(f + b)
		}

		if i%3 > 0 || i%5 > 0 {
			fmt.Println(i)
		}
	}
}
