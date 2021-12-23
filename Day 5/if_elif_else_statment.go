package main

import (
	"fmt"
)

func num_digits(number string) {
	if len(number) < 2 {
		fmt.Println("1")
	} else if len(number) < 3 {
		fmt.Println("2")
	} else {
		fmt.Println("greater than 2")
	}
}

func main() {
	num_1 := "1"
	num_2 := "12"
	num_3 := "123"
	num_4 := "1234"

	num_digits(num_1)
	num_digits(num_2)
	num_digits(num_3)
	num_digits(num_4)
}
