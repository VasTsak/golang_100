package main

import (
	"fmt"
)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	} else {
		return (fibonacci(n-1) + fibonacci(n-2))
	}
}

func main() {
	fmt.Println("Fibbonaci of 5 is: ", fibonacci(5))
	fmt.Println("Fibbonaci of 15 is: ", fibonacci(15))
	fmt.Println("Fibbonaci of 25 is: ", fibonacci(25))
}
